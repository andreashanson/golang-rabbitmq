package producer

type Repository interface {
	Publish(b []byte, queue string) error
}

type Service struct {
	repo Repository
}

func New(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) Publish(b []byte, queue string) error {
	return s.repo.Publish(b, queue)
}
