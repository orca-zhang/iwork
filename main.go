package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/orcastor/iwork-converter/iwork2html"
	"github.com/orcastor/iwork-converter/iwork2text"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf(`Converts pages files to html/json/txt

Usage:
    %s infile.pages outfile.html

`, os.Args[0])
		return
	}

	switch {
	case strings.HasSuffix(os.Args[2], ".txt"):
		if err := iwork2text.Convert(os.Args[1], os.Args[2]); err != nil {
			panic(err)
		}
	default:
		if err := iwork2html.Convert(os.Args[1], os.Args[2]); err != nil {
			panic(err)
		}
	}
}
