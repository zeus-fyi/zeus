package beacon_cookbooks

import filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"

func (t *BeaconCookbookTestSuite) TestBeaconConfigDriver() {
	p := filepaths.Path{}
	err := SetCustomConfigOpts(&p)
	t.Require().Nil(err)
}
