package v1_common_routes

import filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"

var CommonManager ClientManager

type ClientManager struct {
	BucketURL string
	filepaths.Path
}
