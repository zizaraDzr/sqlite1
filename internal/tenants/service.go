package tenants

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAll() ([]Tenant, error) {
	return s.repo.GetAll()
}

func (s *Service) GetByGUID(guid string) (*Tenant, error) {
	return s.repo.GetByGUID(guid)
}

func (s *Service) Create(t *Tenant) error {
	return s.repo.Create(t)
}

func (s *Service) Update(guid string, t *Tenant) error {
	return s.repo.Update(guid, t)
}

func (s *Service) Delete(guid string) error {
	return s.repo.Delete(guid)
}
