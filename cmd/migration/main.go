package main

import (
	"log"

	"github.com/gustavocioccari/toggl-cards/pkg/models"
	"github.com/gustavocioccari/toggl-cards/pkg/repositories/postgres"
	"github.com/joho/godotenv"
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

	err = db.AutoMigrate(&models.Suit{}, &models.Value{}, &models.Card{}, &models.Deck{})
	if err != nil {
		log.Fatal(err)
	}
}
