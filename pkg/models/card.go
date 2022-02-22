package models

import (
	"github.com/go-playground/validator"
)

type Card struct {
	ID      string `json:"id" gorm:"primaryKey"`
	Code    string `json:"code" gorm:"unique" validate:"required"`
	ValueID string `json:"value_id" gorm:"not null" validate:"required"`
	Value   Value  `json:"value,omitempty"`
	SuitID  string `json:"suit_id" gorm:"not null" validate:"required"`
	Suit    Suit   `json:"suit,omitempty"`
}

func (c *Card) Validate() error {
	return validator.New().Struct(c)
}

func (c *Card) TableName() string {
	return "card"
}
