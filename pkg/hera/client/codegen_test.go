package hera_client

import "fmt"

const (
	CodexCodeDavinci002 = "code-davinci-002"
	CodexCodeCushman001 = "code-cushman-001"
	CodexCodeDavinci001 = "code-davinci-001"
)

func (t *HeraClientTestSuite) TestFilesUploadToCodeGen() {
	model := CodexCodeDavinci002
	// uses this model ^ if model param is empty str
	resp, err := t.HeraTestClient.UploadFiles(ctx, demoChartPath, model)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
	fmt.Println(resp)
}
