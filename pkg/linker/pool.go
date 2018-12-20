package linker

import (
	"context"
	"sync"

	"go.uber.org/zap"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/seibert-media/golibs/log"
)

// Pool offers management for multiple linker instances
// It takes care of creating new ones and accessing existing ones
type Pool struct {
	instanceCreator InstanceCreator
	storageManager  StorageManager

	i         sync.RWMutex
	instances map[string]Linker
}

// InstanceCreator defines the function called to initialize new Linkers
type InstanceCreator func(ctx context.Context, storage Storage) (Linker, error)

// StorageManager defines the interface required to manage Storage instances
type StorageManager interface {
	New(ctx context.Context, id string) (Storage, error)
	List(ctx context.Context) (map[string]Storage, error)
}

// Accessor stores a Linker instance and it's access information
// It is a helper type for working with pooled linkers
type Accessor struct {
	Linker
	ID string
}

// New initializes the pool and prepares for storing instances
func New(ctx context.Context, storageManager StorageManager, creator InstanceCreator) *Pool {
	p := &Pool{
		instanceCreator: creator,
		storageManager:  storageManager,

		instances: make(map[string]Linker),
	}

	return p
}

// Load existing instances from the storage manager
func (p *Pool) Load(ctx context.Context) error {
	log.From(ctx).Info("loading existing instances")

	instances, err := p.storageManager.List(ctx)
	if err != nil {
		return errors.Wrap(err, "retrieving storage instances")
	}

	for id, db := range instances {
		instance, err := p.instanceCreator(ctx, db)
		if err != nil {
			return errors.Wrap(err, "recreating instance")
		}

		p.i.Lock()
		p.instances[id] = instance
		p.i.Unlock()

		go instance.Run(ctx)
	}

	return nil
}

// Create new linker, store it in the pool and start running it's background process
func (p *Pool) Create(ctx context.Context, raterURL string, threshold float32) (*Accessor, error) {
	ctx = log.WithFields(ctx, zap.String("rater", raterURL), zap.Float32("threshold", threshold))
	log.From(ctx).Info("creating instance")

	rand, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.Wrap(err, "creating linker uuid")
	}

	if rand.String() == "" {
		return nil, errors.New("invalid uuid generated")
	}

	meta := &Metadata{
		ID:        rand.String(),
		Rater:     raterURL,
		Threshold: threshold,
	}

	log.From(ctx).Info("creating storage")
	db, err := p.storageManager.New(ctx, rand.String())
	if err != nil {
		return nil, errors.Wrap(err, "creating storage")
	}
	db.SetMetadata(ctx, meta)

	log.From(ctx).Info("creating linker")
	linker, err := p.instanceCreator(ctx, db)
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
