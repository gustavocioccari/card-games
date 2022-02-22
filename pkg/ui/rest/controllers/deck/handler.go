package deck

import (
	"net/http"

	"github.com/gustavocioccari/toggl-cards/pkg/models"
	"github.com/gustavocioccari/toggl-cards/pkg/services/deck"
	"github.com/gustavocioccari/toggl-cards/pkg/ui/rest"
	"github.com/labstack/echo/v4"
)

type deckController struct {
	deckService deck.DeckService
}

type DeckController interface {
	Create(c echo.Context) error
}

func NewDeckController(deckService deck.DeckService) DeckController {
	return &deckController{
		deckService: deckService,
	}
}

func (dc *deckController) Create(c echo.Context) error {
	var deck *models.Deck

	shuffled := c.QueryParam("shuffled")
	cards := c.QueryParam("cards")

	if (shuffled == "false" || shuffled == "") && cards == "" {
		deck, err := dc.deckService.CreateDefault()
		if err != nil {
			return rest.InternalServerError(c, err)
		}
		return c.JSON(http.StatusCreated, deck)
	}

	deck, err := dc.deckService.Create(cards, shuffled)
	if err != nil {
		return rest.InternalServerError(c, err)
	}

	return c.JSON(http.StatusCreated, deck)
}
