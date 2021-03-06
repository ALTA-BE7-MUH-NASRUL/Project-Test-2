package user

import (
	"Tugas/Project-Test-2/delivery/helper"
	_middlewares "Tugas/Project-Test-2/delivery/middleware"
	_entities "Tugas/Project-Test-2/entities"
	_userUseCase "Tugas/Project-Test-2/usecase/user"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUseCase _userUseCase.UserUseCaseInterface
}

func NewUserHandler(userUseCase _userUseCase.UserUseCaseInterface) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

func (uh *UserHandler) GetUserHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, tokenerr := _middlewares.ReadTokenId(c)
		if tokenerr != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Bad Request"))
		}
		idn := c.Param("id")
		id, _ := strconv.Atoi(idn)
		users, rows, err := uh.userUseCase.GetUser(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not exist"))
		}
		if idToken != id {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Bad Request"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get data", users))
	}
}

func (uh *UserHandler) DeleteUserHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, tokenerr := _middlewares.ReadTokenId(c)
		if tokenerr != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Bad Request"))
		}
		idn := c.Param("id")
		id, _ := strconv.Atoi(idn)
		_, rows, err := uh.userUseCase.GetUser(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Data not exist"))
		}
		if idToken != id {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Bad Request"))
		}
		_, err = uh.userUseCase.DeleteUser(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed delete data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("Succes delete data"))
	}
}

func (uh *UserHandler) CreateUserHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var user _entities.User
		c.Bind(&user)
		users, err := uh.userUseCase.CreateUser(user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed create data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("Succes create data", users))
	}
}

func (uh *UserHandler) UpdatedUserHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, tokenerr := _middlewares.ReadTokenId(c)
		if tokenerr != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Bad Request"))
		}
		idn := c.Param("id")
		id, _ := strconv.Atoi(idn)
		users, rows, err := uh.userUseCase.GetUser(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Data not exist"))
		}
		if idToken != id {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Bad Request"))
		}
		c.Bind(&users)
		users, err = uh.userUseCase.UpdatedUser(users, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed edit data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("Success edit data", users))
	}
}
