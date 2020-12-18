package venue

func NewService(store *store) Service {
	return &basicService{store}
}

func (s *basicService) Create(id string) error {
	o, err := NewVenue(id)

	if err != nil {
		return err
	}

	s.store.Save(o)

	return nil
}

func (s *basicService) SetName(id, name string) error {
	v := s.store.Find(id)

	if err := v.SetName(name); err == nil {
		return err
	}

	s.store.Save(v)

	return nil
}

func (s *basicService) Query(qm *QueryModel) map[string]*venue {
	return s.Query(qm)
}

type basicService struct {
	*store
}
