package venue

import "errors"

type Service interface {
	Create(id, name string) error
}

//Commands
type Create struct {
	Name string `json:"name,omitempty"`
	ID   string `json:"id,omitempty"`
}

//Events
type VenueCreated struct {
	Name string
	ID   string
}

var (
	ErrMissingID = errors.New("ID not specified")
	ErrMissingName = errors.New("Name not specified")
)
