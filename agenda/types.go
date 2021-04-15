package agenda

import "time"

type Service interface {
	Create(id string) error
	SetTitle(id, title string) error
	SetDescription(id, description string) error
	AddEntry(id, entryID string) error
	SetEntryTitle(id, entryID, title string) error
	SetEntryBusinessUnit(id, entryID, name string) error
	SetEntryVenue(id, entryID, venueID string) error
	SetEntryTransport(id, entryID, transportID string) error
	SetEntryDescription(id, entryID, description string) error
	SetEntryTime(id, entryID string, entryTime time.Time) error
	SetEntryPublish(id, entryID string) error
	Query(qm *QueryModel) []*agenda
	Publish(id string) error
}

type QueryModel struct {
	ID       string
	Title    string
	FromTime time.Time
	ToTime   time.Time
}

const (
	unpublished string = "unpublished"
	published   string = "published"
)

type store interface {
	Save(agenda *agenda)
	Find(id string) *agenda
	Query(query *QueryModel) map[string]*agenda
}
