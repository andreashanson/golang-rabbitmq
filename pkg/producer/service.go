package producer

type Repository interface {
	Publish(b []byte) error
}

type Service struct {
	repo Repository
}

func New(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) Publish(b []byte) error {
	return s.repo.Publish(b)
}
