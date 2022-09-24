package handling

import (
	"context"
	cargo2 "github.com/VolodymyrPobochii/go-mod-work/cargo"
	"github.com/VolodymyrPobochii/go-mod-work/location"
	"github.com/VolodymyrPobochii/go-mod-work/voyage"
	"time"

	"github.com/go-kit/kit/endpoint"
)

type registerIncidentRequest struct {
	ID             cargo2.TrackingID
	Location       location.UNLocode
	Voyage         voyage.Number
	EventType      cargo2.HandlingEventType
	CompletionTime time.Time
}

type registerIncidentResponse struct {
	Err error `json:"error,omitempty"`
}

func (r registerIncidentResponse) error() error { return r.Err }

func makeRegisterIncidentEndpoint(hs Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(registerIncidentRequest)
		err := hs.RegisterHandlingEvent(req.CompletionTime, req.ID, req.Voyage, req.Location, req.EventType)
		return registerIncidentResponse{Err: err}, nil
	}
}
