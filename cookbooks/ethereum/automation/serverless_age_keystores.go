package ethereum_automation_cookbook

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"

	"github.com/rs/zerolog/log"
	signing_automation_ethereum "github.com/zeus-fyi/zeus/pkg/artemis/signing_automation/ethereum"
	age_encryption "github.com/zeus-fyi/zeus/pkg/crypto/age"
	bls_signer "github.com/zeus-fyi/zeus/pkg/crypto/bls"
	"github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/memfs"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
)

var InMemFs = memfs.NewMemFs()

func GenerateValidatorDepositsAndCreateAgeEncryptedKeystores(ctx context.Context, w3Client signing_automation_ethereum.Web3SignerClient, vdg signing_automation_ethereum.ValidatorDepositGenerationParams, age age_encryption.Age, HDPassword string) error {
	err := vdg.GenerateAndEncryptValidatorKeysFromSeedAndPath(ctx)
	if err != nil {
		return err
	}
	dpSlice, err := w3Client.GenerateEphemeryDepositDataWithDefaultWd(ctx, vdg)
	if err != nil {
		return err
	}
	signing_automation_ethereum.PrintJSONSlice(vdg.Fp, dpSlice, vdg.Network)
	err = decryptToInMemFS(vdg.Fp.DirOut, HDPassword)
	if err != nil {
		return err
	}
	// these are the in memfs paths
	p := filepaths.Path{DirIn: "./keystores", DirOut: "./gzip", FnOut: "keystores.tar.gz"}
	b, err := gzipDirectoryToMemoryFS(p)
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
		acc, err := signing_automation_ethereum.DecryptKeystoreCipherIntoEthSignerBLS(context.Background(), input, hdPassword)
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
			return err

		}
	}
	return nil
}

func gzipDirectoryToMemoryFS(p filepaths.Path) ([]byte, error) {
	// Create in-memory file system and populate it with files from directoryPath
	// Create gzip writer on top of output file in memory
	gzipBytes := new(bytes.Buffer)
	gzipWriter := gzip.NewWriter(gzipBytes)

	// Create tar writer on top of gzip writer
	tarWriter := tar.NewWriter(gzipWriter)

	dir, err := InMemFs.Sub(p.DirIn)
	if err != nil {
		return nil, err
	}
	// Iterate over files in directory and add to tar archive
	err = fs.WalkDir(dir, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Create header for current file
		if d.IsDir() || path == p.FnOut {
			return nil
		}
		p.FnIn = path
		file, err := InMemFs.FS.Open(p.FileInPath())
		if err != nil {
			return err
		}
		defer file.Close()
		fileInfo, err := file.Stat()
		if err != nil {
			return err
		}

		header := new(tar.Header)
		header.Name = fileInfo.Name()
		header.Mode = int64(fileInfo.Mode())
		header.Size = fileInfo.Size()
		header.ModTime = fileInfo.ModTime()

		// Write header to tar archive
		err = tarWriter.WriteHeader(header)
		if err != nil {
			return err
		}

		// Copy file contents to tar archive
		_, err = io.Copy(tarWriter, file)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	// Flush and close tar writer, then gzip writer
	err = tarWriter.Close()
	if err != nil {
		return nil, err
	}
	err = gzipWriter.Close()
	if err != nil {
		return nil, err
	}

	// Return compressed bytes
	return gzipBytes.Bytes(), nil
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
