package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	gocsv "github.com/exaream/go-snippet/gocsv"
)

func main() {
	csvPath := filepath.Join("..", "..", "testdata", "KEN_ALL.CSV")
	if err := gocsv.Read(csvPath); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	time.Sleep(2 * time.Second)

	zipPath := filepath.Join("..", "..", "testdata", "ken_all.zip")
	if err := gocsv.WalkInZip(zipPath); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
