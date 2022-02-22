package deck

import (
	"log"
	"testing"

	"github.com/gustavocioccari/toggl-cards/pkg/repositories/postgres"
	"github.com/gustavocioccari/toggl-cards/pkg/repositories/postgres/card"
	"github.com/gustavocioccari/toggl-cards/pkg/repositories/postgres/deck"
	"github.com/stretchr/testify/assert"
)

func TestDeck(t *testing.T) {
	db, err := postgres.GetDB()
	if err != nil {
		log.Fatal(err)
	}

	deckRepository := deck.NewDeckRepository(db)
	cardRepository := card.NewCardRepository(db)
	deckService := NewDeckService(deckRepository, cardRepository)

	t.Run("should be able to create a shuffled deck with default cards", func(t *testing.T) {
		d, err := deckService.Create("", "true")

		if assert.NoError(t, err) {
			assert.Equal(t, d.Remaining, 52)
			assert.Equal(t, d.Shuffled, true)
		}
	})

	t.Run("should be able to create a shuffled deck with selected cards", func(t *testing.T) {
		d, err := deckService.Create("AS,JC,10H,5C,6D", "true")

		if assert.NoError(t, err) {
			assert.Equal(t, d.Remaining, 5)
			assert.Equal(t, d.Shuffled, true)
		}
	})

	t.Run("should be able to create an unshuffled deck with default cards", func(t *testing.T) {
		d, err := deckService.Create("", "false")

		if assert.NoError(t, err) {
			assert.Equal(t, d.Remaining, 52)
			assert.Equal(t, d.Shuffled, false)
		}
	})

	t.Run("should be able to create an unshuffled deck with selected cards", func(t *testing.T) {
		d, err := deckService.Create("AC,JC", "false")

		if assert.NoError(t, err) {
			assert.Equal(t, d.Remaining, 2)
			assert.Equal(t, d.Shuffled, false)
		}
	})
}
