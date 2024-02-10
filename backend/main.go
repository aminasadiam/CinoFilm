package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aminasadiam/CinoFilm/backend/database"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Cant Fine .Env File")
		return
	}
	database.ConnectDatabase()
}

func main() {
	app := echo.New()

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	app.Logger.Fatal(app.Start(os.Getenv("PORT")))
}
