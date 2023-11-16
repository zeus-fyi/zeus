package ai_codegen

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	openai "github.com/sashabaranov/go-openai"
	"github.com/zeus-fyi/zeus/pkg/hera"
	"github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/compression"
	"github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/memfs"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

func CodeGenRoutes(e *echo.Group) *echo.Group {
	e.POST("/ui/openai/codegen/upload", CreateCodeGenUploadPromptRequestHandler)
	e.POST("/ui/openai/codegen", ChatGptApiRequestHandler)
	return e
}

func ChatGptApiRequestHandler(c echo.Context) error {
	request := new(CodeGenAPIRequest)
	if err := c.Bind(request); err != nil {
		return err
	}
	return request.CompleteCodeGenRequest(c)
}

var addYourOwnUserIds = 1

func (ai *CodeGenAPIRequest) CompleteCodeGenRequest(c echo.Context) error {
	resp, err := hera.HeraOpenAI.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: "gpt-4-1106-preview",
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: ai.Prompt,
					Name:    fmt.Sprintf("%d", addYourOwnUserIds),
				},
			},
		},
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, resp.Choices[0].Message.Content)
}

func CreateCodeGenUploadPromptRequestHandler(c echo.Context) error {
	request := new(UICodeGenAPIRequest)
	if err := c.Bind(request); err != nil {
		return err
	}
	return request.CodeGenUploadPromptRequest(c)
}

func (ai *UICodeGenAPIRequest) CodeGenUploadPromptRequest(c echo.Context) error {
	file, err := c.FormFile("prompt")
	if err != nil {
		log.Err(err).Msg("CompleteCodeGenRequest: FormFile")
		return c.JSON(http.StatusInternalServerError, nil)
	}
	src, err := file.Open()
	if err != nil {
		log.Err(err).Msg("CompleteCodeGenRequest: file.Open()")
		return c.JSON(http.StatusInternalServerError, nil)
	}
	defer src.Close()
	in := bytes.Buffer{}
	if _, err = io.Copy(&in, src); err != nil {
		log.Err(err).Msg("CompleteCodeGenRequest: Copy")
		return c.JSON(http.StatusInternalServerError, nil)
	}
	prompt, err := UnGzipTextFiles(&in)
	if err != nil {
		log.Err(err)
		return c.JSON(http.StatusInternalServerError, nil)
	}

	resp, err := hera.HeraOpenAI.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: "gpt-4-1106-preview",
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
					Name:    fmt.Sprintf("%d", addYourOwnUserIds),
				},
			},
		},
	)
	return c.JSON(http.StatusOK, resp.Choices[0].Message.Content)
}

type UICodeGenAPIRequest struct {
	TokenEstimate int    `json:"tokenEstimate,omitempty"`
	Prompt        string `json:"prompt"`
}

type CodeGenAPIRequest struct {
	Prompt string `json:"prompt"`
}

func UnGzipTextFiles(in *bytes.Buffer) (string, error) {
	p := filepaths.Path{DirIn: "/tmp", DirOut: "/tmp", FnIn: "prompt.tar.gz"}
	m := memfs.NewMemFs()
	err := m.MakeFileIn(&p, in.Bytes())
	if err != nil {
		log.Err(err)
		return "", err
	}
	p.DirOut = "/prompt"
	comp := compression.NewCompression()
	err = comp.UnGzipFromInMemFsOutToInMemFS(&p, m)
	if err != nil {
		log.Err(err)
		return "", err
	}
	p.DirIn = "/prompt"
	return AppendFilesToString(m, p.DirOut)
}

func AppendFilesToString(fs memfs.MemFS, dir string) (string, error) {
	var buffer bytes.Buffer
	// Read the directory
	files, ferr := fs.ReadDir(dir)
	if ferr != nil {
		log.Err(ferr)
		return "", ferr
	}
	// Iterate through the files in the directory
	for _, file := range files {
		// Open the file
		f, err := fs.Open(filepath.Join(dir, file.Name()))
		if err != nil {
			log.Err(err)
			return "", err
		}
		// Read the file content
		content, err := io.ReadAll(f)
		if err != nil {
			log.Err(err)
			f.Close()
			return "", err
		}
		// Write the file content to the buffer
		_, err = buffer.Write(content)
		if err != nil {
			log.Err(err)
			f.Close()
			return "", err
		}
		f.Close()
	}
	return buffer.String(), nil
}
