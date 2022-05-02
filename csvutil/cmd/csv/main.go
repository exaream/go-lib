package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/exaream/go-lib/csvutil"
)

func main() {
	file := filepath.Join("..", "..", "testdata", "47OKINAW.CSV")
	if err := csvutil.Read(file); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
