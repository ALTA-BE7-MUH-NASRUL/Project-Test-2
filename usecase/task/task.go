package task

import (
	_entities "Tugas/Project-Test-2/entities"
	_taskRepository "Tugas/Project-Test-2/repo/task"
)

type TaskUseCase struct {
	taskRepository _taskRepository.TaskRepositoryInterface
}

func NewTaskUseCase(taskRepo _taskRepository.TaskRepositoryInterface) TaskUseCaseInterface {
	return &TaskUseCase{
		taskRepository: taskRepo,
	}
}

func (tuc *TaskUseCase) GetTask(id int) (_entities.Task, int, error) {
	task, rows, err := tuc.taskRepository.GetTask(id)
	return task, rows, err
}

func (tuc *TaskUseCase) DeleteTask(id int) (_entities.Task, error) {
	task, err := tuc.taskRepository.DeleteTask(id)
	return task, err
}
func (tuc *TaskUseCase) CreateTask(UserId int, ProjectId int, List string) (_entities.Task, error) {
	task, err := tuc.taskRepository.CreateTask(UserId, ProjectId, List)
	return task, err
}
func (tuc *TaskUseCase) UpdatedTask(task _entities.Task, id int) (_entities.Task, error) {
	task.ID = uint(id)
	task, err := tuc.taskRepository.UpdatedTask(task)
	return task, err
}
func (tuc *TaskUseCase) ReopenTask(task _entities.Task, id int) (_entities.Task, error) {
	task.ID = uint(id)
	task, err := tuc.taskRepository.ReopenTask(task)
	return task, err
}
func (tuc *TaskUseCase) CompletedTask(task _entities.Task, id int) (_entities.Task, error) {
	task.ID = uint(id)
	task, err := tuc.taskRepository.CompletedTask(task)
	return task, err
}
