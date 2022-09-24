package handling

import (
	cargo2 "github.com/VolodymyrPobochii/go-mod-work/cargo"
	"github.com/VolodymyrPobochii/go-mod-work/location"
	"github.com/VolodymyrPobochii/go-mod-work/voyage"
	"time"

	"github.com/go-kit/kit/log"
)

type loggingService struct {
	logger log.Logger
	Service
}

// NewLoggingService returns a new instance of a logging Service.
func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s *loggingService) RegisterHandlingEvent(completed time.Time, id cargo2.TrackingID, voyageNumber voyage.Number,
	unLocode location.UNLocode, eventType cargo2.HandlingEventType) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "register_incident",
			"tracking_id", id,
			"location", unLocode,
			"voyage", voyageNumber,
			"event_type", eventType,
			"completion_time", completed,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.RegisterHandlingEvent(completed, id, voyageNumber, unLocode, eventType)
}
