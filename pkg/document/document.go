package document

import (
	"github.com/Obaied/RAKE.Go"
	"gopkg.in/jdkato/prose.v2"
)

// Document starts with Content which is then analyzed
type Document struct {
	Content  string     `json:"content"`
	Segments []*Segment `json:"segments"`
}

// Segment of a Document
type Segment struct {
	Text     string     `json:"text"`
	Keywords []*Keyword `json:"keywords"`
}

// Keyword as found inside a Segment
type Keyword struct {
	Key   string  `json:"key"`
	Value float64 `json:"value"`
}

// New takes the input string and runs basic analysis
func New(from string) (*Document, error) {
	doc, err := prose.NewDocument(from)
	if err != nil {
		return nil, err
	}

	result := &Document{Content: from}

	for _, s := range doc.Sentences() {
		seg := &Segment{Text: s.Text}

		var keywords []*Keyword
		for _, k := range rake.RunRake(s.Text) {
			keywords = append(
				keywords,
				&Keyword{
					Key:   k.Key,
					Value: k.Value,
				})
		}

		seg.Keywords = keywords
		result.Segments = append(
			result.Segments,
			seg,
		)
	}

	return result, nil
}
