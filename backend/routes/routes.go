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
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowHeaders:     []string{"authorization", "Content-Type"},
		AllowCredentials: true,
		AllowMethods:     []string{echo.OPTIONS, echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	authRoutes(app)
}

func authRoutes(app *echo.Echo) {
	app.POST("/api/register", controllers.Register)
	app.POST("/api/login", controllers.Login)
	app.GET("/api/user", controllers.User)
	app.POST("/api/logout", controllers.Logout)
}
