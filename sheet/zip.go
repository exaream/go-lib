package sheet

import (
	"archive/zip"
	"fmt"
)

const zipExt = ".zip"

func WalkCSVInZip(target string) (rerr error) {
	if ext := fileExt(target); ext != zipExt {
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
		if fileExt(f.Name) != csvExt {
			continue
		}

		if err := readCSVInZip(f); err != nil {
			return err
		}
	}

	return err
}

func readCSVInZip(zf *zip.File) (rerr error) {
	f, err := zf.Open()
	if err != nil {
		return err
	}

	defer func() {
		if err = f.Close(); rerr == nil && err != nil {
			rerr = err
		}
	}()

	return readCSV(f)
}
