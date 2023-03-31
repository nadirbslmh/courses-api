package routes

import (
	"courses-api/controllers"
	"courses-api/middlewares"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	courseController := controllers.InitCourseController()

	loggerConfig := middlewares.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}

	loggerMiddleware := loggerConfig.Init()

	e.Use(loggerMiddleware)

	courseRoutes := e.Group("/api/v1")

	courseRoutes.GET("/courses", courseController.GetAll)
	courseRoutes.GET("/courses/:id", courseController.GetByID)
	courseRoutes.POST("/courses", courseController.Create)
	courseRoutes.PUT("/courses/:id", courseController.Update)
	courseRoutes.DELETE("/courses/:id", courseController.Delete)
	courseRoutes.POST("/courses/:id", courseController.Restore)
	courseRoutes.DELETE("/courses/:id/force", courseController.ForceDelete)
}
