package agenda

import (
	"errors"
	"time"
)

type entry struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description,omitempty"`
	EntryTime    time.Time `json:"entry_time,omitempty"`
	BusinessUnit string    `json:"business_unit,omitempty"`
	VenueID      string    `json:"venue_id,omitempty"`
	TransportID  string    `json:"transport_id,omitempty"`
	changes      []interface{}
	version      int
	state        string
}

func NewEntry(id string) (*entry, error) {
	s := &entry{version: -1}
	if err := s.apply(&created{id}); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *entry) SetTitle(title string) error {
	return s.apply(&titleSet{s.ID, title})
}

func (s *entry) SetBusinessUnit(bu string) error {
	return s.apply(&entryBusinessUnitSet{s.ID, bu})
}

func (s *entry) SetTime(entryTime time.Time) error {
	return s.apply(&entryTimeSet{s.ID, entryTime})
}

func (s *entry) SetDescription(description string) error {
	return s.apply(&descriptionSet{s.ID, description})
}

func (s *entry) SetVenue(venueID string) error {
	return s.apply(&entryVenueSet{s.ID, venueID})
}

func (s *entry) Publish() error {
	return s.apply(&entryPublished{})
}

func (s *entry) apply(e interface{}) error {
	s.when(e)
	if err := s.valid(); err != nil {
		return err
	}
	s.changes = append(s.changes, e)
	return nil
}

func (s *entry) when(e interface{}) {
	switch v := e.(type) {
	case *created:
		s.ID = v.ID
		s.state = unpublished
	case *titleSet:
		s.Title = v.Title
	case *entryBusinessUnitSet:
		s.BusinessUnit = v.BusinessUnit
	case *entryTimeSet:
		s.EntryTime = v.entryTime
	case *descriptionSet:
		s.Description = v.Description
	case *entryVenueSet:
		s.VenueID = v.venueID
	case *entryPublished:
		s.state = published
	default:
		return
	}
	s.version++
}

func (s *entry) valid() error {
	if len(s.ID) <= 0 {
		return errors.New("ID missing")
	}
	if published == s.state {
		if len(s.Title) <= 0 {
			return errors.New("Title missing")
		}
		if s.EntryTime.IsZero() {
			return errors.New("Time not set")
		}
	}
	return nil
}
