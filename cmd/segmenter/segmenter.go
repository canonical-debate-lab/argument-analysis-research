package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"gopkg.in/jdkato/prose.v2"
)

var (
	input     = flag.String("input", "", "the string to segment")
	printJSON = flag.Bool("json", false, "if true, this app will return a json representation of all segments")
)

// Output of the command if json flag is set
type Output struct {
	Input    string   `json:"input"`
	Count    int      `json:"count"`
	Segments []string `json:"segments"`
}

func main() {
	flag.Parse()

	doc, _ := prose.NewDocument(*input)

	st := doc.Sentences()

	if *printJSON {
		o := Output{Input: *input, Count: len(st)}
		for _, sent := range st {
			o.Segments = append(o.Segments, sent.Text)
		}

		b, err := json.MarshalIndent(o, "", " ")
		if err != nil {
			panic(err)
		}
		os.Stdout.Write(b)

		return
	}

	fmt.Printf("Found %d Sentences:\n", len(st))
	for _, sent := range st {
		fmt.Printf("- %s \n\n", sent.Text)
	}
}
