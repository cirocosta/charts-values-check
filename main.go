package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cirocosta/charts-values-check/pkg"
)

var (
	filepath = flag.String("file", "", "file to read")
	fileType = flag.String("type", "values", "type of the file (readme|values)")
)

func main() {
	flag.Parse()

	if *filepath == "" {
		fmt.Fprintf(os.Stderr, "filepath must be specified")
		os.Exit(1)
	}

	var finder pkg.Finder
	switch *fileType {
	case "readme":
		finder = &pkg.ReadmeFinder{}
	case "values":
		finder = &pkg.ValuesFinder{}
	default:
		fmt.Fprintf(os.Stderr, "unknown file type %s - possible values: readme | values",
			*fileType)
		os.Exit(1)

	}

	file, err := os.Open(*filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	results, err := finder.Find(bytes)
	if err != nil {
		panic(err)
	}

	for _, result := range results {
		fmt.Println(result)
	}
}
