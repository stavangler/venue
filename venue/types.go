package venue

import "errors"

type Service interface {
	Create(id string) error
	SetName(id, name string) error
	Query(qm *QueryModel) map[string]*venue
}

type QueryModel struct {
	Name string
	Id   string
}

//Commands
type Create struct {
	ID string `json:"id,omitempty"`
}

type SetName struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

//Events
type Created struct {
	ID string
}

type Published struct{}

type NameSet struct {
	Name string
}

var (
	ErrMissingID   = errors.New("ID not specified")
	ErrMissingName = errors.New("Name not specified")
)
