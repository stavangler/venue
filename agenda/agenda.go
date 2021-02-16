package agenda

import (
	"errors"
	"time"
)

type agenda struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description,omitempty"`
	BusinessUnit string    `json:"businessunit,omitempty"`
	VenueID      string    `json:"venue_id,omitempty"`
	TransportID  string    `json:"transport_id,omitempty"`
	EntryTime    time.Time `json:"entry_time,omitempty"`
	changes      []interface{}
	version      int
	state        string
}

func NewAgenda(id string) (*agenda, error) {
	o := &agenda{version: -1}
	if err := o.apply(&created{id}); err != nil {
		return nil, err
	}
	return o, nil
}

func (s *agenda) apply(e interface{}) error {
	s.when(e)
	if err := s.valid(); err != nil {
		return err
	}
	s.changes = append(s.changes, e)
	return nil
}

func (s *agenda) when(e interface{}) {
	switch v := e.(type) {
	case *created:
		s.ID = v.ID
		s.state = unpublished
	case *agendaPublished:
		s.state = published
	case *titleSet:
		s.Title = v.Title
	}
}

func (s *agenda) valid() error {
	if len(s.ID) <= 0 {
		return errors.New("ID missing")
	}
	if published == s.state {
		if len(s.Title) <= 0 {
			return errors.New("Title missing")
		}
		if s.EntryTime.IsZero() {
			return errors.New("Entry time not set")
		}
	}
	return nil
}
