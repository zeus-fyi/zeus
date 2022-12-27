package hera_client

import (
	"context"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/compression"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

const heraCodeGenRoute = "/v1beta/openai/codegen"

func (h *HeraClient) UploadFiles(ctx context.Context, p filepaths.Path, model string) (interface{}, error) {
	var respJson interface{}
	err := h.ZipFilesInPath(&p)
	if err != nil {
		return respJson, err
	}
	resp, err := h.R().
		SetResult(&respJson).
		SetFormData(map[string]string{
			"model": model,
		}).
		SetFile("prompt", p.FileOutPath()).
		Post(heraCodeGenRoute)

	if err != nil || resp.StatusCode() != http.StatusOK {
		log.Ctx(ctx).Err(err).Msg("HeraClient: Upload Files")
		return respJson, err
	}
	h.PrintRespJson(resp.Body())
	return respJson, err
}

func (h *HeraClient) ZipFilesInPath(p *filepaths.Path) error {
	comp := compression.NewCompression()
	err := comp.GzipCompressDir(p)
	if err != nil {
		log.Err(err).Interface("path", p).Msg("HeraClient: ZipTextFileInPath")
		return err
	}
	return err
}
