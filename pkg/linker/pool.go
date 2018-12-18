package linker

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/linker/db"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// Pool offers management for multiple linker instances
// It takes care of creating new ones and accessing existing ones
type Pool struct {
	creator Creator

	i         sync.RWMutex
	instances map[string]Linker
}

// Creator defines the function called to initialize new Linkers
type Creator func(ctx context.Context, id string) (Linker, error)

// Accessor stores a Linker instance and it's access information
// It is a helper type for working with pooled linkers
type Accessor struct {
	Linker
	ID string
}

// New initializes the pool and prepares for storing instances
func New(ctx context.Context, creator Creator) *Pool {
	p := &Pool{
		creator:   creator,
		instances: make(map[string]Linker),
	}

	filepath.Walk("db", func(path string, info os.FileInfo, err error) error {
		if strings.HasPrefix(path, "db-") {
			instance, err := p.creator(ctx, strings.TrimPrefix(path, "db-"))
			if err != nil {
				return errors.Wrap(err, "creating instance")
			}

			p.i.Lock()
			p.instances[strings.TrimPrefix(path, "db-")] = instance
			p.i.Unlock()
		}
		return nil
	})

	return p
}

// Create new linker, store it in the pool and start running it's background process
func (p *Pool) Create(ctx context.Context, raterURL string, threshold float32) (*Accessor, error) {
	rand, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.Wrap(err, "creating linker uuid")
	}

	if rand.String() == "" {
		return nil, errors.New("invalid uuid generated")
	}

	db, err := db.New(ctx, fmt.Sprintf("db-%s", rand.String()))
	if err != nil {
		return nil, err
	}

	meta := &Metadata{
		ID:        rand.String(),
		Rater:     raterURL,
		Threshold: threshold,
	}

	if err = db.Metadata.Put("config", meta); err != nil {
		return nil, errors.Wrap(err, "creating metadata")
	}
	if err := db.Close(); err != nil {
		return nil, errors.Wrap(err, "closing metadata db")
	}

	linker, err := p.creator(ctx, rand.String())
	if err != nil {
		return nil, errors.Wrap(err, "creating linker")
	}
	go linker.Run(ctx)

	p.i.Lock()
	p.instances[rand.String()] = linker
	p.i.Unlock()

	return &Accessor{
		Linker: linker,
		ID:     rand.String(),
	}, nil
}

// Get linker accessor by its uuid
func (p *Pool) Get(ctx context.Context, uuid string) *Accessor {
	return &Accessor{
		Linker: p.instances[uuid],
		ID:     uuid,
	}
}

// List all linker instances
func (p *Pool) List(ctx context.Context) []Linker {
	p.i.RLock()
	defer p.i.RUnlock()
	acs := make([]Linker, 0, len(p.instances))
	for _, ac := range p.instances {
		acs = append(acs, ac)
	}

	return acs
}
