package project

import (
	"Tugas/Project-Test-2/delivery/helper"
	_middlewares "Tugas/Project-Test-2/delivery/middleware"
	_entities "Tugas/Project-Test-2/entities"
	_projectUseCase "Tugas/Project-Test-2/usecase/project"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProjectHandler struct {
	projectUseCase _projectUseCase.ProjectUseCaseInterface
}

func NewProjectHandler(projectUseCase _projectUseCase.ProjectUseCaseInterface) *ProjectHandler {
	return &ProjectHandler{
		projectUseCase: projectUseCase,
	}
}

func (th *ProjectHandler) GetProjectHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, tokenerr := _middlewares.ReadTokenId(c)
		if tokenerr != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Bad Request"))
		}
		idn := c.Param("id")
		id, _ := strconv.Atoi(idn)
		Project, rows, err := th.projectUseCase.GetProject(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not exist"))
		}
		if idToken != int(Project.UserID) {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Bad Request"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get data", Project))
	}
}

func (th *ProjectHandler) DeleteProjectHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, tokenerr := _middlewares.ReadTokenId(c)
		if tokenerr != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Bad Request"))
		}
		idn := c.Param("id")
		id, _ := strconv.Atoi(idn)
		Project, rows, err := th.projectUseCase.GetProject(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Data not exist"))
		}
		if idToken != int(Project.UserID) {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Bad Request"))
		}
		_, err = th.projectUseCase.DeleteProject(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed delete data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("Succes delete data"))
	}
}

func (th *ProjectHandler) CreateProjectHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var Project _entities.Project
		c.Bind(&Project)
		Project, err := th.projectUseCase.CreateProject(Project)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed create data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("Succes create data", Project))
	}
}

func (th *ProjectHandler) UpdatedProjectHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, tokenerr := _middlewares.ReadTokenId(c)
		if tokenerr != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Bad Request"))
		}
		idn := c.Param("id")
		id, _ := strconv.Atoi(idn)
		Project, rows, err := th.projectUseCase.GetProject(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Data not exist"))
		}
		if idToken != int(Project.UserID) {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Bad Request"))
		}
		c.Bind(&Project)
		Project, err = th.projectUseCase.UpdatedProject(Project, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed edit data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("Success edit data", Project))
	}
}
func (th *ProjectHandler) CompletedProjectHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, tokenerr := _middlewares.ReadTokenId(c)
		if tokenerr != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Bad Request"))
		}
		idn := c.Param("id")
		id, _ := strconv.Atoi(idn)
		Project, rows, err := th.projectUseCase.GetProject(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Data not exist"))
		}
		if idToken != int(Project.UserID) {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Bad Request"))
		}
		c.Bind(&Project)
		Project, err = th.projectUseCase.CompletedProject(Project, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed edit data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("Success edit data", Project))
	}
}
