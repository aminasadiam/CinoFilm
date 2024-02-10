package main

import (
	"github.com/aminasadiam/CinoFilm/backend/database"
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
	database.DB.AutoMigrate()
}
