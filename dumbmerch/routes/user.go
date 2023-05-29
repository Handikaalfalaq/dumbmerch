package routes

import (
	"dumbmerch/handlers"
	middleware "dumbmerch/pkg/middleware"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repositories"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	e.GET("/users", middleware.Auth(h.GetAllUser))
	e.GET("/user/:id", h.GetUserById)
	e.POST("/user", h.CreateNewUser)
	e.PATCH("/user/:id", h.UpdateDataUser)
	e.DELETE("/user/:id", h.DeleteDataUser)
}
