package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"csv"
)

var testdataDir = filepath.Join("..", "..", "testdata")

func main() {
	csvPath := filepath.Join(testdataDir, "13TOKYO.CSV")
	if err := csv.Read(csvPath); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	time.Sleep(2 * time.Second)

	zipPath := filepath.Join(testdataDir, "13tokyo.zip")
	if err := csv.ReadZip(zipPath); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
