package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/acsellers/inflections"
	"github.com/ajiyoshi-vg/scannergen/parse"
)

// WritePackage writes the Go package header to
// writer w with the given package name.
func writePackage(w io.Writer, name string) {
	fmt.Fprintf(w, sPackage, name)
}

func writeImports(w io.Writer, tree *parse.Node, pkgs ...string) {
	var pmap = map[string]struct{}{}

	// add default packages
	for _, pkg := range pkgs {
		pmap[pkg] = struct{}{}
	}

	// check each edge node to see if it is
	// encoded, which might require us to import
	// other packages
	for _, node := range tree.Edges() {
		if node.Tags == nil || len(node.Tags.Encode) == 0 {
			continue
		}
		switch node.Tags.Encode {
		case "json":
			pmap["encoding/json"] = struct{}{}
			// case "gzip":
			// 	pmap["compress/gzip"] = struct{}{}
			// case "snappy":
			// 	pmap["github.com/golang/snappy"] = struct{}{}
		}
	}

	if len(pmap) == 0 {
		return
	}

	// write the import block, including each
	// encoder package that was specified.
	fmt.Fprintln(w, "\nimport (")
	for pkg, _ := range pmap {
		fmt.Fprintf(w, "\t%q\n", pkg)
	}
	fmt.Fprintln(w, ")")
}

func writeSliceFunc(w io.Writer, tree *parse.Node) {

	var buf1, buf2, buf3 bytes.Buffer

	var i, depth int
	var parent = tree

	for _, node := range tree.Edges() {
		if node.Tags.Skip {
			continue
		}

		// temporary variable declaration
		switch node.Kind {
		case parse.Map, parse.Slice:
			fmt.Fprintf(&buf1, "var v%d %s\n", i, "[]byte")
		default:
			fmt.Fprintf(&buf1, "var v%d %s\n", i, node.Type)
		}

		// variable scanning
		fmt.Fprintf(&buf3, "v%d,\n", i)

		// variable setting
		path := node.Path()[2:]

		// if the parent is a ptr struct we
		// need to create a new
		if parent != node.Parent && node.Parent.Kind == parse.Ptr {
			// if node.Parent != nil && node.Parent.Parent != parent {
			// 	fmt.Fprintln(&buf2, "}\n")
			// 	depth--
			// }

			// seriously ... this works?
			if node.Parent != nil && node.Parent.Parent != parent {
				for _, p := range path {
					if p == parent || depth == 0 {
						break
					}
					fmt.Fprintln(&buf2, "}\n")
					depth--
				}
			}
			depth++
			fmt.Fprintf(&buf2, "if v.%s != nil {\n", join(path[:len(path)-1], "."))
		}

		switch node.Kind {
		case parse.Map, parse.Slice, parse.Struct, parse.Ptr:
			fmt.Fprintf(&buf2, "v%d, _ = json.Marshal(&v.%s)\n", i, join(path, "."))
		default:
			fmt.Fprintf(&buf2, "v%d=v.%s\n", i, join(path, "."))
		}

		parent = node.Parent
		i++
	}

	for depth != 0 {
		depth--
		fmt.Fprintln(&buf2, "}\n")
	}

	fmt.Fprintf(w,
		sSliceRow,
		tree.Type,
		tree.Type,
		buf1.String(),
		buf2.String(),
		buf3.String(),
	)
}

func writeRowFunc(w io.Writer, tree *parse.Node) {

	var buf1, buf2, buf3 bytes.Buffer

	var i int
	var parent = tree
	for _, node := range tree.Edges() {
		if node.Tags.Skip {
			continue
		}

		// temporary variable declaration
		switch node.Kind {
		case parse.Map, parse.Slice:
			fmt.Fprintf(&buf1, "var v%d %s\n", i, "[]byte")
		default:
			fmt.Fprintf(&buf1, "var v%d %s\n", i, node.Type)
		}

		// variable scanning
		fmt.Fprintf(&buf2, "&v%d,\n", i)

		// variable setting
		path := node.Path()[2:]

		// if the parent is a ptr struct we
		// need to create a new
		if parent != node.Parent && node.Parent.Kind == parse.Ptr {
			fmt.Fprintf(&buf3, "v.%s=&%s{}\n", join(path[:len(path)-1], "."), node.Parent.Type)
		}

		switch node.Kind {
		case parse.Map, parse.Slice, parse.Struct, parse.Ptr:
			fmt.Fprintf(&buf3, "json.Unmarshal(v%d, &v.%s)\n", i, join(path, "."))
		default:
			fmt.Fprintf(&buf3, "v.%s=v%d\n", join(path, "."), i)
		}

		parent = node.Parent
		i++
	}

	fmt.Fprintf(w,
		sScanRow,
		tree.Type,
		tree.Type,
		buf1.String(),
		buf2.String(),
		tree.Type,
		buf3.String(),
	)
}

func writeRowsFunc(w io.Writer, tree *parse.Node) {
	var buf1, buf2, buf3 bytes.Buffer

	var i int
	var parent = tree
	for _, node := range tree.Edges() {
		if node.Tags.Skip {
			continue
		}

		// temporary variable declaration
		switch node.Kind {
		case parse.Map, parse.Slice:
			fmt.Fprintf(&buf1, "var v%d %s\n", i, "[]byte")
		default:
			fmt.Fprintf(&buf1, "var v%d %s\n", i, node.Type)
		}

		// variable scanning
		fmt.Fprintf(&buf2, "&v%d,\n", i)

		// variable setting
		path := node.Path()[2:]

		// if the parent is a ptr struct we
		// need to create a new
		if parent != node.Parent && node.Parent.Kind == parse.Ptr {
			fmt.Fprintf(&buf3, "v.%s=&%s{}\n", join(path[:len(path)-1], "."), node.Parent.Type)
		}

		switch node.Kind {
		case parse.Map, parse.Slice, parse.Struct, parse.Ptr:
			fmt.Fprintf(&buf3, "json.Unmarshal(v%d, &v.%s)\n", i, join(path, "."))
		default:
			fmt.Fprintf(&buf3, "v.%s=v%d\n", join(path, "."), i)
		}

		parent = node.Parent
		i++
	}

	fmt.Fprintf(w,
		sScanRows,
		inflections.Pluralize(tree.Type),
		tree.Type,
		tree.Type,
		buf1.String(),
		buf2.String(),
		tree.Type,
		buf3.String(),
	)
}

// join is a helper function that joins nodes
// together by name using the seperator.
func join(nodes []*parse.Node, sep string) string {
	var parts []string
	for _, node := range nodes {
		parts = append(parts, node.Name)
	}
	return strings.Join(parts, sep)
}
