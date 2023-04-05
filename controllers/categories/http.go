package categories

import (
	"courses-api/businesses/categories"
	"courses-api/controllers"
	"courses-api/controllers/categories/request"
	"courses-api/controllers/categories/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	categoryUsecase categories.Usecase
}

func NewCategoryController(categoryUC categories.Usecase) *CategoryController {
	return &CategoryController{
		categoryUsecase: categoryUC,
	}
}

func (cc *CategoryController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	categoriesData, err := cc.categoryUsecase.GetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to fetch data", "")
	}

	categories := []response.Category{}

	for _, category := range categoriesData {
		categories = append(categories, response.FromDomain(category))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all categories", categories)
}

func (cc *CategoryController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	categoryID := c.Param("id")

	category, err := cc.categoryUsecase.GetByID(ctx, categoryID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "category not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "category found", response.FromDomain(category))
}

func (cc *CategoryController) Create(c echo.Context) error {
	input := request.Category{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	category, err := cc.categoryUsecase.Create(ctx, input.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to create a category", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "category created", response.FromDomain(category))
}

func (cc *CategoryController) Update(c echo.Context) error {
	input := request.Category{}
	ctx := c.Request().Context()

	categoryID := c.Param("id")

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	category, err := cc.categoryUsecase.Update(ctx, input.ToDomain(), categoryID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to update a category", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "category updated", response.FromDomain(category))
}

func (cc *CategoryController) Delete(c echo.Context) error {
	categoryID := c.Param("id")
	ctx := c.Request().Context()

	err := cc.categoryUsecase.Delete(ctx, categoryID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to delete a category", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "category deleted", "")
}
