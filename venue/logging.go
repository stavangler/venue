package venue

import (
	"time"

	"github.com/go-kit/kit/log"
)

type loggingService struct {
	log.Logger
	s Service
}

func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (ls *loggingService) Create(id, name string) (err error) {
	defer func(begin time.Time) {
		ls.Log(
			"method", "create_venue",
			"took", time.Since(begin),
			"name", name,
			"id", id,
			"err", err,
		)
	}(time.Now())
	return ls.s.Create(id, name)
}
