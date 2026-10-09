package extract

import (
	"archive/zip"
	"bytes"
	"compress/gzip"
	"errors"
	"io"
	"os"
	"strings"
)

func ExtractZip(zipPath string) (io.Reader, error) {
	var buf bytes.Buffer
	// open the zip with zip.OpenReader
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return nil, err
	}


	// loop over files inside
	for _, file := range r.File {
		if strings.HasSuffix(file.Name, ".xml") {
			rc, err := file.Open()
			if err != nil {
				return nil, err
			}

			io.Copy(&buf, rc)
			rc.Close()

			return &buf, nil
		}
	}
	r.Close()
	
	return nil, errors.New("no xml file found")
}

func ExtractGz(gzPath string) (io.Reader, error) {
	var buf bytes.Buffer
	file, err := os.Open(gzPath)
	if err != nil {
		return nil, err
	}

	r, err := gzip.NewReader(file)
	if err != nil {
		return nil, err
	}

	io.Copy(&buf, r)
	r.Close()
	file.Close()

	return &buf, nil
}
