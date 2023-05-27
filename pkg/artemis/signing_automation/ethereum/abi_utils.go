package signing_automation_ethereum

import (
	"context"
	"io"
	"os"
	"path"
	"runtime"

	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/gochain/v4/accounts/abi"
)

func ABIOpenFile(ctx context.Context, abiFile string) (*abi.ABI, error) {
	jsonReader, err := os.Open(abiFile)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("GetABI: ABIOpenFile")
		return nil, err
	}
	return ReadAbi(ctx, jsonReader)
}

func ReadAbi(ctx context.Context, reader io.Reader) (*abi.ABI, error) {
	abiIn, err := abi.JSON(reader)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("readAbi:  abi.JSON")
		return nil, err
	}
	return &abiIn, nil
}

func ForceDirToEthSigningDirLocation() string {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "")
	err := os.Chdir(dir)
	if err != nil {
		panic(err.Error())
	}
	return dir
}
