package csv

import (
	"archive/zip"
	"path/filepath"
	"strings"
)

func ReadZip(target string) (err error) {
	z, err := zip.OpenReader(target)
	if err != nil {
		return err
	}

	defer func() {
		err = z.Close()
	}()

	for _, f := range z.File {
		if strings.ToLower(filepath.Ext(f.Name)) != ext {
			continue
		}

		if err := readZip(f); err != nil {
			return err
		}
	}

	return err
}

func readZip(zf *zip.File) error {
	f, err := zf.Open()
	if err != nil {
		return err
	}

	defer func() {
		err = f.Close()
	}()

	return read(f)
}
