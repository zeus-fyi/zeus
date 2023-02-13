package serverless_inmemdb

import (
	"context"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	aegis_inmemdbs "github.com/zeus-fyi/zeus/pkg/aegis/inmemdbs"
	age_encryption "github.com/zeus-fyi/zeus/pkg/crypto/age"
	bls_signer "github.com/zeus-fyi/zeus/pkg/crypto/bls"
	"github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/memfs"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
)

var (
	fl           = strings_filter.FilterOpts{StartsWith: "keystore"}
	KeystorePath = filepaths.Path{
		PackageName: "",
		DirIn:       "/opt/keystores",
		FnIn:        "keystores.tar.gz.age",
		FnOut:       "",
		Env:         "",
		FilterFiles: &fl,
	}
	InMemFs  = memfs.MemFS{}
	unzipDir = "./keystores"
)

func ImportIntoInMemFs(ctx context.Context, enc age_encryption.Age) error {
	if InMemFs.FS == nil {
		InMemFs = memfs.NewMemFs()
	} else {
		return nil
	}
	err := enc.DecryptAndUnGzipToInMemFs(&KeystorePath, InMemFs, unzipDir)
	if err != nil {
		log.Ctx(ctx).Err(err)
		return err
	}
	InMemFs.Remove("keystores.tar.gz.age")
	InMemFs.Remove("keystores.tar.gz")
	return nil
}

// SignValidatorMessagesFromInMemFs returns a 0x prefixed hex string of the signature
func SignValidatorMessagesFromInMemFs(ctx context.Context, signReqs aegis_inmemdbs.EthereumBLSKeySignatureRequests) (aegis_inmemdbs.EthereumBLSKeySignatureResponses, error) {
	resp := make(map[string]aegis_inmemdbs.EthereumBLSKeySignatureResponse)
	batchResp := aegis_inmemdbs.EthereumBLSKeySignatureResponses{
		Map: resp,
	}
	if len(signReqs.Map) == 0 {
		return batchResp, nil
	}

	KeystorePath.DirIn = unzipDir
	for pubkey, req := range signReqs.Map {
		KeystorePath.FnIn = pubkey
		b, err := InMemFs.ReadFile(KeystorePath.FileInPath())
		if err != nil {
			err = errors.New(fmt.Sprintf("could not read key file %s from inmemfs: "+err.Error(), pubkey))
			log.Ctx(ctx).Err(err)
		} else {
			acc := bls_signer.NewEthSignerBLSFromExistingKey(string(b))
			sig := acc.Sign([]byte(req.Message)).Marshal()
			batchResp.Map[pubkey] = aegis_inmemdbs.EthereumBLSKeySignatureResponse{Signature: "0x" + bls_signer.ConvertBytesToString(sig)}
		}
	}
	if len(batchResp.Map) != len(signReqs.Map) {
		log.Ctx(ctx).Warn().Msg("SignValidatorMessagesFromInMemFs, did not contain all expected validator signatures")
	}
	return batchResp, nil
}
