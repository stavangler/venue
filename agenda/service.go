package agenda

import (
	"errors"
	"time"
)

type agendaService struct {
	store
}

func NewService(store store) Service {
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

func (s *agendaService) Publish(id string) error {
	a := s.store.Find(id)
	if a == nil {
		return errors.New("Unknown agenda id")
	}
	err := a.Publish()
	if err != nil {
		return err
	}
	s.store.Save(a)
	return nil
}

func (s *agendaService) SetTitle(id, title string) error {
	a := s.store.Find(id)
	if a == nil {
		return errors.New("Unknown agenda id")
	}
	err := a.SetTitle(title)
	if err != nil {
		return err
	}
	s.store.Save(a)
	return nil
}

func (s *agendaService) SetDescription(id, description string) error {
	a := s.store.Find(id)
	if a == nil {
		return errors.New("Unknown agenda id")
	}
	err := a.SetDescription(description)
	if err != nil {
		return err
	}
	s.store.Save(a)
	return nil
}

func (s *agendaService) AddEntry(id, entryID string) error {
	a := s.store.Find(id)
	if a == nil {
		return errors.New("Unknown agenda id")
	}
	e, err := NewEntry(entryID)
	if err != nil {
		return err
	}
	err = a.AddEntry(e)
	if err != nil {
		return err
	}
	s.store.Save(a)
	return nil
}
func (s *agendaService) SetEntryTitle(id, entryID, title string) error {
	a := s.store.Find(id)
	if a == nil {
		return errors.New("Unknown agenda id")
	}
	e, err := a.GetEntry(entryID)
	if err != nil {
		return err
	}
	e.SetTitle(title)
	a.UpdateEntry(&e)
	s.store.Save(a)
	return nil
}

func (s *agendaService) SetEntryBusinessUnit(id, entryID, name string) error {
	a := s.store.Find(id)
	if a == nil {
		return errors.New("Unknown agenda id")
	}
	e, err := a.GetEntry(entryID)
	if err != nil {
		return err
	}
	e.SetBusinessUnit(name)
	a.UpdateEntry(&e)
	s.store.Save(a)
	return nil
}

func (s *agendaService) SetEntryVenue(id, entryID, venueid string) error {
	a := s.store.Find(id)
	if a == nil {
		return errors.New("Unknown agenda id")
	}
	e, err := a.GetEntry(entryID)
	if err != nil {
		return err
	}
	e.SetVenue(venueid)
	a.UpdateEntry(&e)
	s.store.Save(a)
	return nil
}
func (s *agendaService) SetEntryTransport(id, entryID, transportID string) error {
	return errors.New("Not implemented")
}
func (s *agendaService) SetEntryDescription(id, entryID, description string) error {
	a := s.store.Find(id)
	if a == nil {
		return errors.New("Unknown agenda id")
	}
	e, err := a.GetEntry(entryID)
	if err != nil {
		return err
	}
	e.SetDescription(description)
	a.UpdateEntry(&e)
	s.store.Save(a)
	return nil
}

func (s *agendaService) SetEntryTime(id, entryID string, entryTime time.Time) error {
	a := s.store.Find(id)
	if a == nil {
		return errors.New("Unknown agenda id")
	}
	e, err := a.GetEntry(entryID)
	if err != nil {
		return err
	}
	e.SetTime(entryTime)
	a.UpdateEntry(&e)
	s.store.Save(a)
	return nil
}

func (s *agendaService) SetEntryPublish(id, entryID string) error {
	a := s.store.Find(id)
	if a == nil {
		return errors.New("Unknown agenda id")
	}
	e, err := a.GetEntry(entryID)
	if err != nil {
		return err
	}
	e.Publish()
	a.UpdateEntry(&e)
	s.store.Save(a)
	return nil
}

func (s *agendaService) Query(qm *QueryModel) []*agenda {
	q := s.store.Query(qm)
	qres := make([]*agenda, 0, len(q))
	for _, a := range q {
		qres = append(qres, a)
	}
	return qres
}
