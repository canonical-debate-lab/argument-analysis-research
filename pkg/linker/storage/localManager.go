package storage

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/linker"

	"github.com/pkg/errors"
	"github.com/seibert-media/golibs/log"
	"go.uber.org/zap"
)

// LocalManager provides helpers for managing a local storage collection
type LocalManager struct {
	path string
}

// NewLocalManager with root path
func NewLocalManager(ctx context.Context, root string) *LocalManager {
	return &LocalManager{path: root}
}

// New LocalStorage for the provided id
func (m *LocalManager) New(ctx context.Context, id string) (linker.Storage, error) {
	db, err := NewLocalStorage(ctx, fmt.Sprintf("%s/db-%s", m.path, id))
	if err != nil {
		return nil, err
	}

	return db, nil
}

// List all instances of LocalStorage under the managers path
func (m *LocalManager) List(ctx context.Context) (map[string]linker.Storage, error) {
	instances := make(map[string]linker.Storage)

	filepath.Walk(m.path, func(path string, info os.FileInfo, err error) error {

		if info.IsDir() {
			return nil
		}

		log.From(ctx).Debug("loading", zap.String("path", path))

		file := filepath.Base(path)

		if strings.HasPrefix(file, "db-") {
			instance, err := NewLocalStorage(ctx, path)
			if err != nil {
				return errors.Wrap(err, "recreating storage")
			}

			instances[strings.TrimPrefix(file, "db-")] = instance
		}
		return nil
	})

	return instances, nil
}
