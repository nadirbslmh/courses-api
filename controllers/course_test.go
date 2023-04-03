package controllers

import (
	"bytes"
	"courses-api/database"
	"courses-api/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name                   string
	path                   string
	expectedStatus         int
	expectedBodyStartsWith string
}

var controller CourseController = InitCourseController()

func InitEcho() *echo.Echo {
	database.InitDatabase()
	database.Migrate()

	e := echo.New()

	return e
}

func TestGetAllCourses_Success(t *testing.T) {
	testcase := testCase{
		name:                   "success",
		path:                   "/api/v1/courses",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitEcho()

	req := httptest.NewRequest(http.MethodGet, testcase.path, nil)

	recorder := httptest.NewRecorder()

	ctx := e.NewContext(req, recorder)

	ctx.SetPath(testcase.path)

	if assert.NoError(t, controller.GetAll(ctx)) {
		assert.Equal(t, http.StatusOK, recorder.Code)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}

}

func TestCreateCourse_Success(t *testing.T) {
	testcase := testCase{
		name:                   "success",
		path:                   "/api/v1/courses",
		expectedStatus:         http.StatusCreated,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitEcho()

	category, err := database.SeedCategory()

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	var courseInput models.CourseInput = models.CourseInput{
		Title:       "test",
		Description: "test",
		CategoryID:  category.ID,
		Level:       "test",
	}

	jsonBody, err := json.Marshal(&courseInput)

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	bodyReader := bytes.NewReader(jsonBody)

	request := httptest.NewRequest(http.MethodPost, testcase.path, bodyReader)

	recorder := httptest.NewRecorder()

	request.Header.Add("Content-Type", "application/json")

	ctx := e.NewContext(request, recorder)

	ctx.SetPath(testcase.path)

	if assert.NoError(t, controller.Create(ctx)) {
		assert.Equal(t, http.StatusCreated, testcase.expectedStatus)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}

	t.Cleanup(func() {
		database.CleanSeeders()
	})
}

func TestGetCourseByID_Success(t *testing.T) {
	testcase := testCase{
		name:                   "success",
		path:                   "/api/v1/courses",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitEcho()

	course, err := database.SeedCourse()

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	courseID := strconv.Itoa(int(course.ID))

	request := httptest.NewRequest(http.MethodGet, testcase.path, nil)

	recorder := httptest.NewRecorder()

	ctx := e.NewContext(request, recorder)

	ctx.SetPath(testcase.path)

	ctx.SetParamNames("id")
	ctx.SetParamValues(courseID)

	if assert.NoError(t, controller.GetByID(ctx)) {
		assert.Equal(t, http.StatusOK, testcase.expectedStatus)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}

	t.Cleanup(func() {
		database.CleanSeeders()
	})
}
