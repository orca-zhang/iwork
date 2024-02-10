package main

import (
	"fmt"
	"os"

	"github.com/orcastor/iwork-converter/iwork2html"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf(`Converts pages files to html

Usage:
    %s infile.pages outfile.html

`, os.Args[0])
		return
	}

	if err := iwork2html.Convert(os.Args[1], os.Args[2]); err != nil {
		panic(err)
	}
}
