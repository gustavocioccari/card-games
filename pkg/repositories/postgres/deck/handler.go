package deck

import (
	"github.com/gustavocioccari/toggl-cards/pkg/models"
	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

type DeckRepository interface {
	Create(deck *models.Deck) error
	GetByID(id string) (*models.Deck, error)
	TakeCard(id string, limit int) ([]models.Card, error)
	RemoveCards(id string, cards []models.Card) error
	UpdateRemaining(id string, remaining int) error
}

func NewDeckRepository(db *gorm.DB) DeckRepository {
	return &repo{db: db}
}

func (r *repo) Create(deck *models.Deck) error {
	return r.db.Create(deck).Error
}

func (r *repo) GetByID(id string) (*models.Deck, error) {
	var deck *models.Deck

	err := r.db.Where(&models.Deck{ID: id}).Preload("Cards.Suit").Preload("Cards.Value").First(&deck).Error
	if err != nil {
		return nil, err
	}
	return deck, nil
}

func (r *repo) TakeCard(id string, limit int) ([]models.Card, error) {
	var cards []models.Card

	err := r.db.Model(&models.Deck{ID: id}).Limit(limit).Preload("Suit").Preload("Value").Association("Cards").Find(&cards)
	if err != nil {
		return nil, err
	}

	return cards, nil
}

func (r *repo) RemoveCards(id string, cards []models.Card) error {
	return r.db.Model(&models.Deck{ID: id}).Association("Cards").Delete(&cards)
}

func (r *repo) UpdateRemaining(id string, remaining int) error {
	return r.db.Model(&models.Deck{ID: id}).Updates(&models.Deck{Remaining: remaining}).Error
}
