package async

import (
	"context"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/linker"

	"github.com/pkg/errors"
	"github.com/seibert-media/golibs/log"
	"go.uber.org/zap"
)

// twoWayRate takes a segment and rates it against all existing segments in both ways. If the resulting weight matches the threshold one or both links get added to the links
func (l *asyncLinker) twoWayRate(ctx context.Context, base, seg *linker.Segment) error {
	if err := l.rate(ctx, base, seg); err != nil {
		// Do not fail for now; return err
	}

	if err := l.rate(ctx, seg, base); err != nil {
		// Do not fail for now; return err
	}

	return nil
}

// rate takes a segment and rates it against all existing segments. If the resulting weight matches the threshold a link gets added to the linkers state
func (l *asyncLinker) rate(ctx context.Context, base, seg *linker.Segment) error {
	log.From(ctx).Info("rating", zap.Stringer("src", base), zap.Stringer("trg", seg))
	weight, err := l.rater.Rate(ctx, base.Text, seg.Text)
	if err != nil {
		log.From(ctx).Error("rating", zap.Stringer("src", base), zap.Stringer("trg", seg), zap.Error(err))
		return errors.Wrapf(err, "rating src:%s-trg:%s", base, seg)
	}

	if weight >= l.Metadata.Threshold {
		go func() {
			l.link <- &linker.Edge{
				Source: base.Node,
				Target: seg.Node,
				Weight: weight,
			}
		}()
	}

	return nil
}
