package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/ajiyoshi-vg/sqlgen/parse"
	//"github.com/drone/sqlgen/schema"
)

var (
	input    = flag.String("file", "", "input file name; required")
	output   = flag.String("o", "", "output file name; required")
	pkgName  = flag.String("pkg", "main", "output package name; required")
	typeName = flag.String("type", "", "type to generate; required")
	genFuncs = flag.Bool("funcs", true, "generate sql helper functions")
)

func main() {
	flag.Parse()

	// parses the syntax tree into something a bit
	// easier to work with.
	tree, err := parse.Parse(*input, *typeName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	// if the code is generated in a different folder
	// that the struct we need to import the struct
	if tree.Pkg != *pkgName && *pkgName != "main" {
		// TODO
	}

	var buf bytes.Buffer

	writePackage(&buf, *pkgName)
	if *genFuncs {
		writeImports(&buf, tree, "database/sql")
		writeRowFunc(&buf, tree)
		writeRowsFunc(&buf, tree)
		writeSliceFunc(&buf, tree)
	}

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
