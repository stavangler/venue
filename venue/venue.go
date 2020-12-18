package venue

import (
	"fmt"
	"reflect"
)

var (
	published  = "published"
	unpublised = "unpublished"
)

type venue struct {
	changes []interface{}
	version int
	ID      string `json:"id"`
	Name    string `json:"name,omitempty"`
	state   string
}

//NewVenue creates a new venue domain model.
func NewVenue(id string) (*venue, error) {
	o := &venue{
		version: -1,
	}
	if err := o.apply(
		&Created{
			ID: id,
		},
	); err != nil {
		return nil, err
	}
	return o, nil
}

func (s *venue) Publish() error {
	return s.apply(&Published{})
}

//SetName gives the venue a name
func (s *venue) SetName(name string) error {
	return s.apply(&NameSet{
		Name: name,
	})
}

func (s *venue) apply(e interface{}) error {
	s.when(e)
	if err := s.ensureValidState(); err != nil {
		return err
	}
	s.changes = append(s.changes, e)
	return nil
}

func (s *venue) when(e interface{}) {
	switch reflect.TypeOf(e).String() {
	case "*venue.Created":
		s.ID = e.(*Created).ID
		s.state = unpublised
	case "*venue.NameSet":
		s.Name = e.(*NameSet).Name
	case "*venue.Published":
		s.state = published
	default:
		fmt.Printf("Default: %v, %s", e, reflect.TypeOf(e).String())
	}
}

func (s *venue) ensureValidState() error {
	if len(s.ID) <= 0 {
		return ErrMissingID
	}

	if published == s.state {
		if len(s.Name) <= 0 {
			return ErrMissingName
		}
	}
	return nil
}
