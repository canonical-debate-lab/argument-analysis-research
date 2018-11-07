package main

import (
	"context"
	"encoding/json"
	"flag"
	"os"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/document"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/keyword"
)

var (
	input     = flag.String("input", "", "the string to segment")
	printJSON = flag.Bool("json", false, "if true, this app will return a json representation of all segments")
)

func main() {
	flag.Parse()

	if *printJSON {
		doc, err := document.New(context.Background(), *input, keyword.Extract)

		b, err := json.MarshalIndent(doc, "", " ")
		if err != nil {
			panic(err)
		}

		os.Stdout.Write(b)

		return
	}
}
