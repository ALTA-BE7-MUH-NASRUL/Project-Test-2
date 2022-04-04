package project

import (
	_entities "Tugas/Project-Test-2/entities"

	"gorm.io/gorm"
)

type ProjectRepository struct {
	database *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{
		database: db,
	}
}

func (tr *ProjectRepository) CreateProject(_entities.Project) (_entities.Project, error) {
	var Project _entities.Project
	Project.Status = "not completed"
	tx := tr.database.Save(&Project)
	if tx.Error != nil {
		return Project, tx.Error
	}
	if tx.RowsAffected == 0 {
		return Project, tx.Error
	}
	return Project, nil
}

func (tr *ProjectRepository) GetProject(id int) (_entities.Project, int, error) {
	var Project _entities.Project
	tx := tr.database.Find(&Project, id)
	if tx.Error != nil {
		return Project, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return Project, 0, tx.Error
	}
	return Project, int(tx.RowsAffected), nil
}

func (tr *ProjectRepository) UpdatedProject(Project _entities.Project) (_entities.Project, error) {
	tx := tr.database.Save(&Project)
	if tx.Error != nil {
		return Project, tx.Error
	}
	if tx.RowsAffected == 0 {
		return Project, tx.Error
	}
	return Project, nil
}

func (tr *ProjectRepository) DeleteProject(id int) (_entities.Project, error) {
	var Project _entities.Project
	tx := tr.database.Delete(&Project, id)
	if tx.Error != nil {
		return Project, tx.Error
	}
	if tx.RowsAffected == 0 {
		return Project, tx.Error

	}
	return Project, nil
}

func (tr *ProjectRepository) CompletedProject(Project _entities.Project) (_entities.Project, error) {
	Project.Status = "completed"
	tx := tr.database.Save(&Project)
	if tx.Error != nil {
		return Project, tx.Error
	}
	if tx.RowsAffected == 0 {
		return Project, tx.Error
	}
	return Project, nil
}
