package main

import (
	"encoding/json"
	"flag"
	"os"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/document"
)

var (
	input     = flag.String("input", "", "the string to segment")
	printJSON = flag.Bool("json", false, "if true, this app will return a json representation of all segments")
)

func main() {
	flag.Parse()

	if *printJSON {
		doc, err := document.New(*input)

		b, err := json.MarshalIndent(doc, "", " ")
		if err != nil {
			panic(err)
		}

		os.Stdout.Write(b)

		return
	}
}
