package serverless_inmemdb

import (
	"context"
	"github.com/rs/zerolog/log"
	aegis_inmemdbs "github.com/zeus-fyi/zeus/pkg/aegis/inmemdbs"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
)

var (
	fl           = strings_filter.FilterOpts{StartsWith: "deposit"}
	KeystorePath = filepaths.Path{
		PackageName: "",
		DirIn:       "/opt/keystores",
		DirOut:      "/opt/keystores",
		FnIn:        "",
		FnOut:       "",
		Env:         "",
		FilterFiles: &fl,
	}
)

// ImportIntoInMemDB assumes you're using the same password for your keystores.
func ImportIntoInMemDB(ctx context.Context, password string) {
	if aegis_inmemdbs.ValidatorInMemDB == nil {
		aegis_inmemdbs.InitValidatorDB()
	}
	vs := aegis_inmemdbs.DecryptedValidators{HDPassword: password}
	err := KeystorePath.WalkAndApplyFuncToFileType(".json", vs.ReadValidatorFromKeystore)
	if err != nil {
		log.Ctx(ctx).Err(err)
		panic(err)
	}
	aegis_inmemdbs.InsertValidatorsInMemDb(ctx, vs.Validators)
}
