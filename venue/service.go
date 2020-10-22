package venue

func NewService() Service {
	return &basicService{}
}

func (s *basicService) Create(id, name string) error {
	_, err := NewVenue(id, name)

	if err != nil {
		return err
	}

	//err = s.store.Save(o)

	return nil
}

type basicService struct {
}
