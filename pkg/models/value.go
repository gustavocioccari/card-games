package models

import (
	"github.com/go-playground/validator"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Value struct {
	ID   string `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"unique,not null" validate:"required"`
}

func (v *Value) Validate() error {
	return validator.New().Struct(v)
}

func (Value) TableName() string {
	return "value"
}

func (v *Value) BeforeCreate(tx *gorm.DB) (err error) {
	v.ID = uuid.NewV4().String()
	return
}
