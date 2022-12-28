package hera_client

import (
	"fmt"
	"os"

	gogpt "github.com/sashabaranov/go-gpt3"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	test_base "github.com/zeus-fyi/zeus/test"
)

var maxTokensByModel = map[string]int{
	gogpt.GPT3TextDavinci003: 2048,
	gogpt.GPT3TextDavinci002: 2048,
	gogpt.GPT3TextDavinci001: 2048,
}

func (t *HeraClientTestSuite) TestTokenCountApproximate() {
	bytes, err := os.ReadFile("./mocks/hera/tokenizer_example/example.txt")
	t.Require().Nil(err)
	tokenCount := t.HeraTestClient.GetTokenApproximate(string(bytes))
	t.Assert().Equal(61, tokenCount)
	// NOTE open gpt-3 https://beta.openai.com/tokenizer returns 64 tokens as the count
	// there's no opensource transformer for this, so use this + some margin when sending requests
	// 2048 is the max token count for most models, the max size - prompt size, is your limitation on completion
	// tokens
}

var demoChartPath = filepaths.Path{
	PackageName: "",
	DirIn:       "./mocks/hera/custom_make",
	DirOut:      "./mocks/outputs/hera",
	FnIn:        "prompt", // filename for your gzip workload
	FnOut:       "",
	Env:         "",
	FilterFiles: strings_filter.FilterOpts{},
}

func (t *HeraClientTestSuite) TestFilesUploadToCodeGen() {
	model := gogpt.GPT3TextDavinci003
	// uses this model ^ if model param is empty str
	promptTokenEst := t.HeraTestClient.GetTokenApproximate(demoChartPath.FileInPath())
	tokenSizeMax, _ := maxTokensByModel[model]
	maxTokens := tokenSizeMax - promptTokenEst

	// the python func changes the directory, so set it back to the test path
	test_base.ForceDirToTestDirLocation()
	resp, err := t.HeraTestClient.UploadFiles(ctx, demoChartPath, model, fmt.Sprintf("%d", maxTokens))
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
	fmt.Println(resp)

	demoChartPath.FnOut = "openai.txt"
	for _, choice := range resp.Choices {
		werr := demoChartPath.WriteToFileOutPath([]byte(choice.Text))
		t.Require().Nil(werr)
	}

}
