package agenda

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
			"method", "create_agenda",
			"took", time.Since(begin),
			"id", id,
			"err", err,
		)
	}(time.Now())
	return ls.s.Create(id)
}

func (ls *loggingService) SetTitle(id, title string) (err error) {
	defer func(begin time.Time) {
		ls.Log(
			"method", "set_agenda_title",
			"took", time.Since(begin),
			"id", id,
			"err", err,
		)
	}(time.Now())
	return ls.s.SetTitle(id, title)
}

func (ls *loggingService) SetDescription(id, description string) (err error) {
	defer func(begin time.Time) {
		ls.Log(
			"method", "set_agenda_description",
			"took", time.Since(begin),
			"id", id,
			"err", err,
		)
	}(time.Now())
	return ls.s.SetDescription(id, description)
}

func (ls *loggingService) SetEntryBusinessUnit(id, entryID, name string) (err error) {
	defer func(begin time.Time) {
		ls.Log(
			"method", "set_business_unit",
			"took", time.Since(begin),
			"id", id,
			"name", name,
			"err", err,
		)
	}(time.Now())
	return ls.s.SetEntryBusinessUnit(id, entryID, name)
}

func (ls *loggingService) AddEntry(id, entryID string) (err error) {
	return ls.s.AddEntry(id, entryID)
}
func (ls *loggingService) SetEntryTitle(id, entryID, title string) (err error) {
	return ls.s.SetEntryTitle(id, entryID, title)
}
func (ls *loggingService) SetEntryVenue(id, entryID, venueID string) (err error) {
	return ls.s.SetEntryVenue(id, entryID, venueID)
}
func (ls *loggingService) SetEntryTransport(id, entryID, transportID string) (err error) {
	return ls.s.SetEntryTransport(id, entryID, transportID)
}
func (ls *loggingService) SetEntryDescription(id, entryID, description string) (err error) {
	return ls.s.SetEntryDescription(id, entryID, description)
}
func (ls *loggingService) SetEntryTime(id, entryID string, entryTime time.Time) (err error) {
	return ls.s.SetEntryTime(id, entryID, entryTime)
}

func (ls *loggingService) Query(qm *QueryModel) []*agenda {
	defer func(begin time.Time) {
		ls.Log(
			"method", "query",
			"took", time.Since(begin),
			"id", qm.ID,
			"name", qm.Title,
		)
	}(time.Now())
	return ls.s.Query(qm)
}
