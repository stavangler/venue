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

func (ls *loggingService) Create(id string) (err error) {
	defer func(begin time.Time) {
		ls.Log(
			"method", "create_venue",
			"took", time.Since(begin),
			"id", id,
			"err", err,
		)
	}(time.Now())
	return ls.s.Create(id)
}

func (ls *loggingService) SetName(id, name string) (err error) {
	defer func(begin time.Time) {
		ls.Log(
			"method", "set_name",
			"took", time.Since(begin),
			"id", id,
			"name", name,
			"err", err,
		)
	}(time.Now())
	return ls.s.SetName(id, name)
}

func (ls *loggingService) Query(qm *QueryModel) map[string]*venue{
	defer func(begin time.Time) {
		ls.Log(
			"method", "query",
			"took", time.Since(begin),
			"id", qm.Id,
			"name", qm.Name,
		)
	}(time.Now())
	return ls.s.Query(qm)
}
