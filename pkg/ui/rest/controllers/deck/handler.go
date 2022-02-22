package deck

import (
	"net/http"
	"strconv"

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
	GetByID(c echo.Context) error
	DrawCard(c echo.Context) error
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

	res := map[string]interface{}{
		"id":        deck.ID,
		"shuffled":  deck.Shuffled,
		"remaining": deck.Remaining,
	}

	return c.JSON(http.StatusCreated, res)
}

func (dc *deckController) GetByID(c echo.Context) error {
	id := c.Param("id")

	deck, err := dc.deckService.GetByID(id)
	if err != nil {
		return rest.NotFound(c, err)
	}

	return c.JSON(http.StatusOK, deck)
}

func (dc *deckController) DrawCard(c echo.Context) error {
	id := c.Param("id")
	limit := c.QueryParam("limit")
	limitInt := 1

	if limit != "" {
		var err error

		limitInt, err = strconv.Atoi(limit)
		if err != nil {
			return rest.InternalServerError(c, err)
		}
	}

	cards, err := dc.deckService.DrawCard(id, limitInt)
	if err != nil {
		return rest.InternalServerError(c, err)
	}

	return c.JSON(http.StatusOK, cards)
}
