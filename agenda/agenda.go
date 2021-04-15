package agenda

import (
	"errors"
)

type agenda struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description,omitempty"`
	Entries     []entry `json:"entries,omitempty"`
	changes     []interface{}
	version     int
	state       string
}

func NewAgenda(id string) (*agenda, error) {
	o := &agenda{version: -1}
	if err := o.apply(&created{id}); err != nil {
		return nil, err
	}
	return o, nil
}

func (s *agenda) SetTitle(title string) error {
	return s.apply(&titleSet{s.ID, title})
}

func (s *agenda) SetDescription(description string) error {
	return s.apply(&descriptionSet{s.ID, description})
}

func (s *agenda) AddEntry(e *entry) error {
	return s.apply(&entryAdded{s.ID, e})
}

func (s *agenda) Publish() error {
	return s.apply(&agendaPublished{})
}

func (s *agenda) GetEntry(entryID string) (entry, error) {
	for _, e := range s.Entries {
		if e.ID == entryID {
			return e, nil
		}
	}
	return entry{}, errors.New("Not found")
}

func (s *agenda) UpdateEntry(e *entry) error {
	return s.apply(&entryUpdated{e.ID, e})
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
	case *descriptionSet:
		s.Description = v.Description
	case *entryAdded:
		found := false
		for _, entry := range s.Entries {
			if entry.ID == v.ID {
				found = true
			}
		}
		if found == false {
			s.Entries = append(s.Entries, *v.Entry)
		}
	case *entryUpdated:
		for i, entry := range s.Entries {
			if entry.ID == v.ID {
				s.Entries[i] = *v.Entry
				continue
			}
		}
	default:
		return
	}
	s.version++
}

func (s *agenda) valid() error {
	if len(s.ID) <= 0 {
		return errors.New("ID missing")
	}
	if published == s.state {
		if len(s.Title) <= 0 {
			return errors.New("Title missing")
		}
		for _, e := range s.Entries {
			if e.state != published {
				return errors.New("not all events in agenda are published")
			}
		}
	}
	return nil
}
