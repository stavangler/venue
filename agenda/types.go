package agenda

import "time"

type Service interface {
	Create(id string) error
	SetBusinessUnit(id, name string) error
	AddEntry(id, entryID string) error
	SetEntryTitle(id, entryID, title string) error
	SetEntryVenue(id, entryID, venueID string) error
	SetEntryTransport(id, entryID, transportID string) error
	SetEntryDescription(id, entryID, description string) error
	SetEntryTime(id, entryID string, entryTime time.Time) error
	Query(qm *QueryModel) []*agenda
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
