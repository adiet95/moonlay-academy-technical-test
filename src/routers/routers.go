package routers

import (
	"errors"

	"github.com/adiet95/moonlay-academy-technical-test/src/database"
	"github.com/labstack/echo/v4"
)

func New() (*echo.Echo, error) {
	mainRoute := echo.New()

	db, err := database.New()
	if err != nil {
		return nil, errors.New("failed init database")
	}

	// modules.New(mainRoute, db)

	return mainRoute, nil
}
