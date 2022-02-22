package value

import (
	"github.com/gustavocioccari/toggl-cards/pkg/models"
	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

type ValueRepository interface {
	Create(value *models.Value) error
	GetByName(name string) (*models.Value, error)
}

func NewValueRepository(db *gorm.DB) ValueRepository {
	return &repo{db: db}
}

func (r *repo) Create(value *models.Value) error {
	return r.db.Create(value).Error
}

func (r *repo) GetByName(name string) (*models.Value, error) {
	var value *models.Value

	err := r.db.Where(&models.Value{Name: name}).First(&value).Error
	if err != nil {
		return nil, err
	}
	return value, nil
}
