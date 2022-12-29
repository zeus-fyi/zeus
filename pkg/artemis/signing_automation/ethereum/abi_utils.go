package signing_automation_ethereum

import (
	"io"
	"os"
	"path"
	"runtime"

	"github.com/gochain/gochain/v4/accounts/abi"
	"github.com/rs/zerolog/log"
)

func ABIOpenFile(abiFile string) (*abi.ABI, error) {
	jsonReader, err := os.Open(abiFile)
	if err != nil {
		log.Err(err).Msg("GetABI: ABIOpenFile")
		return nil, err
	}
	return readAbi(jsonReader)
}

func readAbi(reader io.Reader) (*abi.ABI, error) {
	abiIn, err := abi.JSON(reader)
	if err != nil {
		log.Err(err).Msg("readAbi:  abi.JSON")
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
