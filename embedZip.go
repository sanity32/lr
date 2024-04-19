package lr

import (
	"errors"

	"archive/zip"
	"bytes"
	"io"
	"os"
)

type EmbedZipFile []byte

var ErrZipFileIsEmpty = errors.New("zip file is empty")

func (EmbedZipFile) readZipFile(zf *zip.File) ([]byte, error) {
	f, err := zf.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return io.ReadAll(f)
}

func (e EmbedZipFile) ExtractFirst(filename string) (err error) {
	r, err := zip.NewReader(bytes.NewReader(e), int64(len(e)))
	if err != nil {
		return err
	}
	if len(r.File) == 0 {
		return ErrZipFileIsEmpty
	}

	if bb, err := e.readZipFile(r.File[0]); err != nil {
		return err
	} else {
		// filename = r.File[0].Name
		return os.WriteFile(filename, bb, 0777)
	}
}

func (e EmbedZipFile) Extract() error {
	r, err := zip.NewReader(bytes.NewReader(e), int64(len(e)))
	if err != nil {
		return err
	}
	for _, zippedFile := range r.File {
		unzippedBB, err := e.readZipFile(zippedFile)
		if err != nil {
			return err
		}
		if err := os.WriteFile(zippedFile.Name, unzippedBB, 0777); err != nil {
			return err
		}
	}
	return nil
}
