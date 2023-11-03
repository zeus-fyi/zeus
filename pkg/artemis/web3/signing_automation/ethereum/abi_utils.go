package signing_automation_ethereum

import (
	"context"
	"io"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/rs/zerolog/log"
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

func ReadAbiString(ctx context.Context, abiContents string) (*abi.ABI, error) {
	reader := strings.NewReader(abiContents)
	abiIn, err := abi.JSON(reader)
	if err != nil {
		log.Err(err).Msg("readAbi:  abi.JSON")
		return nil, err
	}
	return &abiIn, nil
}

func MustReadAbiString(ctx context.Context, abiContents string) *abi.ABI {
	reader := strings.NewReader(abiContents)
	abiIn, err := abi.JSON(reader)
	if err != nil {
		panic(err)
	}
	return &abiIn
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
