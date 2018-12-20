package async

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"sync"
	"time"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/document"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/linker"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/rater"

	"github.com/pkg/errors"
	"github.com/seibert-media/golibs/log"
	"go.uber.org/zap"
)

// asyncLinker implements the Linker interface by storing all documents and comparing them for building a link matrix
type asyncLinker struct {
	Metadata *linker.Metadata `json:"metadata"`

	rater rater.Rater

	docs chan *document.Document

	dm        sync.RWMutex
	documents map[string]*document.Document

	segs chan *linker.Segment

	sm       sync.RWMutex
	segments map[string]*linker.Segment

	link chan *linker.Edge

	lm    sync.RWMutex
	links map[string]*linker.Edge

	db linker.Storage
}

// New Linker for processing and persisting documents
func New(ctx context.Context, storage linker.Storage) (linker.Linker, error) {
	return &asyncLinker{

		docs:      make(chan *document.Document),
		documents: make(map[string]*document.Document),
		segments:  make(map[string]*linker.Segment),
		links:     make(map[string]*linker.Edge),
		segs:      make(chan *linker.Segment),
		link:      make(chan *linker.Edge),

		db: storage,
	}, nil
}

// Run the linker to compare all documents and segments
func (l *asyncLinker) Run(ctx context.Context) error {
	go l.handleDocs(ctx)
	go l.handleSegments(ctx)

	rate := time.Second / 2
	throttle := time.Tick(rate)

	if l.db != nil {
		if err := l.loadFromDB(ctx); err != nil {
			return errors.Wrap(err, "loading from db")
		}
	}

	for edge := range l.link {
		<-throttle

		l.lm.Lock()
		log.From(ctx).Info("storing link", zap.Stringer("edge", edge))
		l.links[edge.Hash()] = edge
		l.lm.Unlock()
		go l.db.InsertLink(ctx, edge)
	}

	return errors.New("linker closed unexpectedly")
}

// InsertDocument stores the document and all its segments into the linkers state to get them analyzed
func (l *asyncLinker) InsertDocument(ctx context.Context, doc *document.Document) error {
	log.From(ctx).Info("checking document")
	if len(doc.Segments) == 0 {
		log.From(ctx).Error("checking document", zap.Error(linker.ErrNotSegmented))
		return errors.Wrap(linker.ErrNotSegmented, "checking document")
	}

	h := sha256.New()
	h.Write([]byte(doc.Content))
	doc.Hash = hex.EncodeToString(h.Sum(nil))

	log.From(ctx).Info("inserting document", zap.String("hash", doc.Hash))
	l.docs <- doc
	go l.db.InsertDocument(ctx, doc)

	return nil
}

// ListDocuments currently stored in the linker
func (l *asyncLinker) ListDocuments(ctx context.Context) []*document.Document {
	l.dm.RLock()
	defer l.dm.RUnlock()

	docs := make([]*document.Document, 0, len(l.documents))
	for _, doc := range l.documents {
		docs = append(docs, doc)
	}

	return docs
}

// ListLinks currently stored in the linker
func (l *asyncLinker) ListLinks(ctx context.Context) []*linker.Edge {
	l.lm.RLock()
	defer l.lm.RUnlock()

	links := make([]*linker.Edge, 0, len(l.links))
	for _, edge := range l.links {
		links = append(links, edge)
	}

	return links
}
