package main

import (
	"log"

	"github.com/gustavocioccari/toggl-cards/pkg/models"
	"github.com/gustavocioccari/toggl-cards/pkg/repositories/postgres"
	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}

func main() {
	db, err := postgres.GetDB()
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	if err = suit(db); err != nil {
		log.Fatal("suit", err)
	}

	if err = value(db); err != nil {
		log.Fatal("value", err)
	}

	if err = card(db); err != nil {
		log.Fatal("card", err)
	}

	log.Println("finish seed")
}

func suit(db *gorm.DB) error {
	suits := []models.Suit{
		{
			ID:   uuid.NewV4().String(),
			Name: "CLUBS",
		},
		{
			ID:   uuid.NewV4().String(),
			Name: "DIAMONDS",
		},
		{
			ID:   uuid.NewV4().String(),
			Name: "HEARTS",
		},
		{
			ID:   uuid.NewV4().String(),
			Name: "SPADES",
		},
	}

	return db.Save(&suits).Error
}

func value(db *gorm.DB) error {
	value := []models.Value{
		{
			ID:   uuid.NewV4().String(),
			Name: "ACE",
		},
		{
			ID:   uuid.NewV4().String(),
			Name: "KING",
		},
		{
			ID:   uuid.NewV4().String(),
			Name: "QUEEN",
		},
		{
			ID:   uuid.NewV4().String(),
			Name: "JACK",
		},
		{
			ID:   uuid.NewV4().String(),
			Name: "10",
		},
		{
			ID:   uuid.NewV4().String(),
			Name: "9",
		},
		{
			ID:   uuid.NewV4().String(),
			Name: "8",
		},
		{
			ID:   uuid.NewV4().String(),
			Name: "7",
		},
		{
			ID:   uuid.NewV4().String(),
			Name: "6",
		},
		{
			ID:   uuid.NewV4().String(),
			Name: "5",
		},
		{
			ID:   uuid.NewV4().String(),
			Name: "4",
		},
		{
			ID:   uuid.NewV4().String(),
			Name: "3",
		},
		{
			ID:   uuid.NewV4().String(),
			Name: "2",
		},
	}

	return db.Save(&value).Error
}

func card(db *gorm.DB) error {
	var cards []models.Card

	var values []models.Value
	var suits []models.Suit

	err := db.Find(&values).Error
	if err != nil {
		log.Fatal("Error finding values", err)
	}

	err = db.Find(&suits).Error
	if err != nil {
		log.Fatal("Error finding suits", err)
	}

	for _, suit := range suits {
		for _, value := range values {
			card := models.Card{
				ID:      uuid.NewV4().String(),
				SuitID:  suit.ID,
				ValueID: value.ID,
				Code:    string(value.Name[0]) + string(suit.Name[0]),
			}

			cards = append(cards, card)
		}
	}
	return db.Save(&cards).Error
}
