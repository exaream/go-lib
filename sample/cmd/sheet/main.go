package main

import (
	"path/filepath"

	"github.com/exaream/go-snippet/sheet"
)

func main() {
	file := filepath.Join("..", "..", "testdata", "47OKINAW.CSV")
	sheet.ReadCSV(file)
}
