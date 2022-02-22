package card

import (
	"github.com/gustavocioccari/toggl-cards/pkg/models"
	"github.com/gustavocioccari/toggl-cards/pkg/repositories/postgres/card"
)

type service struct {
	repo card.CardRepository
}

type CardService interface {
	FindOneByCode(code string) (*models.Card, error)
}

func NewCardService(cardRepository card.CardRepository) CardService {
	return &service{
		repo: cardRepository,
	}
}

func (s *service) FindOneByCode(code string) (*models.Card, error) {
	card, err := s.repo.FindOneByCode(code)
	if err != nil {
		return nil, err
	}

	return card, nil
}
