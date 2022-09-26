package tracking

import (
	"github.com/VolodymyrPobochii/go-mod-work/domain/tracking"
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentingService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	tracking.Service
}

// NewInstrumentingService returns an instance of an instrumenting Service.
func NewInstrumentingService(counter metrics.Counter, latency metrics.Histogram, s tracking.Service) tracking.Service {
	return &instrumentingService{
		requestCount:   counter,
		requestLatency: latency,
		Service:        s,
	}
}

func (s *instrumentingService) Track(id string) (tracking.Cargo, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "track").Add(1)
		s.requestLatency.With("method", "track").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.Track(id)
}
