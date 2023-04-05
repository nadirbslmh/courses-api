package users

import (
	"courses-api/businesses/users"
	"courses-api/controllers"
	"courses-api/controllers/users/request"
	"courses-api/controllers/users/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authUseCase users.Usecase
}

func NewAuthController(authUC users.Usecase) *AuthController {
	return &AuthController{
		authUseCase: authUC,
	}
}

func (ctrl *AuthController) Register(c echo.Context) error {
	userInput := request.User{}
	ctx := c.Request().Context()

	if err := c.Bind(&userInput); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}

	err := userInput.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	user, err := ctrl.authUseCase.Register(ctx, userInput.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "error when inserting data", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "user registered", response.FromDomain(user))
}

func (ctrl *AuthController) Login(c echo.Context) error {
	userInput := request.User{}
	ctx := c.Request().Context()

	if err := c.Bind(&userInput); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}

	err := userInput.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	token, err := ctrl.authUseCase.Login(ctx, userInput.ToDomain())

	var isFailed bool = err != nil || token == ""

	if isFailed {
		return controllers.NewResponse(c, http.StatusUnauthorized, "failed", "invalid email or password", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "token created", token)
}
