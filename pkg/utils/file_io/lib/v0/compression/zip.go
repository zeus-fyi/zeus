package compression

import (
	"archive/zip"
	"bytes"
	"io"
	"path"

	"github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/memfs"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

// ZipKeystoreFileInMemory zips the keystroes/keystore.tar.gzip.age to a single keystores.zip
func ZipKeystoreFileInMemory(p filepaths.Path, inMemFs memfs.MemFS) (*bytes.Buffer, error) {
	// Open the file for reading
	fileToZip, err := inMemFs.Open(p.FileInPath())
	if err != nil {
		return nil, err
	}
	defer fileToZip.Close()

	// Create the zip file
	zipFile := new(bytes.Buffer)
	// Create a new zip archive
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Add the file to the zip archive
	fileInfo, err := fileToZip.Stat()
	if err != nil {
		return nil, err
	}
	header, err := zip.FileInfoHeader(fileInfo)
	if err != nil {
		return nil, err
	}
	header.Method = zip.Deflate
	header.Name = path.Join("keystores", header.Name)
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(writer, fileToZip)
	if err != nil {
		return nil, err
	}
	err = zipWriter.Close()
	if err != nil {
		return nil, err
	}
	return zipFile, nil
}
