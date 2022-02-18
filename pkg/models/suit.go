package models

import (
	"github.com/go-playground/validator"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Suit struct {
	ID   string `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"unique,not null" validate:"required"`
}

func (s *Suit) Validate() error {
	return validator.New().Struct(s)
}

func (Suit) TableName() string {
	return "suit"
}

func (s *Suit) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.NewV4().String()
	return
}
