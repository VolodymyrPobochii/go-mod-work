package handling

import (
	cargo2 "github.com/VolodymyrPobochii/go-mod-work/domain/cargo"
	"github.com/VolodymyrPobochii/go-mod-work/domain/handling"
	"github.com/VolodymyrPobochii/go-mod-work/domain/location"
	"github.com/VolodymyrPobochii/go-mod-work/domain/voyage"
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentingService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	handling.Service
}

// NewInstrumentingService returns an instance of an instrumenting Service.
func NewInstrumentingService(counter metrics.Counter, latency metrics.Histogram, s handling.Service) handling.Service {
	return &instrumentingService{
		requestCount:   counter,
		requestLatency: latency,
		Service:        s,
	}
}

func (s *instrumentingService) RegisterHandlingEvent(completed time.Time, id cargo2.TrackingID, voyageNumber voyage.Number,
	loc location.UNLocode, eventType cargo2.HandlingEventType) error {

	defer func(begin time.Time) {
		s.requestCount.With("method", "register_incident").Add(1)
		s.requestLatency.With("method", "register_incident").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.RegisterHandlingEvent(completed, id, voyageNumber, loc, eventType)
}
