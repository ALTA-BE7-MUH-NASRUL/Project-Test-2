package task

import (
	"Tugas/Project-Test-2/delivery/helper"
	_middlewares "Tugas/Project-Test-2/delivery/middleware"
	_taskUseCase "Tugas/Project-Test-2/usecase/task"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	taskUseCase _taskUseCase.TaskUseCaseInterface
}

func NewTaskHandler(taskUseCase _taskUseCase.TaskUseCaseInterface) *TaskHandler {
	return &TaskHandler{
		taskUseCase: taskUseCase,
	}
}

func (th *TaskHandler) GetTaskHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, tokenerr := _middlewares.ReadTokenId(c)
		if tokenerr != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Bad Request"))
		}
		idn := c.Param("id")
		id, _ := strconv.Atoi(idn)
		task, rows, err := th.taskUseCase.GetTask(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not exist"))
		}
		if idToken != id {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Bad Request"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get data", task))
	}
}

func (th *TaskHandler) DeleteTaskHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, tokenerr := _middlewares.ReadTokenId(c)
		if tokenerr != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Bad Request"))
		}
		idn := c.Param("id")
		id, _ := strconv.Atoi(idn)
		task, rows, err := th.taskUseCase.GetTask(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Data not exist"))
		}
		if idToken != int(task.UserID) {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Bad Request"))
		}
		_, err = th.taskUseCase.DeleteTask(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed delete data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("Succes delete data"))
	}
}

func (th *TaskHandler) CreateTaskHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		type task struct {
			UserId    int    `json:"UserId"`
			ProjectId int    `json:"ProjectId"`
			List      string `json:"list"`
		}
		var tasks task
		idToken, tokenerr := _middlewares.ReadTokenId(c)
		if tokenerr != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Bad Request"))
		}
		fmt.Println(tasks)
		c.Bind(&tasks)
		if idToken != tasks.UserId {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Bad Request"))
		}
		tax, err := th.taskUseCase.CreateTask(tasks.UserId, tasks.ProjectId, tasks.List)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed create data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("Succes create data", tax))
	}
}

func (th *TaskHandler) UpdatedTaskHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, tokenerr := _middlewares.ReadTokenId(c)
		if tokenerr != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Bad Request"))
		}
		idn := c.Param("id")
		id, _ := strconv.Atoi(idn)
		task, rows, err := th.taskUseCase.GetTask(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Data not exist"))
		}
		if idToken != int(task.UserID) {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Bad Request"))
		}
		c.Bind(&task)
		task, err = th.taskUseCase.UpdatedTask(task, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed edit data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("Success edit data", task))
	}
}
func (th *TaskHandler) CompletedTaskHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, tokenerr := _middlewares.ReadTokenId(c)
		if tokenerr != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Bad Request"))
		}
		idn := c.Param("id")
		id, _ := strconv.Atoi(idn)
		task, rows, err := th.taskUseCase.GetTask(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Data not exist"))
		}
		if idToken != int(task.UserID) {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Bad Request"))
		}
		c.Bind(&task)
		task, err = th.taskUseCase.CompletedTask(task, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed edit data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("Success edit data", task))
	}
}
func (th *TaskHandler) ReopenTaskHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, tokenerr := _middlewares.ReadTokenId(c)
		if tokenerr != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Bad Request"))
		}
		idn := c.Param("id")
		id, _ := strconv.Atoi(idn)
		task, rows, err := th.taskUseCase.GetTask(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Data not exist"))
		}
		if idToken != int(task.UserID) {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Bad Request"))
		}
		c.Bind(&task)
		task, err = th.taskUseCase.ReopenTask(task, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed edit data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("Success edit data", task))
	}
}
