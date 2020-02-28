// The avrogo command generates Go types for the Avro schemas specified on the
// command line. Each schema file results in a Go file with the same basename but with a ".go" suffix.
//
// Usage:
//
//	usage: avrogo [flags] schema-file...
//	  -d string
//	    	directory to write Go files to (default ".")
//	  -p string
//	    	package name (defaults to $GOPACKAGE)
//	  -t	generated files will have _test.go suffix
//	  -map string
//	    	map from Avro namespace to Go package.
//
// By default, a type is generated for each Avro definition
// in the schema. Some additional metadata fields are
// recognized:
//
// - If a definition has a "go.package" metadata
// field, the type from that package will be used instead.
// - If a definition has a "go.name" metadata field,
// the associated string will be used for the Go type name.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Generate the tests.

//go:generate go run ./generatetestcode.go

var (
	dirFlag  = flag.String("d", ".", "directory to write Go files to")
	pkgFlag  = flag.String("p", os.Getenv("GOPACKAGE"), "package name (defaults to $GOPACKAGE)")
	testFlag = flag.Bool("t", strings.HasSuffix(os.Getenv("GOFILE"), "_test.go"), "generated files will have _test.go suffix (defaults to true if $GOFILE is a test file)")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: avrogo [flags] schema-file...\n")
		flag.PrintDefaults()
		os.Exit(2)
	}
	flag.Parse()
	files := flag.Args()
	if len(files) == 0 {
		flag.Usage()
	}
	if *pkgFlag == "" {
		fmt.Fprintf(os.Stderr, "avrogo: -p flag must specify a package name or set $GOPACKAGE\n")
		os.Exit(1)
	}
	if err := generateFiles(files); err != nil {
		fmt.Fprintf(os.Stderr, "avrogo: %v\n", err)
		os.Exit(1)
	}
}

func generateFiles(files []string) error {
	for _, f := range files {
		if err := generateFile(f); err != nil {
			return err
		}
	}
	return nil
}

func generateFile(f string) error {
	data, err := ioutil.ReadFile(f)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	if err := generate(&buf, data, *pkgFlag); err != nil {
		return err
	}
	resultData, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Printf("%s\n", buf.Bytes())
		return fmt.Errorf("cannot format source: %v", err)
	}

	outFile := filepath.Base(f)
	outFile = strings.TrimSuffix(outFile, filepath.Ext(f)) + "_gen"
	if *testFlag {
		outFile += "_test"
	}
	outFile += ".go"
	outFile = filepath.Join(*dirFlag, outFile)
	if err := ioutil.WriteFile(outFile, resultData, 0666); err != nil {
		return err
	}
	return nil
}
