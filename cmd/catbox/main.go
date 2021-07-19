package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/wabarc/go-catbox"
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:\n\n")
		fmt.Fprintf(os.Stderr, "  catbox [options] [file1] ... [fileN]\n")

		flag.PrintDefaults()
	}
	var basePrint = func() {
		fmt.Print("A CLI tool help upload files to Catbox.\n\n")
		flag.Usage()
		fmt.Fprint(os.Stderr, "\n")
	}

	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		basePrint()
		os.Exit(0)
	}

}

func main() {
	files := flag.Args()

	cat := catbox.New(nil)
	for _, path := range files {
		if url, err := cat.Upload(path); err != nil {
			fmt.Fprintf(os.Stderr, "catbox: %v\n", err)
		} else {
			fmt.Fprintf(os.Stdout, "%s  %s\n", url, path)
		}
	}
}
