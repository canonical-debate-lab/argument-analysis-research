package linker

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"sync"
	"time"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/document"

	"github.com/pkg/errors"
	"github.com/seibert-media/golibs/log"
	"go.uber.org/zap"
)

// Linker stores all documents and compares them for building a link matrix
type Linker struct {
	Rater     Rater
	Threshold float32

	docs chan *document.Document

	dm        sync.RWMutex
	Documents map[string]*document.Document

	segs chan *Segment

	sm       sync.RWMutex
	Segments []*Segment

	link chan *Edge

	lm    sync.RWMutex
	Links []*Edge
}

// New Linker from documents
func New(rater Rater, threshold float32) *Linker {
	return &Linker{
		Rater:     rater,
		Threshold: threshold,

		docs:      make(chan *document.Document),
		Documents: make(map[string]*document.Document),
		segs:      make(chan *Segment),
		link:      make(chan *Edge),
	}
}

// Run the linker to compare all documents and segments
func (l *Linker) Run(ctx context.Context) error {
	go l.handleDocs(ctx)
	go l.handleSegments(ctx)

	rate := time.Second / 2
	throttle := time.Tick(rate)

	for edge := range l.link {
		<-throttle

		l.lm.Lock()
		log.From(ctx).Info("storing link", zap.Stringer("edge", edge))
		l.Links = append(l.Links, edge)
		l.lm.Unlock()
	}

	return errors.New("linker closed unexpectedly")
}

func (l *Linker) handleDocs(ctx context.Context) {

	rate := time.Second / 2
	throttle := time.Tick(rate)

	for doc := range l.docs {
		<-throttle

		l.dm.Lock()
		if _, exists := l.Documents[doc.Hash]; exists {
			l.dm.Unlock()
			continue
		}
		l.Documents[doc.Hash] = doc
		l.dm.Unlock()

		for s, seg := range doc.Segments {
			l.segs <- &Segment{
				Node: &Node{Doc: doc.Hash, Seg: s},
				Text: seg.Text,
			}
		}
	}
}

func (l *Linker) handleSegments(ctx context.Context) {
	for seg := range l.segs {
		log.From(ctx).Info("processing", zap.Stringer("segment", seg))
		l.sm.RLock()
		segs := make([]*Segment, 0, len(l.Segments)+1)
		segs = append(segs, l.Segments...)
		l.sm.RUnlock()

		rate := time.Second / 4
		throttle := time.Tick(rate)

		for _, trg := range segs {
			<-throttle

			go l.TwoWayRate(ctx, seg, trg)
		}

		segs = append(segs, seg)
		l.sm.Lock()
		l.Segments = segs
		l.sm.Unlock()
	}
}

// TwoWayRate takes a segment and rates it against all existing segments in both ways. If the resulting weight matches the threshold one or both links get added to the links
func (l *Linker) TwoWayRate(ctx context.Context, base, seg *Segment) error {
	if err := l.Rate(ctx, base, seg); err != nil {
		// Do not fail for now; return err
	}

	if err := l.Rate(ctx, seg, base); err != nil {
		// Do not fail for now; return err
	}

	return nil
}

// Rate takes a segment and rates it against all existing segments. If the resulting weight matches the threshold a link gets added to the linkers state
func (l *Linker) Rate(ctx context.Context, base, seg *Segment) error {
	log.From(ctx).Info("rating", zap.Stringer("src", base), zap.Stringer("trg", seg))
	weight, err := l.Rater.Rate(ctx, base.Text, seg.Text)
	if err != nil {
		log.From(ctx).Error("rating", zap.Stringer("src", base), zap.Stringer("trg", seg), zap.Error(err))
		return errors.Wrapf(err, "rating src:%s-trg:%s", base, seg)
	}

	if weight >= l.Threshold {
		go func() {
			l.link <- &Edge{
				Source: base.Node,
				Target: seg.Node,
				Weight: weight,
			}
		}()
	}

	return nil
}

// InsertDocument stores the document and all its segments into the linkers state to get them analyzed
func (l *Linker) InsertDocument(ctx context.Context, doc *document.Document) error {
	log.From(ctx).Info("checking document")
	if len(doc.Segments) == 0 {
		log.From(ctx).Error("checking document", zap.Error(ErrNotSegmented))
		return errors.Wrap(ErrNotSegmented, "checking document")
	}

	h := sha256.New()
	h.Write([]byte(doc.Content))
	doc.Hash = hex.EncodeToString(h.Sum(nil))

	log.From(ctx).Info("inserting document", zap.String("hash", doc.Hash))
	l.docs <- doc

	return nil
}

// ListDocuments currently stored in the linker
func (l *Linker) ListDocuments(ctx context.Context) []*document.Document {
	l.dm.RLock()
	defer l.dm.RUnlock()

	docs := make([]*document.Document, 0, len(l.Documents))
	for _, doc := range l.Documents {
		docs = append(docs, doc)
	}

	return docs
}

// ListLinks currently stored in the linker
func (l *Linker) ListLinks(ctx context.Context) []*Edge {
	l.lm.RLock()
	defer l.lm.RUnlock()

	links := make([]*Edge, len(l.Links))
	copy(links, l.Links)

	return links
}
