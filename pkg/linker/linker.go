package linker

import (
	"context"

	"github.com/pkg/errors"
	"github.com/seibert-media/golibs/log"
	"go.uber.org/zap"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/document"
)

// Linker stores all documents and compares them for building a link matrix
type Linker struct {
	Rater     Rater
	Docs      []*document.Document
	Threshold float32
}

// New Linker from documents
func New(rater Rater, documents []*document.Document, threshold float32) *Linker {
	return &Linker{
		Rater:     rater,
		Docs:      documents,
		Threshold: threshold,
	}
}

// Run the linker to compare all documents and segments
func (l *Linker) Run(ctx context.Context) ([]DocumentLinks, error) {
	links := []DocumentLinks{}
	for d, doc := range l.Docs {
		log.From(ctx).Debug("checking document", zap.Int("index", d))
		if len(doc.Segments) == 0 {
			log.From(ctx).Error("checking document", zap.Int("index", d), zap.Error(ErrNotSegmented))
			return nil, errors.Wrapf(ErrNotSegmented, "checking document %d", d)
		}

		docLinks := DocumentLinks{}
		for s, seg := range doc.Segments {
			log.From(ctx).Debug("searching links", zap.Int("document", d), zap.Int("segment", s))
			l, err := l.FindLinks(ctx, seg)
			if err != nil {
				log.From(ctx).Error("searching links", zap.Int("document", d), zap.Int("segment", s), zap.Error(err))
				return nil, errors.Wrap(err, "searching links")
			}

			docLinks = append(docLinks, l)
		}

		links = append(links, docLinks)
	}

	return links, nil
}

// FindLinks for the passed in base segment and return the links
func (l *Linker) FindLinks(ctx context.Context, base *document.Segment) (Links, error) {
	res := Links{}
	for d, doc := range l.Docs {
		for s, seg := range doc.Segments {
			log.From(ctx).Debug("rating", zap.Int("document", d), zap.Int("segment", s), zap.String("text", seg.Text), zap.String("base", base.Text))
			dist, err := l.Rater.Rate(ctx, base, seg)
			if err != nil {
				log.From(ctx).Error("rating", zap.Int("document", d), zap.Int("segment", s), zap.String("text", seg.Text), zap.String("base", base.Text), zap.Error(err))
				return nil, errors.Wrapf(err, "rating doc:%d-seg:%d", d, s)
			}

			if dist >= l.Threshold {
				res = append(res, Link{
					Document: d,
					Segment:  s,
					Dist:     dist,
				})
			}
		}
	}

	return res, nil
}
