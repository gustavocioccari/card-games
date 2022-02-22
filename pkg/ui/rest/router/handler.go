package router

import (
	"net/http"

	deckService "github.com/gustavocioccari/toggl-cards/pkg/services/deck"
	deckController "github.com/gustavocioccari/toggl-cards/pkg/ui/rest/controllers/deck"
	"github.com/gustavocioccari/toggl-cards/pkg/ui/rest/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRouter(deckService deckService.DeckService) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middlewares.Cors())
	e.Use(middlewares.Logger())

	deckController := deckController.NewDeckController(deckService)

	v1 := e.Group("toggl-cards/v1")
	{
		v1.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "Server is running")
		})

		groupDeck := v1.Group("/decks")
		{
			groupDeck.POST("", deckController.Create)
			groupDeck.GET("/:id/open", deckController.GetByID)
			groupDeck.PATCH("/:id/draw", deckController.DrawCard)
		}
	}

	return e
}
