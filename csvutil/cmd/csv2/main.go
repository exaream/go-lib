package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/exaream/go-snippet/csvutil"
	"github.com/exaream/go-snippet/fileutil"
)

type Country struct {
	Name       string `csv:"国名"`
	ISOCode    string `csv:"ISOコード"`
	Population int    `csv:"人口"`
}

func main() {
	target := filepath.Join("..", "..", "testdata", "country.csv")
	countries := []Country{
		{Name: "中国", ISOCode: "CN/CHN", Population: 1444200000},
		{Name: "インド", ISOCode: "IN/IND", Population: 1393400000},
		{Name: "アメリカ", ISOCode: "US/USA", Population: 332900000},
		{Name: "日本", ISOCode: "JP/JPN", Population: 126100000},
	}

	if !fileutil.Exist(target) {
		if err := csvutil.Create(target, countries); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	// TODO: Check the difference with and without *
	records := make([]*Country, cap(countries))
	if err := csvutil.ReadStructured(target, records); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// TODO: Remove the following
	// for _, c := range countries {
	// 	fmt.Println(c.Name)
	// }
}
