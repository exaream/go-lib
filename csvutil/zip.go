package csvutil

import (
	"archive/zip"
	"fmt"

	"github.com/exaream/go-snippet/fileutil"
)

const zipExt = ".zip"

func WalkInZip(target string) (rerr error) {
	if ext := fileutil.Ext(target); ext != zipExt {
		return fmt.Errorf("invalid file extension: %s", ext)
	}

	z, err := zip.OpenReader(target)
	if err != nil {
		return err
	}

	defer func() {
		if err = z.Close(); rerr == nil && err != nil {
			rerr = err
		}
	}()

	for _, f := range z.File {
		if fileutil.Ext(f.Name) != csvExt {
			continue
		}

		if err := readInZip(f); err != nil {
			return err
		}
	}

	return err
}

func readInZip(zf *zip.File) (rerr error) {
	f, err := zf.Open()
	if err != nil {
		return err
	}

	defer func() {
		if err = f.Close(); rerr == nil && err != nil {
			rerr = err
		}
	}()

	return read(f)
}
