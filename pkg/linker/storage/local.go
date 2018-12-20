package storage

import (
	"context"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/document"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/linker"

	"github.com/pkg/errors"
	"github.com/seibert-media/golibs/log"
	bolt "go.etcd.io/bbolt"
	"go.uber.org/zap"
	"gopkg.in/djherbis/stow.v3"
)

// Local wraps around boltdb and stow to persist linker data
type Local struct {
	*bolt.DB

	metadata *stow.Store

	documents, segments, links *stow.Store
}

// NewLocalStorage opens a boltdb at path and provides the linker.Storage interface
func NewLocalStorage(ctx context.Context, path string) (linker.Storage, error) {
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		log.From(ctx).Error("opening db", zap.Error(err))
		return nil, errors.Wrap(err, "opening db")
	}

	return &Local{
		DB: db,

		metadata: stow.NewJSONStore(db, []byte("metadata")),

		documents: stow.NewJSONStore(db, []byte("documents")),
		segments:  stow.NewJSONStore(db, []byte("segments")),
		links:     stow.NewJSONStore(db, []byte("links")),
	}, nil
}

// SetMetadata into stow by its hash
func (l *Local) SetMetadata(ctx context.Context, meta *linker.Metadata) {
	if err := l.metadata.Put("config", meta); err != nil {
		log.From(ctx).Error("setting metadata", zap.String("id", meta.ID), zap.Error(err))
	}
}

// Metadata retrieves the stored metadata and returns it as pointer
func (l *Local) Metadata(ctx context.Context) (*linker.Metadata, error) {
	var meta *linker.Metadata
	if err := l.metadata.Get("config", &meta); err != nil {
		log.From(ctx).Error("loading metadata", zap.Error(err))
		return nil, errors.Wrap(err, "loading metadata")
	}

	return meta, nil
}

// InsertDocument into stow by its hash
func (l *Local) InsertDocument(ctx context.Context, doc *document.Document) {
	if err := l.documents.Put(doc.Hash, doc); err != nil {
		log.From(ctx).Error("inserting document", zap.String("hash", doc.Hash), zap.Error(err))
	}
}

// InsertSegment into stow by its hash
func (l *Local) InsertSegment(ctx context.Context, seg *linker.Segment) {
	if err := l.segments.Put(seg.Hash(), seg); err != nil {
		log.From(ctx).Error("inserting segment", zap.String("hash", seg.Hash()), zap.Error(err))
	}
}

// InsertLink into stow by its hash
func (l *Local) InsertLink(ctx context.Context, link *linker.Edge) {
	if err := l.links.Put(link.Hash(), link); err != nil {
		log.From(ctx).Error("inserting link", zap.String("hash", link.Hash()), zap.Error(err))
	}
}

// Documents takes all documents currently stored in stow and returns them as hash map
func (l *Local) Documents(ctx context.Context) (map[string]*document.Document, error) {
	m := make(map[string]*document.Document)
	if err := l.documents.ForEach(func(doc *document.Document) {
		m[doc.Hash] = doc
	}); err != nil {
		return nil, errors.Wrap(err, "loading documents")
	}

	return m, nil
}

// Segments takes all segments currently stored in stow and returns them as hash map
func (l *Local) Segments(ctx context.Context) (map[string]*linker.Segment, error) {
	m := make(map[string]*linker.Segment)
	if err := l.segments.ForEach(func(seg *linker.Segment) {
		m[seg.Hash()] = seg
	}); err != nil {
		return nil, errors.Wrap(err, "loading segments")
	}

	return m, nil
}

// Links takes all links currently stored in stow and returns them as hash map
func (l *Local) Links(ctx context.Context) (map[string]*linker.Edge, error) {
	m := make(map[string]*linker.Edge)
	if err := l.links.ForEach(func(link *linker.Edge) {
		m[link.Hash()] = link
	}); err != nil {
		return nil, errors.Wrap(err, "loading links")
	}

	return m, nil
}
