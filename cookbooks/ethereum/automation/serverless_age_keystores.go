package ethereum_automation_cookbook

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	signing_automation_ethereum "github.com/zeus-fyi/zeus/pkg/artemis/signing_automation/ethereum"
	age_encryption "github.com/zeus-fyi/zeus/pkg/crypto/age"
	bls_signer "github.com/zeus-fyi/zeus/pkg/crypto/bls"
	"github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/memfs"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
)

var InMemFs = memfs.NewMemFs()

func GenerateAgeKeystores(keystoresPath filepaths.Path, age age_encryption.Age, HDPassword string) error {
	err := decryptToInMemFS(keystoresPath.DirIn, HDPassword)
	if err != nil {
		return err
	}
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
	p.DirOut = keystoresPath.DirOut
	err = age.EncryptFromInMemFS(InMemFs, &p)
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

	for _, file := range files {
		if file.IsDir() || file.Name() == ".DS_Store" {
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
