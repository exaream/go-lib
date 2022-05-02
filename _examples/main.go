package main

import (
	"fmt"

	"github.com/exaream/go-lib/security/nullbyte"
)

func main() {
	s := "\x00aaa\x00bbb\x00"

	fmt.Printf("Original    %q\n", s)
	if r, ok := nullbyte.Contains(s); ok {
		fmt.Printf("Contains()  %t\n", r)
	}

	if r, ok := nullbyte.Trim(s); ok {
		fmt.Printf("Trim()      %q\n", r)
		if r2, ok := nullbyte.Contains(r); ok {
			fmt.Printf("Contains()  %t\n", r2)
		}
	}

	if r, ok := nullbyte.RemoveAll(s); ok {
		fmt.Printf("RemoveAll() %q\n", r)
		if r2, ok := nullbyte.Contains(r); ok {
			fmt.Printf("Contains()  %t\n", r2)
		}
	}
}
