package zip

import (
	"archive/zip"
	"errors"
	"io"
	"strings"
)

func ExtractXML(zipPath string) (io.Reader, error) {
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

			return rc, nil
		}
	}
	
	defer r.Close()
	
	return nil, errors.New("no xml found")
}
