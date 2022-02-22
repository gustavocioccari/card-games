package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gustavocioccari/toggl-cards/pkg/repositories/postgres"
	cardRepo "github.com/gustavocioccari/toggl-cards/pkg/repositories/postgres/card"
	deckRepo "github.com/gustavocioccari/toggl-cards/pkg/repositories/postgres/deck"
	"github.com/gustavocioccari/toggl-cards/pkg/services/deck"
	"github.com/gustavocioccari/toggl-cards/pkg/ui/rest/router"
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
		log.Fatal(err)
	}

	deckRepository := deckRepo.NewDeckRepository(db)
	cardRepository := cardRepo.NewCardRepository(db)

	deckService := deck.NewDeckService(deckRepository, cardRepository)

	router := router.SetupRouter(deckService)

	if err := router.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))); err != nil {
		log.Fatalln("Error on start rest:", err)
	}
}
