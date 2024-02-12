package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/aminasadiam/CinoFilm/backend/database"
	"github.com/aminasadiam/CinoFilm/backend/models"
	"github.com/aminasadiam/CinoFilm/backend/security/password"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/nanorand/nanorand"
)

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

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

func Login(c echo.Context) error {
	data := new(LoginData)
	if err := c.Bind(data); err != nil {
		return c.String(400, "Bad request")
	}

	var user models.User
	database.DB.Where("username = ?", data.Username).First(&user)
	if user.ID == 0 {
		return c.JSON(http.StatusNotFound, "user not found")
	}

	if !user.IsActive {
		return c.JSON(http.StatusNotFound, "user is not active")
	}

	if compare := password.CheckPasswordHash(data.Password, user.Password); !compare {
		return c.JSON(http.StatusNotFound, "incorrect password")
	}

	cliam := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := cliam.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to login : "+err.Error())
	}

	cookie := http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}

	c.SetCookie(&cookie)

	return c.JSON(http.StatusOK, "success")
}
