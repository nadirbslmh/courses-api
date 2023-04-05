package courses

import (
	"courses-api/businesses/courses"
	"courses-api/controllers"
	"courses-api/controllers/courses/request"
	"courses-api/controllers/courses/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CourseController struct {
	courseUsecase courses.Usecase
}

func NewCourseController(courseUC courses.Usecase) *CourseController {
	return &CourseController{
		courseUsecase: courseUC,
	}
}

func (cc *CourseController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	coursesData, err := cc.courseUsecase.GetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to fetch data", "")
	}

	courses := []response.Course{}

	for _, course := range coursesData {
		courses = append(courses, response.FromDomain(course))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all courses", courses)
}

func (cc *CourseController) GetByID(c echo.Context) error {
	var courseID string = c.Param("id")
	ctx := c.Request().Context()

	course, err := cc.courseUsecase.GetByID(ctx, courseID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "course not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "course found", response.FromDomain(course))
}

func (cc *CourseController) Create(c echo.Context) error {
	input := request.Course{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	course, err := cc.courseUsecase.Create(ctx, input.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to create a course", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "course created", response.FromDomain(course))
}

func (cc *CourseController) Update(c echo.Context) error {
	var courseID string = c.Param("id")
	ctx := c.Request().Context()

	input := request.Course{}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	course, err := cc.courseUsecase.Update(ctx, input.ToDomain(), courseID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "update course failed", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "course updated", response.FromDomain(course))
}

func (cc *CourseController) Delete(c echo.Context) error {
	var courseID string = c.Param("id")
	ctx := c.Request().Context()

	err := cc.courseUsecase.Delete(ctx, courseID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "delete course failed", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "course deleted", "")
}

func (cc *CourseController) Restore(c echo.Context) error {
	var courseID string = c.Param("id")
	ctx := c.Request().Context()

	course, err := cc.courseUsecase.Restore(ctx, courseID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "course not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "course restored", response.FromDomain(course))
}

func (cc *CourseController) ForceDelete(c echo.Context) error {
	var courseID string = c.Param("id")
	ctx := c.Request().Context()

	err := cc.courseUsecase.ForceDelete(ctx, courseID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "force delete course failed", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "course deleted permanently", "")
}
