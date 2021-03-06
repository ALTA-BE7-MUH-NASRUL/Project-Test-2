package project

import (
	_entities "Tugas/Project-Test-2/entities"
)

type ProjectUseCaseInterface interface {
	GetProject(id int) (_entities.Project, int, error)
	DeleteProject(id int) (_entities.Project, error)
	CreateProject(UserId int, Title string) (_entities.Project, error)
	UpdatedProject(Project _entities.Project, id int) (_entities.Project, error)
	CompletedProject(Project _entities.Project, id int) (_entities.Project, error)
}
