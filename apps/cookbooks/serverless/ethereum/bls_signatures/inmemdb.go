package main

import (
	"context"
	"github.com/hashicorp/go-memdb"
	"github.com/rs/zerolog/log"
	aegis_inmemdbs "github.com/zeus-fyi/zeus/pkg/aegis/inmemdbs"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	"path"
)

var (
	fl           = strings_filter.FilterOpts{StartsWith: "deposit"}
	KeystorePath = filepaths.Path{
		PackageName: "",
		DirIn:       "./keystores",
		DirOut:      "./keystores",
		FnIn:        "",
		FnOut:       "",
		Env:         "",
		FilterFiles: &fl,
	}
	InMemDB *memdb.MemDB
)

// TODO check if inmemdb is empty and regenerate if true

// ImportIntoInMemDB assumes you're using the same password for your keystores, it
// also implies usage of network in your path directory. It will reset the in-memory
// database and import all the validators from the keystore directory when called.
func ImportIntoInMemDB(ctx context.Context, network, password string) {
	aegis_inmemdbs.InitValidatorDB()
	fl.Contains = network
	KeystorePath.DirIn = path.Join(KeystorePath.DirIn, network)
	vs := aegis_inmemdbs.DecryptedValidators{HDPassword: password}
	err := KeystorePath.WalkAndApplyFuncToFileType(".json", vs.ReadValidatorFromKeystore)
	if err != nil {
		log.Ctx(ctx).Err(err)
		panic(err)
	}
	aegis_inmemdbs.InsertValidatorsInMemDb(ctx, vs.Validators)
	InMemDB = aegis_inmemdbs.ValidatorInMemDB
}
