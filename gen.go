package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/ajiyoshi-vg/sqlgen/parse"
)

var (
	input   = flag.String("file", "", "input file name; required")
	output  = flag.String("o", "", "output file name; required")
	pkgName = flag.String("pkg", "main", "output package name; required")
)

func main() {
	flag.Parse()

	root, err := parse.ParseAll(*input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	// if the code is generated in a different folder
	// that the struct we need to import the struct
	if root.Pkg != *pkgName && *pkgName != "main" {
		// TODO
	}

	var buf bytes.Buffer
	writePackage(&buf, root.Pkg)
	writeImports(&buf, root, "database/sql")

	root.Each(func(n *parse.Node) {
		writeRowFunc(&buf, n)
		writeRowsFunc(&buf, n)
		writeSliceFunc(&buf, n)
	})

	// formats the generated file using gofmt
	pretty, err := format(&buf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	// create output source for file. defaults to
	// stdout but may be file.
	var out io.WriteCloser = os.Stdout
	if *output != "" {
		out, err = os.Create(*output)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return
		}
		defer out.Close()
	}

	io.Copy(out, pretty)
}
