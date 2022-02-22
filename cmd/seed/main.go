package main

import (
	"log"
	"strconv"

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
			Name: "SPADES",
			Rank: 1,
		},
		{
			ID:   uuid.NewV4().String(),
			Name: "DIAMONDS",
			Rank: 2,
		},
		{
			ID:   uuid.NewV4().String(),
			Name: "CLUBS",
			Rank: 3,
		},
		{
			ID:   uuid.NewV4().String(),
			Name: "HEARTS",
			Rank: 4,
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
			Name: "2",
		},
		{
			ID:   uuid.NewV4().String(),
			Name: "3",
		},
		{
			ID:   uuid.NewV4().String(),
			Name: "4",
		},
		{
			ID:   uuid.NewV4().String(),
			Name: "5",
		},
		{
			ID:   uuid.NewV4().String(),
			Name: "6",
		},
		{
			ID:   uuid.NewV4().String(),
			Name: "7",
		},
		{
			ID:   uuid.NewV4().String(),
			Name: "8",
		},
		{
			ID:   uuid.NewV4().String(),
			Name: "9",
		},
		{
			ID:   uuid.NewV4().String(),
			Name: "10",
		},
		{
			ID:   uuid.NewV4().String(),
			Name: "JACK",
		},
		{
			ID:   uuid.NewV4().String(),
			Name: "QUEEN",
		},
		{
			ID:   uuid.NewV4().String(),
			Name: "KING",
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
			code := string(value.Name[0]) + string(suit.Name[0])

			if len(value.Name) > 1 {
				if _, err := strconv.Atoi(string(value.Name[1])); err == nil {
					code = string(value.Name[0]) + string(value.Name[1]) + string(suit.Name[0])
				}
			}

			card := models.Card{
				ID:      uuid.NewV4().String(),
				SuitID:  suit.ID,
				ValueID: value.ID,
				Code:    code,
			}

			cards = append(cards, card)
		}
	}
	return db.Save(&cards).Error
}
