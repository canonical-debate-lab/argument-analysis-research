package async

import (
	"context"
	"time"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/linker"

	"github.com/seibert-media/golibs/log"
	"go.uber.org/zap"
)

func (l *asyncLinker) handleSegments(ctx context.Context) {
	for seg := range l.segs {
		log.From(ctx).Info("processing", zap.Stringer("segment", seg))
		l.sm.RLock()
		segs := make([]*linker.Segment, 0, len(l.segments))
		for _, seg := range l.segments {
			segs = append(segs, seg)
		}
		l.sm.RUnlock()

		rate := time.Second / 4
		throttle := time.Tick(rate)

		for _, trg := range segs {
			<-throttle

			go l.twoWayRate(ctx, seg, trg)
		}

		l.sm.Lock()
		l.segments[seg.Hash()] = seg
		l.sm.Unlock()
	}
}
