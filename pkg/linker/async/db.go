package async

import (
	"context"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/rater"

	"github.com/pkg/errors"
)

// loadFromDB retrieves all persisted data from the persistence layer
func (l *asyncLinker) loadFromDB(ctx context.Context) error {
	l.lockAll()
	defer l.unlockAll()

	var err error

	l.Metadata, err = l.db.Metadata(ctx)
	if err != nil {
		return errors.Wrap(err, "restoring metadata")
	}

	l.rater = rater.NewHTTPRater(l.Metadata.Rater)

	l.documents, err = l.db.Documents(ctx)
	if err != nil {
		return errors.Wrap(err, "restoring documents")
	}

	l.segments, err = l.db.Segments(ctx)
	if err != nil {
		return errors.Wrap(err, "restoring segments")
	}

	l.links, err = l.db.Links(ctx)
	if err != nil {
		return errors.Wrap(err, "restoring links")
	}

	return nil
}

// lockAll internal data objects
func (l *asyncLinker) lockAll() {
	l.dm.Lock()
	l.sm.Lock()
	l.lm.Lock()
}

// unlockAll internal data objects
func (l *asyncLinker) unlockAll() {
	l.dm.Unlock()
	l.sm.Unlock()
	l.lm.Unlock()
}
