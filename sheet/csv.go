package sheet

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

const csvExt = ".csv"

type reader io.Reader

func ReadCSV(target string) (rerr error) {
	f, err := os.Open(path.Clean(target))
	if err != nil {
		return err
	}

	defer func() {
		if err = f.Close(); rerr == nil && err != nil {
			rerr = err
		}
	}()

	if ext(f.Name()) != csvExt {
		return fmt.Errorf("%s is not a CSV file", f.Name())
	}

	return readCSV(f)
}

func fileExt(filePath string) string {
	return strings.ToLower(filepath.Ext(filePath))
}

func readCSV(r reader) error {
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
