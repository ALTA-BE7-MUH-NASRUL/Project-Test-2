package route

import (
	_authHandler "Tugas/Project-Test-2/delivery/handler/auth"
	_projectHandler "Tugas/Project-Test-2/delivery/handler/project"
	_taskHandler "Tugas/Project-Test-2/delivery/handler/task"
	_userHandler "Tugas/Project-Test-2/delivery/handler/user"
	_middlewares "Tugas/Project-Test-2/delivery/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, uh *_userHandler.UserHandler, th *_taskHandler.TaskHandler, ph *_projectHandler.ProjectHandler) {
	// user
	e.POST("/users", uh.CreateUserHandler())
	e.GET("/users/:id", uh.GetUserHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/users/:id", uh.DeleteUserHandler(), _middlewares.JWTMiddleware())
	e.PUT("/users/:id", uh.UpdatedUserHandler(), _middlewares.JWTMiddleware())

	// Task
	e.POST("/tasks", th.CreateTaskHandler(), _middlewares.JWTMiddleware())
	e.GET("/tasks/:id", th.GetTaskHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/tasks/:id", th.DeleteTaskHandler(), _middlewares.JWTMiddleware())
	e.PUT("/tasks/:id", th.UpdatedTaskHandler(), _middlewares.JWTMiddleware())
	e.PUT("/tasks/completed/:id", th.CompletedTaskHandler(), _middlewares.JWTMiddleware())
	e.PUT("/tasks/reopen/:id", th.ReopenTaskHandler(), _middlewares.JWTMiddleware())

	// Project
	e.POST("/projects", ph.CreateProjectHandler(), _middlewares.JWTMiddleware())
	e.GET("/projects/:id", ph.GetProjectHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/projects/:id", ph.DeleteProjectHandler(), _middlewares.JWTMiddleware())
	e.PUT("/projects/:id", ph.UpdatedProjectHandler(), _middlewares.JWTMiddleware())
	e.PUT("/projects/completed/:id", ph.CompletedProjectHandler(), _middlewares.JWTMiddleware())
}

func RegisterAuthPath(e *echo.Echo, ah *_authHandler.AuthHandler) {
	e.POST("/auth", ah.LoginHandler())
}
