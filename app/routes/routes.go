package routes

import (
	"courses-api/app/middlewares"
	"courses-api/controllers/categories"
	"courses-api/controllers/courses"
	"courses-api/controllers/users"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	LoggerMiddleware   echo.MiddlewareFunc
	JWTMiddleware      echojwt.Config
	AuthController     users.AuthController
	CourseController   courses.CourseController
	CategoryController categories.CategoryController
}

func (cl *ControllerList) RegisterRoutes(e *echo.Echo) {
	e.Use(cl.LoggerMiddleware)

	users := e.Group("/api/v1/users")

	users.POST("/register", cl.AuthController.Register)
	users.POST("/login", cl.AuthController.Login)

	course := e.Group("/api/v1/courses", echojwt.WithConfig(cl.JWTMiddleware))
	course.Use(middlewares.VerifyToken)

	course.GET("", cl.CourseController.GetAll)
	course.GET("/:id", cl.CourseController.GetByID)
	course.POST("", cl.CourseController.Create)
	course.PUT("/:id", cl.CourseController.Update)
	course.DELETE("/:id", cl.CourseController.Delete)
	course.POST("/:id", cl.CourseController.Restore)
	course.DELETE("/:id/force", cl.CourseController.ForceDelete)

	category := e.Group("/api/v1/categories", echojwt.WithConfig(cl.JWTMiddleware))
	category.Use(middlewares.VerifyToken)

	category.GET("", cl.CategoryController.GetAll)
	category.POST("", cl.CategoryController.Create)
	category.PUT("/:id", cl.CategoryController.Update)
	category.DELETE("/:id", cl.CategoryController.Delete)
}
