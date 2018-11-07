package segmenter

import (
	"context"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/document"

	"github.com/seibert-media/golibs/log"
	"go.uber.org/zap"
	"gopkg.in/jdkato/prose.v2"
)

// Pipe takes a channel on which to receive string and returns a channel outputting Documents
func Pipe(ctx context.Context, in chan string) chan *document.Document {
	log.From(ctx).Info("initializing segmenting pipe")

	out := make(chan *document.Document)
	go func() {
		for content := range in {
			doc, err := prose.NewDocument(content)
			if err != nil {
				log.From(ctx).Error("segmenting", zap.String("content", content), zap.Error(err))
				continue
			}

			result := &document.Document{Content: content}

			for _, s := range doc.Sentences() {
				seg := &document.Segment{Text: s.Text}

				result.Segments = append(
					result.Segments,
					seg,
				)
			}

			out <- result
		}
	}()

	return out
}
