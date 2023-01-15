package ethereum_automation_cookbook

import (
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
)

var (
	fl           = strings_filter.FilterOpts{StartsWith: "deposit"}
	KeystorePath = filepaths.Path{
		PackageName: "",
		DirIn:       "./ethereum/automation/validator_keys/ephemery",
		DirOut:      "./ethereum/automation/validator_keys/ephemery",
		FnIn:        "keystore-ephemery-m_12381_3600_0_0_0.json",
		FnOut:       "",
		Env:         "",
		FilterFiles: &fl,
	}
)
