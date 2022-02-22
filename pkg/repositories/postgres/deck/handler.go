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
}

func NewDeckRepository(db *gorm.DB) DeckRepository {
	return &repo{db: db}
}

func (r *repo) Create(deck *models.Deck) error {
	return r.db.Create(deck).Error
}

func (r *repo) GetByID(id string) (*models.Deck, error) {
	var deck *models.Deck

	err := r.db.Where(&models.Deck{ID: id}).First(&deck).Error
	if err != nil {
		return nil, err
	}
	return deck, nil
}
