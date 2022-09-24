package handling

import (
	"context"
	"encoding/json"
	cargo2 "github.com/VolodymyrPobochii/go-mod-work/cargo"
	"github.com/VolodymyrPobochii/go-mod-work/location"
	"github.com/VolodymyrPobochii/go-mod-work/voyage"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
)

// MakeHandler returns a handler for the handling service.
func MakeHandler(hs Service, logger kitlog.Logger) http.Handler {
	r := mux.NewRouter()

	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	registerIncidentHandler := kithttp.NewServer(
		makeRegisterIncidentEndpoint(hs),
		decodeRegisterIncidentRequest,
		encodeResponse,
		opts...,
	)

	r.Handle("/handling/v1/incidents", registerIncidentHandler).Methods("POST")

	return r
}

func decodeRegisterIncidentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		CompletionTime time.Time `json:"completion_time"`
		TrackingID     string    `json:"tracking_id"`
		VoyageNumber   string    `json:"voyage"`
		Location       string    `json:"location"`
		EventType      string    `json:"event_type"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return registerIncidentRequest{
		CompletionTime: body.CompletionTime,
		ID:             cargo2.TrackingID(body.TrackingID),
		Voyage:         voyage.Number(body.VoyageNumber),
		Location:       location.UNLocode(body.Location),
		EventType:      stringToEventType(body.EventType),
	}, nil
}

func stringToEventType(s string) cargo2.HandlingEventType {
	types := map[string]cargo2.HandlingEventType{
		cargo2.Receive.String(): cargo2.Receive,
		cargo2.Load.String():    cargo2.Load,
		cargo2.Unload.String():  cargo2.Unload,
		cargo2.Customs.String(): cargo2.Customs,
		cargo2.Claim.String():   cargo2.Claim,
	}
	return types[s]
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case cargo2.ErrUnknown:
		w.WriteHeader(http.StatusNotFound)
	case ErrInvalidArgument:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
