package csv

import (
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

type Reader io.Reader

func Read(target string) (err error) {
	f, err := os.Open(path.Clean(target))
	if err != nil {
		return err
	}

	defer func() {
		err = f.Close()
	}()

	if strings.ToLower(filepath.Ext(f.Name())) != ext {
		return fmt.Errorf("%s is not a CSV file", f.Name)
	}

	return read(f)
}

func read(r Reader) (err error) {
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

	return err
}
