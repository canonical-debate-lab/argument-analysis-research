package async

import (
	"context"

	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

// DocumentCountStat .
var DocumentCountStat = stats.Int64("linker/documents/total", "Count of documents", "1")

// DocumentQueuedCountStat .
var DocumentQueuedCountStat = stats.Int64("linker/documents/queue/total", "Count of queued documents", "1")

// DocumentCountView .
var DocumentCountView = &view.View{
	Name:        "document/count",
	Measure:     DocumentCountStat,
	Description: "The number of documents stored in linker",
	TagKeys:     []tag.Key{},
	Aggregation: view.Count(),
}

// DocumentQueuedCountView .
var DocumentQueuedCountView = &view.View{
	Name:        "document/queue/count",
	Measure:     DocumentQueuedCountStat,
	Description: "The number of documents stored in linker",
	TagKeys:     []tag.Key{},
	Aggregation: view.Count(),
}

// RecordDocumentCount by locking and retrieving the amount of documents for prometheus metrics
func (l *asyncLinker) RecordDocumentCount(ctx context.Context) {
	l.dm.RLock()
	defer l.dm.RUnlock()

	stats.Record(ctx, DocumentCountStat.M(int64(len(l.documents))))
}
