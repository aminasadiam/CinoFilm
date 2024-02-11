package main

import (
	"log"
	"os"

	"github.com/aminasadiam/CinoFilm/backend/database"
	"github.com/aminasadiam/CinoFilm/backend/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
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
	routes.SetupApp(app)
	app.Logger.Fatal(app.Start(os.Getenv("PORT")))
}
