package main

import (
	"path/filepath"

	"github.com/exaream/go-snippet/csvutil"
)

func main() {
	file := filepath.Join("..", "..", "testdata", "47OKINAW.CSV")
	csvutil.Read(file)
}
