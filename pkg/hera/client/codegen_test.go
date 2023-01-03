package hera_client

import (
	"fmt"

	gogpt "github.com/sashabaranov/go-gpt3"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	test_base "github.com/zeus-fyi/zeus/test"
)

var maxTokensByModel = map[string]int{
	gogpt.GPT3TextDavinci003: 2048,
	gogpt.GPT3TextDavinci002: 2048,
	gogpt.GPT3TextDavinci001: 2048,
}

var demoChartPath = filepaths.Path{
	PackageName: "",
	DirIn:       "./mocks/hera/prompt_example",
	DirOut:      "./mocks/outputs/hera",
	FnIn:        "prompt", // filename for your gzip workload
	FnOut:       "",
	Env:         "",
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
