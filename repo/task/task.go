package task

import (
	_entities "Tugas/Project-Test-2/entities"

	"gorm.io/gorm"
)

type TaskRepository struct {
	database *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{
		database: db,
	}
}

func (tr *TaskRepository) CreateTask(UserId int, ProjectId int, List string) (_entities.Task, error) {
	var task _entities.Task
	var project _entities.Project
	err := tr.database.Find(&project, ProjectId)
	if err.Error != nil {
		return task, err.Error
	}
	if err.RowsAffected == 0 {
		return task, err.Error
	}
	task.UserID = uint(UserId)
	task.ProjectID = uint(ProjectId)
	task.List = List
	task.Status = "not completed"
	tx := tr.database.Save(&task)
	if tx.Error != nil {
		return task, tx.Error
	}
	if tx.RowsAffected == 0 {
		return task, tx.Error
	}
	return task, nil
}

func (tr *TaskRepository) GetTask(id int) (_entities.Task, int, error) {
	var task _entities.Task
	tx := tr.database.Find(&task, id)
	if tx.Error != nil {
		return task, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return task, 0, tx.Error
	}
	return task, int(tx.RowsAffected), nil
}

func (tr *TaskRepository) UpdatedTask(task _entities.Task) (_entities.Task, error) {
	tx := tr.database.Save(&task)
	if tx.Error != nil {
		return task, tx.Error
	}
	if tx.RowsAffected == 0 {
		return task, tx.Error
	}
	return task, nil
}

func (tr *TaskRepository) DeleteTask(id int) (_entities.Task, error) {
	var task _entities.Task
	tx := tr.database.Delete(&task, id)
	if tx.Error != nil {
		return task, tx.Error
	}
	if tx.RowsAffected == 0 {
		return task, tx.Error

	}
	return task, nil
}

func (tr *TaskRepository) ReopenTask(task _entities.Task) (_entities.Task, error) {
	task.Status = "not completed"
	tx := tr.database.Save(&task)
	if tx.Error != nil {
		return task, tx.Error
	}
	if tx.RowsAffected == 0 {
		return task, tx.Error
	}
	return task, nil
}

func (tr *TaskRepository) CompletedTask(task _entities.Task) (_entities.Task, error) {
	task.Status = "completed"
	tx := tr.database.Save(&task)
	if tx.Error != nil {
		return task, tx.Error
	}
	if tx.RowsAffected == 0 {
		return task, tx.Error
	}
	return task, nil
}
