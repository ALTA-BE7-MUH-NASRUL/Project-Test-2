package task

import (
	_entities "Tugas/Project-Test-2/entities"
)

type TaskUseCaseInterface interface {
	GetTask(id int) (_entities.Task, int, error)
	DeleteTask(id int) (_entities.Task, error)
	CreateTask(UserId int, ProjectId int, list string) (_entities.Task, error)
	UpdatedTask(task _entities.Task, id int) (_entities.Task, error)
	ReopenTask(task _entities.Task, id int) (_entities.Task, error)
	CompletedTask(task _entities.Task, id int) (_entities.Task, error)
}
