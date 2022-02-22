package models

import (
	"github.com/go-playground/validator"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Deck struct {
	ID        string `json:"id" gorm:"primaryKey"`
	Shuffled  bool   `json:"shuffled" gorm:"not null" validate:"required"`
	Seed      int64  `json:"-" gorm:"not null" validate:"required"`
	Remaining int    `json:"remaining" gorm:"not null" validate:"required"`
	Cards     []Card `json:"cards" gorm:"many2many:deck_card" validate:"required"`
}

func (d *Deck) Validate() error {
	return validator.New().Struct(d)
}

func (Deck) TableName() string {
	return "deck"
}

func (d *Deck) BeforeCreate(tx *gorm.DB) (err error) {
	d.ID = uuid.NewV4().String()
	return
}
