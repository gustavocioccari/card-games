package suit

import (
	"github.com/gustavocioccari/toggl-cards/pkg/models"
	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

type SuitRepository interface {
	Create(suit *models.Suit) error
	GetByName(name string) (*models.Suit, error)
}

func NewSuitRepository(db *gorm.DB) SuitRepository {
	return &repo{db: db}
}

func (r *repo) Create(suit *models.Suit) error {
	return r.db.Create(suit).Error
}

func (r *repo) GetByName(name string) (*models.Suit, error) {
	var suit *models.Suit

	err := r.db.Where(&models.Suit{Name: name}).First(&suit).Error
	if err != nil {
		return nil, err
	}
	return suit, nil
}
