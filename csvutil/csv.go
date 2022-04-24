package csvutil

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/exaream/go-snippet/fileutil"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

const csvExt = ".csv"

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

	if fileutil.Ext(f.Name()) != csvExt {
		return fmt.Errorf("%s is not a CSV file", f.Name())
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
