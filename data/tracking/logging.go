package tracking

import (
	"github.com/VolodymyrPobochii/go-mod-work/domain/tracking"
	"time"

	"github.com/go-kit/kit/log"
)

type loggingService struct {
	logger log.Logger
	tracking.Service
}

// NewLoggingService returns a new instance of a logging Service.
func NewLoggingService(logger log.Logger, s tracking.Service) tracking.Service {
	return &loggingService{logger, s}
}

func (s *loggingService) Track(id string) (c tracking.Cargo, err error) {
	defer func(begin time.Time) {
		s.logger.Log("method", "track", "tracking_id", id, "took", time.Since(begin), "err", err)
	}(time.Now())
	return s.Service.Track(id)
}
