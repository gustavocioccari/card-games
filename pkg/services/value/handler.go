package value

import (
	"github.com/gustavocioccari/toggl-cards/pkg/models"
	"github.com/gustavocioccari/toggl-cards/pkg/repositories/postgres/value"
)

type service struct {
	repo value.ValueRepository
}

type ValueService interface {
	GetByName(name string) (*models.Value, error)
}

func NewValueService(valueRepository value.ValueRepository) ValueService {
	return &service{
		repo: valueRepository,
	}
}

func (s *service) GetByName(name string) (*models.Value, error) {
	value, err := s.repo.GetByName(name)
	if err != nil {
		return nil, err
	}

	return value, nil
}
