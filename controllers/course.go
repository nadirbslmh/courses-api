package controllers

import (
	"courses-api/models"
	"courses-api/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CourseController struct {
	service services.CourseService
}

func InitCourseController() CourseController {
	return CourseController{
		service: services.InitCourseService(),
	}
}

func (cc *CourseController) GetAll(c echo.Context) error {
	courses, err := cc.service.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to fetch courses data",
		})
	}

	return c.JSON(http.StatusOK, models.Response[[]models.Course]{
		Status:  "success",
		Message: "all courses",
		Data:    courses,
	})
}

func (cc *CourseController) GetByID(c echo.Context) error {
	var courseID string = c.Param("id")

	course, err := cc.service.GetByID(courseID)

	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response[string]{
			Status:  "failed",
			Message: "course not found",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Course]{
		Status:  "success",
		Message: "course found",
		Data:    course,
	})
}

func (cc *CourseController) Create(c echo.Context) error {
	var courseInput models.CourseInput

	if err := c.Bind(&courseInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	course, err := cc.service.Create(courseInput)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, models.Response[models.Course]{
		Status:  "success",
		Message: "course created",
		Data:    course,
	})
}

func (cc *CourseController) Update(c echo.Context) error {
	var courseID string = c.Param("id")

	var courseInput models.CourseInput

	if err := c.Bind(&courseInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	course, err := cc.service.Update(courseInput, courseID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Course]{
		Status:  "success",
		Message: "course updated",
		Data:    course,
	})
}

func (cc *CourseController) Delete(c echo.Context) error {
	var courseID string = c.Param("id")

	err := cc.service.Delete(courseID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response[string]{
		Status:  "success",
		Message: "course deleted",
	})
}

func (cc *CourseController) Restore(c echo.Context) error {
	var courseID string = c.Param("id")

	course, err := cc.service.Restore(courseID)

	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response[string]{
			Status:  "failed",
			Message: "course not found",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Course]{
		Status:  "success",
		Message: "course restored",
		Data:    course,
	})
}

func (cc *CourseController) ForceDelete(c echo.Context) error {
	var courseID string = c.Param("id")

	err := cc.service.ForceDelete(courseID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response[string]{
		Status:  "success",
		Message: "course deleted permanently",
	})
}
