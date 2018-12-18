package async

import (
	"context"
	"time"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/document"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/linker"
)

func (l *asyncLinker) handleDocs(ctx context.Context) {
	rate := time.Second / 100
	throttle := time.Tick(rate)

	for doc := range l.docs {
		<-throttle

		l.dm.Lock()
		if _, exists := l.documents[doc.Hash]; exists {
			l.dm.Unlock()
			continue
		}
		l.documents[doc.Hash] = doc
		l.dm.Unlock()

		for s, seg := range doc.Segments {
			go func(doc *document.Document, seg *document.Segment, s int) {
				l.segs <- &linker.Segment{
					Node: &linker.Node{Doc: doc.Hash, Seg: s},
					Text: seg.Text,
				}
			}(doc, seg, s)
		}
	}
}
