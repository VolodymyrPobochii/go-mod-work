package booking

import (
	"github.com/VolodymyrPobochii/go-mod-work/domain/booking"
	cargo2 "github.com/VolodymyrPobochii/go-mod-work/domain/cargo"
	"github.com/VolodymyrPobochii/go-mod-work/domain/location"
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentingService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	booking.Service
}

// NewInstrumentingService returns an instance of an instrumenting Service.
func NewInstrumentingService(counter metrics.Counter, latency metrics.Histogram, s booking.Service) booking.Service {
	return &instrumentingService{
		requestCount:   counter,
		requestLatency: latency,
		Service:        s,
	}
}

func (s *instrumentingService) BookNewCargo(origin, destination location.UNLocode, deadline time.Time) (cargo2.TrackingID, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "book").Add(1)
		s.requestLatency.With("method", "book").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.BookNewCargo(origin, destination, deadline)
}

func (s *instrumentingService) LoadCargo(id cargo2.TrackingID) (c booking.Cargo, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "load").Add(1)
		s.requestLatency.With("method", "load").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.LoadCargo(id)
}

func (s *instrumentingService) RequestPossibleRoutesForCargo(id cargo2.TrackingID) []cargo2.Itinerary {
	defer func(begin time.Time) {
		s.requestCount.With("method", "request_routes").Add(1)
		s.requestLatency.With("method", "request_routes").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.RequestPossibleRoutesForCargo(id)
}

func (s *instrumentingService) AssignCargoToRoute(id cargo2.TrackingID, itinerary cargo2.Itinerary) (err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "assign_to_route").Add(1)
		s.requestLatency.With("method", "assign_to_route").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.AssignCargoToRoute(id, itinerary)
}

func (s *instrumentingService) ChangeDestination(id cargo2.TrackingID, l location.UNLocode) (err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "change_destination").Add(1)
		s.requestLatency.With("method", "change_destination").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.ChangeDestination(id, l)
}

func (s *instrumentingService) Cargos() []booking.Cargo {
	defer func(begin time.Time) {
		s.requestCount.With("method", "list_cargos").Add(1)
		s.requestLatency.With("method", "list_cargos").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.Cargos()
}

func (s *instrumentingService) Locations() []booking.Location {
	defer func(begin time.Time) {
		s.requestCount.With("method", "list_locations").Add(1)
		s.requestLatency.With("method", "list_locations").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.Locations()
}
