package agenda

import (
	"testing"
)

type spy struct {
}

func TestService(t *testing.T) {
	svc := NewService(NewStore())
	svc.Create("12354")
}
