package tracking

import (
	"context"
	"github.com/VolodymyrPobochii/go-mod-work/domain/tracking"

	"github.com/go-kit/kit/endpoint"
)

type trackCargoRequest struct {
	ID string
}

type trackCargoResponse struct {
	Cargo *tracking.Cargo `json:"cargo,omitempty"`
	Err   error           `json:"error,omitempty"`
}

func (r trackCargoResponse) error() error { return r.Err }

func makeTrackCargoEndpoint(ts tracking.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(trackCargoRequest)
		c, err := ts.Track(req.ID)
		return trackCargoResponse{Cargo: &c, Err: err}, nil
	}
}
