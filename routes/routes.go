package routes

import (
	"courses-api/controllers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	courseController := controllers.InitCourseController()

	courseRoutes := e.Group("/api/v1")

	courseRoutes.GET("/courses", courseController.GetAll)
	courseRoutes.GET("/courses/:id", courseController.GetByID)
	courseRoutes.POST("/courses", courseController.Create)
	courseRoutes.PUT("/courses/:id", courseController.Update)
	courseRoutes.DELETE("/courses/:id", courseController.Delete)
}
