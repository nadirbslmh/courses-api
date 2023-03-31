package controllers

import (
	"courses-api/middlewares"
	"courses-api/models"
	"courses-api/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	service services.UserService
}

func InitUserController(jwtAuth *middlewares.JWTConfig) UserController {
	return UserController{
		service: services.InitUserService(jwtAuth),
	}
}

func (uc *UserController) Register(c echo.Context) error {
	var userInput models.UserInput

	if err := c.Bind(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	err := userInput.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	user, err := uc.service.Register(userInput)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, models.Response[models.User]{
		Status:  "success",
		Message: "user registered",
		Data:    user,
	})
}

func (uc *UserController) Login(c echo.Context) error {
	var userInput models.UserInput

	if err := c.Bind(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	err := userInput.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	token, err := uc.service.Login(userInput)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "invalid email or password",
		})
	}

	return c.JSON(http.StatusOK, models.Response[string]{
		Status:  "success",
		Message: "login success",
		Data:    token,
	})
}
