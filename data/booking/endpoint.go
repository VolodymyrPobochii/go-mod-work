package booking

import (
	"context"
	"github.com/VolodymyrPobochii/go-mod-work/domain/booking"
	cargo2 "github.com/VolodymyrPobochii/go-mod-work/domain/cargo"
	"github.com/VolodymyrPobochii/go-mod-work/domain/location"
	"time"

	"github.com/go-kit/kit/endpoint"
)

type bookCargoRequest struct {
	Origin          location.UNLocode
	Destination     location.UNLocode
	ArrivalDeadline time.Time
}

type bookCargoResponse struct {
	ID  cargo2.TrackingID `json:"tracking_id,omitempty"`
	Err error             `json:"error,omitempty"`
}

func (r bookCargoResponse) error() error { return r.Err }

func makeBookCargoEndpoint(s booking.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(bookCargoRequest)
		id, err := s.BookNewCargo(req.Origin, req.Destination, req.ArrivalDeadline)
		return bookCargoResponse{ID: id, Err: err}, nil
	}
}

type loadCargoRequest struct {
	ID cargo2.TrackingID
}

type loadCargoResponse struct {
	Cargo *booking.Cargo `json:"cargo,omitempty"`
	Err   error          `json:"error,omitempty"`
}

func (r loadCargoResponse) error() error { return r.Err }

func makeLoadCargoEndpoint(s booking.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(loadCargoRequest)
		c, err := s.LoadCargo(req.ID)
		return loadCargoResponse{Cargo: &c, Err: err}, nil
	}
}

type requestRoutesRequest struct {
	ID cargo2.TrackingID
}

type requestRoutesResponse struct {
	Routes []cargo2.Itinerary `json:"routes,omitempty"`
	Err    error              `json:"error,omitempty"`
}

func (r requestRoutesResponse) error() error { return r.Err }

func makeRequestRoutesEndpoint(s booking.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requestRoutesRequest)
		itin := s.RequestPossibleRoutesForCargo(req.ID)
		return requestRoutesResponse{Routes: itin, Err: nil}, nil
	}
}

type assignToRouteRequest struct {
	ID        cargo2.TrackingID
	Itinerary cargo2.Itinerary
}

type assignToRouteResponse struct {
	Err error `json:"error,omitempty"`
}

func (r assignToRouteResponse) error() error { return r.Err }

func makeAssignToRouteEndpoint(s booking.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(assignToRouteRequest)
		err := s.AssignCargoToRoute(req.ID, req.Itinerary)
		return assignToRouteResponse{Err: err}, nil
	}
}

type changeDestinationRequest struct {
	ID          cargo2.TrackingID
	Destination location.UNLocode
}

type changeDestinationResponse struct {
	Err error `json:"error,omitempty"`
}

func (r changeDestinationResponse) error() error { return r.Err }

func makeChangeDestinationEndpoint(s booking.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(changeDestinationRequest)
		err := s.ChangeDestination(req.ID, req.Destination)
		return changeDestinationResponse{Err: err}, nil
	}
}

type listCargosRequest struct{}

type listCargosResponse struct {
	Cargos []booking.Cargo `json:"cargos,omitempty"`
	Err    error           `json:"error,omitempty"`
}

func (r listCargosResponse) error() error { return r.Err }

func makeListCargosEndpoint(s booking.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(listCargosRequest)
		return listCargosResponse{Cargos: s.Cargos(), Err: nil}, nil
	}
}

type listLocationsRequest struct {
}

type listLocationsResponse struct {
	Locations []booking.Location `json:"locations,omitempty"`
	Err       error              `json:"error,omitempty"`
}

func makeListLocationsEndpoint(s booking.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(listLocationsRequest)
		return listLocationsResponse{Locations: s.Locations(), Err: nil}, nil
	}
}
