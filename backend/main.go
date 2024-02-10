package main

import (
	"log"
	"os"

	"github.com/aminasadiam/CinoFilm/backend/controllers"
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

	app.GET("/", controllers.Hello)

	app.Logger.Fatal(app.Start(os.Getenv("PORT")))
}
