package animal

type Repository interface {
	GetAll() ([]*Animal, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetAll() ([]*Animal, error) {
	var animals []*Animal
	animals, err := s.repo.GetAll()

	if err != nil {
		return animals, err
	}

	return animals, nil
}
