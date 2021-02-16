package agenda

import (
	"errors"
	"time"
)

type agendaService struct {
	*store
}

func NewService(store *store) Service {
	return &agendaService{store}
}

func (s *agendaService) Create(id string) error {
	exist := s.store.Find(id)
	if exist != nil {
		return errors.New("Already exists")
	}
	o, err := NewAgenda(id)
	if err != nil {
		return err
	}
	s.store.Save(o)
	return nil
}
func (s *agendaService) SetBusinessUnit(id, name string) error {
	return errors.New("Not implemented")
}
func (s *agendaService) AddEntry(id, entryID string) error {
	return errors.New("Not implemented")
}
func (s *agendaService) SetEntryTitle(id, entryID, title string) error {
	return errors.New("Not implemented")
}
func (s *agendaService) SetEntryVenue(id, entryID, venueid string) error {
	return errors.New("Not implemented")
}
func (s *agendaService) SetEntryTransport(id, entryID, transportID string) error {
	return errors.New("Not implemented")
}
func (s *agendaService) SetEntryDescription(id, entryID, description string) error {
	return errors.New("Not implemented")
}
func (s *agendaService) SetEntryTime(id, entryID string, entryTime time.Time) error {
	return errors.New("Not implemented")
}
func (s *agendaService) Query(qm *QueryModel) []*agenda {
	return make([]*agenda, 0)
}
