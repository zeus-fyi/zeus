package ethereum_automation_cookbook

import (
	"archive/zip"
	"context"
	"encoding/json"
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/rs/zerolog/log"
	age_encryption "github.com/zeus-fyi/zeus/pkg/aegis/crypto/age"
	bls_signer "github.com/zeus-fyi/zeus/pkg/aegis/crypto/bls"
	signing_automation_ethereum2 "github.com/zeus-fyi/zeus/pkg/artemis/web3 /signing_automation/ethereum"
	"github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/compression"
	"github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/memfs"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
)

var InMemFs = memfs.NewMemFs()

func GenerateValidatorDepositsAndCreateAgeEncryptedKeystores(ctx context.Context, w3Client signing_automation_ethereum2.Web3SignerClient, vdg signing_automation_ethereum2.ValidatorDepositGenerationParams, age age_encryption.Age, HDPassword string) error {
	err := vdg.GenerateAndEncryptValidatorKeysFromSeedAndPath(ctx)
	if err != nil {
		return err
	}
	dpSlice, err := w3Client.GenerateEphemeryDepositDataWithDefaultWd(ctx, vdg)
	if err != nil {
		return err
	}
	signing_automation_ethereum2.PrintJSONSlice(vdg.Fp, dpSlice, vdg.Network)
	err = decryptToInMemFS(vdg.Fp.DirOut, HDPassword)
	if err != nil {
		return err
	}
	// these are the in memfs paths
	p := filepaths.Path{DirIn: "./keystores", DirOut: "./gzip", FnOut: "keystores.tar.gz"}
	b, err := compression.GzipDirectoryToMemoryFS(p, InMemFs)
	if err != nil {
		return err
	}
	err = InMemFs.MakeFileOut(&p, b)
	if err != nil {
		return err
	}
	p.FnIn = "keystores.tar.gz"
	p.DirIn = "./gzip"
	p.DirOut = vdg.Fp.DirIn
	err = age.EncryptFromInMemFS(InMemFs, &p)
	if err != nil {
		return err
	}
	vdg.Fp.FnIn = p.FnIn
	vdg.Fp.FnOut = "keystores"
	err = zipKeystoreFolder(vdg.Fp)
	if err != nil {
		return err
	}
	return nil
}

func decryptToInMemFS(directoryPath, hdPassword string) error {
	files, ferr := os.ReadDir(directoryPath)
	if ferr != nil {
		return ferr
	}
	filter := strings_filter.FilterOpts{
		DoesNotStartWithThese: []string{".DS_Store", "keystores.tar", "keystores.zip"},
		StartsWithThese:       nil,
		StartsWith:            "keystore",
		Contains:              "",
	}
	for _, file := range files {
		if file.IsDir() || !strings_filter.FilterStringWithOpts(file.Name(), &filter) {
			continue
		}
		fp := filepath.Join(directoryPath, file.Name())
		jsonByteArray, err := os.ReadFile(fp)
		if err != nil {
			log.Err(err)
			panic(err)
		}
		input := make(map[string]interface{})
		err = json.Unmarshal(jsonByteArray, &input)
		if err != nil {
			log.Err(err)
			panic(err)
		}
		acc, err := signing_automation_ethereum2.DecryptKeystoreCipherIntoEthSignerBLS(context.Background(), input, hdPassword)
		if err != nil {
			log.Err(err)
			panic(err)
		}
		p := filepaths.Path{
			DirIn: "./keystores",
			FnIn:  strings_filter.AddHexPrefix(acc.PublicKeyString()),
		}
		err = InMemFs.MakeFileIn(&p, []byte(bls_signer.ConvertBytesToString(acc.BLSPrivateKey.Marshal())))
		if err != nil {
			log.Err(err)
			panic(err)
		}
	}
	return nil
}

// zipKeystoreFolder zips the keystore folder which contains keystore.tar.gz.age
func zipKeystoreFolder(p filepaths.Path) error {
	// Open the file for reading
	fileToZip, err := os.Open(p.FileInPath())
	if err != nil {
		return err
	}
	defer fileToZip.Close()
	// Create the zip file
	zipFileName := p.FnOut + ".zip"
	generationPath := filepaths.Path{
		DirOut:      p.DirOut,
		FnOut:       zipFileName,
		Env:         "",
		FilterFiles: nil,
	}
	zipFile, err := os.Create(generationPath.FileOutPath())
	if err != nil {
		return err
	}
	defer zipFile.Close()

	// Create a new zip archive
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Add the file to the zip archive
	fileInfo, err := fileToZip.Stat()
	if err != nil {
		return err
	}
	header, err := zip.FileInfoHeader(fileInfo)
	if err != nil {
		return err
	}
	header.Method = zip.Deflate
	header.Name = path.Join("keystores", header.Name)
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)
	if err != nil {
		return err
	}
	return nil
}
