package main

import (
	"github.com/aminasadiam/CinoFilm/backend/database"
	"github.com/aminasadiam/CinoFilm/backend/models"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	database.ConnectDatabase()
}

func main() {
	database.DB.AutoMigrate(&models.User{})
}
