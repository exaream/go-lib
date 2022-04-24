package gocsv

import (
	"archive/zip"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

const ext = ".csv"

type reader io.Reader

func Read(target string) (rerr error) {
	f, err := os.Open(path.Clean(target))
	if err != nil {
		return err
	}

	defer func() {
		if err = f.Close(); rerr == nil && err != nil {
			rerr = err
		}
	}()

	if strings.ToLower(filepath.Ext(f.Name())) != ext {
		return fmt.Errorf("%s is not a CSV file", f.Name)
	}

	return read(f)
}

func read(r reader) error {
	cr := csv.NewReader(transform.NewReader(r, japanese.ShiftJIS.NewDecoder()))

	for {
		record, err := cr.Read()

		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return err
		}

		fmt.Println(record)
	}

	return nil
}

func WalkInZip(target string) (rerr error) {
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
		if strings.ToLower(filepath.Ext(f.Name)) != ext {
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
