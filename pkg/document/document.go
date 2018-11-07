package document

import (
	"context"

	"github.com/pkg/errors"
	"github.com/seibert-media/golibs/log"
	"go.uber.org/zap"
	"gopkg.in/jdkato/prose.v2"
)

// Document starts with Content which is then analyzed
type Document struct {
	Content  string     `json:"content"`
	Segments []*Segment `json:"segments,omitempty"`
}

// Segment of a Document
type Segment struct {
	Text     string     `json:"text"`
	Keywords []*Keyword `json:"keywords,omitempty"`
}

// Keyword as found inside a Segment
type Keyword struct {
	Key   string  `json:"key"`
	Value float64 `json:"value"`
}

// Step taking a segment and returning updated segment or error
type Step func(Segment) (Segment, error)

// New takes the input string for segmenting and
// runs all segments through the steps before returning the final document
func New(ctx context.Context, from string, steps ...Step) (*Document, error) {
	doc, err := prose.NewDocument(from)
	if err != nil {
		return nil, err
	}

	result := &Document{Content: from}

	for _, s := range doc.Sentences() {
		seg := Segment{Text: s.Text}

		for i, step := range steps {
			seg, err = step(seg)
			if err != nil {
				log.From(ctx).Error("applying step", zap.Int("step", i), zap.String("segment", s.Text))
				return nil, errors.Wrapf(err, "applying step %d", i)
			}
		}

		result.Segments = append(
			result.Segments,
			&seg,
		)
	}

	return result, nil
}
