package routes

import (
	"github.com/aminasadiam/CinoFilm/backend/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupApp(app *echo.Echo) {
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
	}))

	authRoutes(app)
}

func authRoutes(app *echo.Echo) {
	app.POST("/api/register", controllers.Register)
	app.POST("/api/login", controllers.Login)
}
