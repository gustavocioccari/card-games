package rest

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gustavocioccari/toggl-cards/pkg/models"
	"github.com/labstack/echo/v4"
)

func InternalServerError(c echo.Context, err error) error {
	return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
}

func NotFound(c echo.Context, data interface{}) error {
	err, match := data.(error)
	if match {
		data = echo.Map{"error": err.Error()}
	}

	return c.JSON(http.StatusNotFound, data)
}

func BadRequest(c echo.Context, data interface{}) error {
	err, match := data.(error)
	if match {
		data = echo.Map{"error": err.Error()}
	}

	return c.JSON(http.StatusBadRequest, data)
}

func ValidationError(c echo.Context, err error) error {
	var errors []models.ValidationError

	for _, e := range err.(validator.ValidationErrors) {
		log.Println(e)

		var element models.ValidationError

		element.Field = e.StructNamespace()
		element.Tag = e.Tag()
		element.Value = e.Param()

		errors = append(errors, element)
	}

	return c.JSON(http.StatusBadRequest, echo.Map{"errors": errors})
}
