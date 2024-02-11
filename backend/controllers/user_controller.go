package controllers

import (
	"net/http"
	"strconv"

	"github.com/aminasadiam/CinoFilm/backend/database"
	"github.com/aminasadiam/CinoFilm/backend/models"
	"github.com/aminasadiam/CinoFilm/backend/security/password"
	"github.com/labstack/echo/v4"
	"github.com/nanorand/nanorand"
)

func Register(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.String(400, "Bad request")
	}

	user.Password, _ = password.HashPassword(user.Password)

	code, _ := nanorand.Gen(6)
	i, _ := strconv.Atoi(code)
	user.ActiveCode = uint(i)

	database.DB.Create(&user)

	return c.JSON(http.StatusOK, user)
}
