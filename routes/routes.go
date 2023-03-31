package routes

import (
	"courses-api/controllers"
	"courses-api/middlewares"
	"courses-api/utils"

	echojwt "github.com/labstack/echo-jwt/v4"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	loggerConfig := middlewares.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}

	loggerMiddleware := loggerConfig.Init()

	e.Use(loggerMiddleware)

	jwtConfig := middlewares.JWTConfig{
		SecretKey:       utils.GetConfig("JWT_SECRET_KEY"),
		ExpiresDuration: 1,
	}

	authMiddlewareConfig := jwtConfig.Init()

	userController := controllers.InitUserController(&jwtConfig)

	users := e.Group("/api/v1/users")

	users.POST("/register", userController.Register)
	users.POST("/login", userController.Login)

	courseController := controllers.InitCourseController()

	courseRoutes := e.Group("/api/v1", echojwt.WithConfig(authMiddlewareConfig))
	courseRoutes.Use(middlewares.VerifyToken)

	courseRoutes.GET("/courses", courseController.GetAll)
	courseRoutes.GET("/courses/:id", courseController.GetByID)
	courseRoutes.POST("/courses", courseController.Create)
	courseRoutes.PUT("/courses/:id", courseController.Update)
	courseRoutes.DELETE("/courses/:id", courseController.Delete)
	courseRoutes.POST("/courses/:id", courseController.Restore)
	courseRoutes.DELETE("/courses/:id/force", courseController.ForceDelete)
}
