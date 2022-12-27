package hera_client

import (
	"fmt"
	"os"
)

const (
	CodexCodeDavinci002 = "code-davinci-002"
	CodexCodeCushman001 = "code-cushman-001"
	CodexCodeDavinci001 = "code-davinci-001"

	GPT3TextDavinci003      = "text-davinci-003"
	GPT3TextDavinci002      = "text-davinci-002"
	GPT3TextCurie001        = "text-curie-001"
	GPT3TextBabbage001      = "text-babbage-001"
	GPT3TextAda001          = "text-ada-001"
	GPT3TextDavinci001      = "text-davinci-001"
	GPT3DavinciInstructBeta = "davinci-instruct-beta"
	GPT3Davinci             = "davinci"
	GPT3CurieInstructBeta   = "curie-instruct-beta"
	GPT3Curie               = "curie"
	GPT3Ada                 = "ada"
	GPT3Babbage             = "babbage"
)

func (t *HeraClientTestSuite) TestTokenCountApproximate() {
	bytes, err := os.ReadFile("./mocks/hera/tokenizer_example/example.txt")
	t.Require().Nil(err)
	tokenCount := t.HeraTestClient.GetTokenApproximate(string(bytes))
	t.Assert().Equal(61, tokenCount)
	// NOTE open gpt-3 https://beta.openai.com/tokenizer returns 64 tokens as the count
	// there's no opensource transformer for this, so use this + some margin when sending requests
}

func (t *HeraClientTestSuite) TestFilesUploadToCodeGen() {
	model := GPT3TextDavinci003
	// uses this model ^ if model param is empty str
	resp, err := t.HeraTestClient.UploadFiles(ctx, demoChartPath, model)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
	fmt.Println(resp)
}
