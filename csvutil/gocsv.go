package csvutil

import (
	"fmt"
	"os"
	"path"

	"github.com/exaream/go-lib/fileutil"
	"github.com/gocarina/gocsv"
)

// Create creates a structured CSV file.
func Create[T comparable](target string, records []T) (rerr error) {
	f, err := os.Create(target)
	if err != nil {
		return err
	}

	defer func() {
		if err = f.Close(); rerr == nil && err != nil {
			rerr = err
		}
	}()

	if err := gocsv.MarshalFile(&records, f); err != nil {
		return err
	}

	return nil
}

// ReadStructured output a structured CSV file.
func ReadStructured[T comparable](target string, records []T) (rerr error) {
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

	if err := gocsv.UnmarshalFile(f, &records); err != nil {
		return err
	}

	for _, v := range records {
		fmt.Printf("%+v\n", v)
	}

	return nil
}
