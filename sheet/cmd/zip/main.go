package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/exaream/go-snippet/sheet"
)

func main() {
	file := filepath.Join("..", "..", "testdata", "47okinaw.zip")
	if err := sheet.WalkCSVInZip(file); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
