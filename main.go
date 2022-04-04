package main

import (
	"Tugas/Project-Test-2/config"
	_authHandler "Tugas/Project-Test-2/delivery/handler/auth"
	_projectHandler "Tugas/Project-Test-2/delivery/handler/project"
	_taskHandler "Tugas/Project-Test-2/delivery/handler/task"
	_userHandler "Tugas/Project-Test-2/delivery/handler/user"
	_middlewares "Tugas/Project-Test-2/delivery/middleware"
	_routes "Tugas/Project-Test-2/delivery/route"
	_authRepository "Tugas/Project-Test-2/repo/auth"
	_projectRepository "Tugas/Project-Test-2/repo/project"
	_taskRepository "Tugas/Project-Test-2/repo/task"
	_userRepository "Tugas/Project-Test-2/repo/user"
	_authUseCase "Tugas/Project-Test-2/usecase/auth"
	_projectUseCase "Tugas/Project-Test-2/usecase/project"
	_taskUseCase "Tugas/Project-Test-2/usecase/task"
	_userUseCase "Tugas/Project-Test-2/usecase/user"
	_utils "Tugas/Project-Test-2/utils"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config := config.GetConfig()
	db := _utils.InitDB(config)

	userRepo := _userRepository.NewUserRepository(db)
	userUseCase := _userUseCase.NewUserUseCase(userRepo)
	userHandler := _userHandler.NewUserHandler(userUseCase)
	taskRepo := _taskRepository.NewTaskRepository(db)
	taskUseCase := _taskUseCase.NewTaskUseCase(taskRepo)
	taskHandler := _taskHandler.NewTaskHandler(taskUseCase)
	authRepo := _authRepository.NewAuthRepository(db)
	authUseCase := _authUseCase.NewAuthUseCase(authRepo)
	authHandler := _authHandler.NewAuthHandler(authUseCase)
	projectRepo := _projectRepository.NewProjectRepository(db)
	projectUseCase := _projectUseCase.NewProjectUseCase(projectRepo)
	projectHandler := _projectHandler.NewProjectHandler(projectUseCase)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(_middlewares.CustomLogger())
	_routes.RegisterPath(e, userHandler, taskHandler, projectHandler)
	_routes.RegisterAuthPath(e, authHandler)
	log.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))

}
