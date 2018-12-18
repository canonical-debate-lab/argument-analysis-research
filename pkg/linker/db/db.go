package db

import (
	"context"

	"github.com/pkg/errors"
	"github.com/seibert-media/golibs/log"
	bolt "go.etcd.io/bbolt"
	"go.uber.org/zap"
	"gopkg.in/djherbis/stow.v3"
)

// DB wraps around boltdb and stow to persist linker data
type DB struct {
	*bolt.DB

	Metadata *stow.Store

	Documents, Segments, Links *stow.Store
}

// New opens the DB at path and provides the persistence layer
func New(ctx context.Context, path string) (*DB, error) {
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		log.From(ctx).Error("opening db", zap.Error(err))
		return nil, errors.Wrap(err, "opening db")
	}

	return &DB{
		DB: db,

		Metadata: stow.NewJSONStore(db, []byte("metadata")),

		Documents: stow.NewJSONStore(db, []byte("documents")),
		Segments:  stow.NewJSONStore(db, []byte("statements")),
		Links:     stow.NewJSONStore(db, []byte("links")),
	}, nil
}
