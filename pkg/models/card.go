package models

import (
	"github.com/go-playground/validator"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Card struct {
	ID      string `json:"id" gorm:"primaryKey"`
	ValueID string `json:"value_id" gorm:"not null" validate:"required"`
	Value   Value  `json:"value"`
	SuitID  string `json:"suit_id" gorm:"not null" validate:"required"`
	Suit    Suit   `json:"suit"`
	Code    string `json:"code" gorm:"not null"`
}

func (c *Card) Validate() error {
	return validator.New().Struct(c)
}

func (Card) TableName() string {
	return "card"
}

func (c *Card) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.NewV4().String()
	return
}
