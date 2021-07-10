package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"

	"golang-testing/domain"
	"golang-testing/domain/interfaces"
	"golang-testing/internal/delivery/http/helper"
)

type ResponseError struct {
	Message string `json:"message"`
}

type ProjectHandler struct {
	PUseCase interfaces.ProjectUsecase
}

func NewProjectHandler(e *echo.Echo, us interfaces.ProjectUsecase) {
	handler := &ProjectHandler{
		PUseCase: us,
	}
	route := e.Group("/api")
	route.GET("/projects", handler.GetAll)
	route.GET("/projects/:id", handler.GetByID)
	route.POST("/projects", handler.Create)
	route.PUT("/projects/:id", handler.Update)
	route.DELETE("/projects/:id", handler.Delete)
}

func (a *ProjectHandler) GetAll(c echo.Context) error {

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	listAr, err := a.PUseCase.GetAll(ctx)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, listAr)
}

func (a *ProjectHandler) GetByID(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	art, err := a.PUseCase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	helper.RespondSuccess(c, http.StatusOK, art)
	return nil
}

func isRequestValid(m *domain.Project) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *ProjectHandler) Create(c echo.Context) error {
	var project domain.Project

	err := c.Bind(&project)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&project); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	err = project.Validate()
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	err = a.PUseCase.Create(ctx, &project)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, project)
}

func (a *ProjectHandler) Update(c echo.Context) error {
	var project domain.Project

	projectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	project.ID = int64(projectID)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = a.PUseCase.Update(ctx, &project)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, project)
}

func (a *ProjectHandler) Delete(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}
	id := int64(idP)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = a.PUseCase.Delete(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, "Success delete item..")
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
