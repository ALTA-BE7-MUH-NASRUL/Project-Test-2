package project

import (
	_entities "Tugas/Project-Test-2/entities"
)

type ProjectRepositoryInterface interface {
	GetProject(id int) (_entities.Project, int, error)
	DeleteProject(id int) (_entities.Project, error)
	CreateProject(Project _entities.Project) (_entities.Project, error)
	UpdatedProject(Project _entities.Project) (_entities.Project, error)
	CompletedProject(Project _entities.Project) (_entities.Project, error)
}
