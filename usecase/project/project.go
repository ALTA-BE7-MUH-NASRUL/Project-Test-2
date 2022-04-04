package project

import (
	_entities "Tugas/Project-Test-2/entities"
	_projcetRepository "Tugas/Project-Test-2/repo/project"
)

type ProjectUseCase struct {
	projcetRepository _projcetRepository.ProjectRepositoryInterface
}

func NewProjectUseCase(projectRepo _projcetRepository.ProjectRepositoryInterface) ProjectUseCaseInterface {
	return &ProjectUseCase{
		projcetRepository: projectRepo,
	}
}

func (tuc *ProjectUseCase) GetProject(id int) (_entities.Project, int, error) {
	Project, rows, err := tuc.projcetRepository.GetProject(id)
	return Project, rows, err
}

func (tuc *ProjectUseCase) DeleteProject(id int) (_entities.Project, error) {
	Project, err := tuc.projcetRepository.DeleteProject(id)
	return Project, err
}
func (tuc *ProjectUseCase) CreateProject(user _entities.Project) (_entities.Project, error) {
	Project, err := tuc.projcetRepository.CreateProject(user)
	return Project, err
}
func (tuc *ProjectUseCase) UpdatedProject(Project _entities.Project, id int) (_entities.Project, error) {
	Project.ID = uint(id)
	Project, err := tuc.projcetRepository.UpdatedProject(Project)
	return Project, err
}
func (tuc *ProjectUseCase) CompletedProject(Project _entities.Project, id int) (_entities.Project, error) {
	Project.ID = uint(id)
	Project, err := tuc.projcetRepository.CompletedProject(Project)
	return Project, err
}
