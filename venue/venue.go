package venue

import (
	"fmt"
	"reflect"
)

type venue struct {
	changes []interface{}
	version int
	id      string
	name    string
}

func NewVenue(id, name string) (*venue, error) {
	o := &venue{
		version: -1,
	}
	if err := o.apply(
		&VenueCreated{
			Name: name,
			ID:   id,
		},
	); err != nil {
		return nil, err
	}
	return o, nil
}

func (o *venue) apply(e interface{}) error {
	o.when(e)
	if err := o.ensureValidState(); err != nil {
		return err
	}
	o.changes = append(o.changes, e)
	return nil
}

func (o *venue) when(e interface{}) {
	switch reflect.TypeOf(e).String() {
	case "*venue.VenueCreated":
		o.id = e.(*VenueCreated).ID
		o.name = e.(*VenueCreated).Name
	default:
		fmt.Printf("Default: %v, %s", e, reflect.TypeOf(e).String())
	}
}

func (o *venue) ensureValidState() error {
	if len(o.id) <= 0 {
		return ErrMissingID
	}

	if len(o.name) <= 0 {
		return ErrMissingName
	}
	return nil
}
