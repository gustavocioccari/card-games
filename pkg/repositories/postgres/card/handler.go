package card

import (
	"github.com/gustavocioccari/toggl-cards/pkg/models"
	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

type CardRepository interface {
	Create(Card *models.Card) error
	FindOneByCode(code string) (*models.Card, error)
	FindByCode(codes []string) ([]models.Card, error)
}

func NewCardRepository(db *gorm.DB) CardRepository {
	return &repo{db: db}
}

func (r *repo) Create(card *models.Card) error {
	return r.db.Create(card).Error
}

func (r *repo) FindOneByCode(code string) (*models.Card, error) {
	var card *models.Card

	err := r.db.
		Where(&models.Card{Code: code}).
		Preload("Suit").
		Preload("Value").
		First(&card).Error
	if err != nil {
		return nil, err
	}
	return card, nil
}

func (r *repo) FindByCode(codes []string) ([]models.Card, error) {
	var cards []models.Card

	err := r.db.
		Where("code IN ?", codes).
		Preload("Suit").
		Preload("Value").
		Find(&cards).Error
	if err != nil {
		return nil, err
	}

	return cards, nil
}
