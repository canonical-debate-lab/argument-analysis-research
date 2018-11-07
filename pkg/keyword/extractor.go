package keyword

import (
	"github.com/Obaied/RAKE.Go"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/document"
)

// Extract keywords from segment and store them in the returned segment
func Extract(from document.Segment) (document.Segment, error) {
	var keywords []*document.Keyword
	for _, k := range rake.RunRake(from.Text) {
		keywords = append(keywords,
			&document.Keyword{
				Key:   k.Key,
				Value: k.Value,
			})

		from.Keywords = keywords
	}
	return from, nil
}
