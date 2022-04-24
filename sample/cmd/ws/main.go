package main

import (
	"path/filepath"

	"github.com/exaream/go-snippet/gocsv"
)

func main() {
	target := filepath.Join("..", "..", "testdata", "47OKINAW.CSV")
	gocsv.Read(target)
}
