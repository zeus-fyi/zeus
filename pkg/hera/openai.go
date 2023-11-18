package hera

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/sashabaranov/go-openai"
)

var HeraOpenAI OpenAI

var maxTokensByModel = map[string]int{
	openai.GPT3TextDavinci003: 2048,
	openai.GPT3TextDavinci002: 2048,
	openai.GPT3TextDavinci001: 2048,
	openai.GPT3TextAda001:     2048,
	openai.GPT3TextBabbage001: 2048,
}

type OpenAI struct {
	*openai.Client
}

func InitHeraOpenAI(bearer string) {
	HeraOpenAI = OpenAI{}
	HeraOpenAI.Client = openai.NewClient(bearer)
}

type OpenAIParams struct {
	Model     string
	MaxTokens int
	Prompt    string
}

func (ai *OpenAI) MakeCodeGenRequest(ctx context.Context, userID int, params OpenAIParams) (openai.CompletionResponse, error) {
	if len(params.Model) <= 0 {
		params.Model = openai.GPT4
	}

	req := openai.CompletionRequest{
		Model:     params.Model,
		MaxTokens: params.MaxTokens,
		Prompt:    params.Prompt,
		User:      fmt.Sprintf("%d", userID),
	}

	resp, err := ai.CreateCompletion(ctx, req)
	if err != nil {
		log.Err(err).Msg("MakeCodeGenRequest")
		return resp, err
	}
	return resp, err
}
