package agenda

import (
	"testing"
	"time"
)

type spy struct {
}

func TestService(t *testing.T) {
	svc := NewService(NewStore())
	if err := svc.Create("12354"); err != nil {
		t.Error(err)
	}
	if err := svc.SetTitle("12354", "foo"); err != nil {
		t.Error(err)
	}
	if err := svc.SetDescription("12354", "bar"); err != nil {
		t.Error(err)
	}
	if err := svc.AddEntry("12354", "111"); err != nil {
		t.Error(err)
	}
	if err := svc.SetEntryTitle("12354", "111", "tull"); err != nil {
		t.Error(err)
	}
	if err := svc.SetEntryBusinessUnit("12354", "111", "dev"); err != nil {
		t.Error(err)
	}
	if err := svc.SetEntryTime("12354", "111", time.Now()); err != nil {
		t.Error(err)
	}
	if err := svc.SetEntryDescription("12354", "111", "abc"); err != nil {
		t.Error(err)
	}
	if err := svc.SetEntryVenue("12354", "111", "mcdonalds"); err != nil {
		t.Error(err)
	}
	res := svc.Query(&QueryModel{ID: "12354"})
	t.Log(res)
}
