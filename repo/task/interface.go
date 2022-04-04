package task

import (
	_entities "Tugas/Project-Test-2/entities"
)

type TaskRepositoryInterface interface {
	GetTask(id int) (_entities.Task, int, error)
	DeleteTask(id int) (_entities.Task, error)
	CreateTask(UserId int, ProjectId int, List string) (_entities.Task, error)
	UpdatedTask(task _entities.Task) (_entities.Task, error)
	ReopenTask(task _entities.Task) (_entities.Task, error)
	CompletedTask(task _entities.Task) (_entities.Task, error)
}
