package main

import (
	"encoding/json"
	"flag"
	"io"
	"os"

	"github.com/ngicks/elastic-type/generate"
	"github.com/ngicks/elastic-type/mapping"
)

var (
	pkgName = flag.String("pkg", "example", "package to generate.")
	input   = flag.String("i", "--", "input filename. set -- if you want to read from stdin.")
	outHigh = flag.String("out-high", "", "output filename to write high level types. panic if empty.")
	outRaw  = flag.String("out-raw", "", "output filename to write raw types. panic if empty.")
)

func main() {
	flag.Parse()

	if *pkgName == "" || *input == "" || *outHigh == "" || *outRaw == "" {
		panic("pkgName, input, outHigh or outRaw is empty")
	}

	var inFile *os.File
	if *input == "--" {
		inFile = os.Stdin
	} else {
		var err error
		inFile, err = os.Open(*input)
		if err != nil {
			panic(err)
		}
		defer inFile.Close()
	}

	var err error
	bin, err := io.ReadAll(inFile)
	if err != nil {
		panic(err)
	}

	var settings mapping.MappingSettings

	err = json.Unmarshal(bin, &settings)
	if err != nil {
		panic(err)
	}

	indexName, mappings := getFirst(settings)

	high, raw, err := generate.Generate(
		mappings.Properties,
		indexName,
		generate.GlobalOption{},
		generate.MapOption{},
	)
	if err != nil {
		panic(err)
	}

	err = generate.WriteFile(*outHigh, *outRaw, high, raw, *pkgName)
	if err != nil {
		panic(err)
	}
}

func getFirst(s mapping.MappingSettings) (indexName string, mappings mapping.Mappings) {
	for k, v := range s {
		return k, *v.Mappings
	}
	panic("nah")
}
