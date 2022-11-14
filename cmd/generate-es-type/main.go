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
	input   = flag.String(
		"i",
		"--",
		"input filename. set -- if you want to read from stdin.\n"+
			"Contents of the file must be what you can fetch from\n"+
			"<elasticsearch origin>/<index_name>/_mappings",
	)
	outHigh = flag.String(
		"out-high",
		"",
		"output filename to write high level types. panic if empty.",
	)
	outRaw = flag.String(
		"out-raw",
		"",
		"output filename to write raw types. panic if empty.",
	)
	mapOptPath = flag.String(
		"map-option",
		"",
		"path to a json file that can be unmarshalled to generate.MapOption.",
	)
	globalOptPath = flag.String(
		"global-option",
		"",
		"path to a json file that con be unmarshalled to generate.GlobalOption. "+
			"TypeNameGenerator field must be empty.",
	)
	prefixWithIndexName = flag.Bool(
		"prefix-with-index-name",
		false,
		"whether it add prefix to every type name with index name.",
	)
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

	mapOpt := generate.MapOption{}
	globalOpt := generate.GlobalOption{}

	indexName, mappings := getFirst(settings)

	if *mapOptPath != "" {
		decode(*mapOptPath, &mapOpt)
	}
	if *globalOptPath != "" {
		decode(*globalOptPath, &globalOpt)
	}
	if *prefixWithIndexName {
		globalOpt.TypeNameGenerator.PostProcess = []generate.TypeNamePostProcessRule{
			generate.Prefix(indexName, 1),
			generate.PascalCaseUnderscoreHyphen(),
		}
	}

	high, raw, err := generate.Generate(
		mappings,
		indexName,
		globalOpt,
		mapOpt,
	)
	if err != nil {
		panic(err)
	}

	err = generate.WriteFile(*outHigh, *outRaw, high, raw, *pkgName)
	if err != nil {
		panic(err)
	}
}

func decode(filename string, v any) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	bin, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(bin, v)
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
