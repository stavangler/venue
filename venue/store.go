package venue

import (
	"strings"
)

type store struct {
	venues map[string]*venue
}

func NewStore() *store {
	return &store{}
}

func (s *store) Save(venue *venue) {
	s.venues[strings.ToLower(venue.id)] = venue
}

func (s *store) Find(id string) *venue {
	return s.venues[strings.ToLower(id)]
}

func (s *store) Query(query *QueryModel) map[string]*venue {
	result := make(map[string]*venue, 0)

	if len(query.Id) > 0 {
		id := strings.ToLower(query.Id)
		if v, ok := s.venues[id]; ok {
			result[id] = v
		}
	}

	if len(query.Name) > 0 {
		for _, v := range s.venues {
			if strings.EqualFold(v.name, query.Name) {
				result[strings.ToLower(v.id)] = v
				break
			}
		}

	}
	return result
}
