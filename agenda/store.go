package agenda

import (
	"strings"
)

type memstore struct {
	agendas map[string]*agenda
}

func NewStore() *memstore {
	return &memstore{
		agendas: make(map[string]*agenda, 0),
	}
}

func (s *memstore) Save(agenda *agenda) {
	s.agendas[strings.ToLower(agenda.ID)] = agenda
}

func (s *memstore) Find(id string) *agenda {
	return s.agendas[strings.ToLower(id)]
}

func (s *memstore) Query(query *QueryModel) map[string]*agenda {
	result := make(map[string]*agenda, 0)

	if len(query.ID) > 0 {
		id := strings.ToLower(query.ID)
		if v, ok := s.agendas[id]; ok {
			result[id] = v
		}
	}

	if len(query.Title) > 0 {
		for _, v := range s.agendas {
			if strings.EqualFold(v.Title, query.Title) {
				result[strings.ToLower(v.ID)] = v
				break
			}
		}

	}
	return result
}
