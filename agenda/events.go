package agenda

import "time"

type created struct {
	ID string
}

type titleSet struct {
	ID    string
	Title string
}

type descriptionSet struct {
	ID          string
	Description string
}

type entryAdded struct {
	ID    string
	Entry *entry
}
type entryUpdated struct {
	ID    string
	Entry *entry
}

type entryPublished struct {
}

type entryBusinessUnitSet struct {
	ID           string
	BusinessUnit string
}

type entryTimeSet struct {
	ID        string
	entryTime time.Time
}

type entryVenueSet struct {
	ID      string
	venueID string
}

type agendaPublished struct{}
