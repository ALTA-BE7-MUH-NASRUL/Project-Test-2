package task

import (
	_entities "Tugas/Project-Test-2/entities"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTask(t *testing.T) {
	t.Run("TestGetTaskSuccess", func(t *testing.T) {
		taskUseCase := NewTaskUseCase(mockTaskRepository{})
		data, rows, err := taskUseCase.GetTask(1)
		assert.Nil(t, err)
		assert.Equal(t, "membuat fitur crud", data.List)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestGetTaskError", func(t *testing.T) {
		taskUseCase := NewTaskUseCase(mockTaskRepositoryError{})
		data, rows, err := taskUseCase.GetTask(1)
		assert.NotNil(t, err)
		assert.Equal(t, 0, rows)
		assert.Equal(t, _entities.Task{}, data)
	})
}

func TestDeleteTask(t *testing.T) {
	t.Run("TestDeleteTaskSuccess", func(t *testing.T) {
		taskUseCase := NewTaskUseCase(mockTaskRepository{})
		data, err := taskUseCase.DeleteTask(1)
		assert.Nil(t, err)
		assert.Equal(t, "deleted", data.Status)
	})

	t.Run("TestDeleteError", func(t *testing.T) {
		taskUseCase := NewTaskUseCase(mockTaskRepositoryError{})
		data, err := taskUseCase.DeleteTask(1)
		assert.NotNil(t, err)
		assert.Equal(t, "completed", data.Status)
	})
}

func TestCreateTask(t *testing.T) {
	t.Run("TestCreateTaskSuccess", func(t *testing.T) {
		taskUseCase := NewTaskUseCase(mockTaskRepository{})
		data, err := taskUseCase.CreateTask(1, 1, "membuat crud")
		assert.Nil(t, err)
		assert.Equal(t, "not completed", data.Status)
	})

	t.Run("TestCreateUserError", func(t *testing.T) {
		taskUseCase := NewTaskUseCase(mockTaskRepositoryError{})
		data, err := taskUseCase.CreateTask(0, 1, "")
		assert.NotNil(t, err)
		assert.Equal(t, _entities.Task{}, data)
	})
}

func TestUpdatedTask(t *testing.T) {
	t.Run("TestUpdatedTaskSuccess", func(t *testing.T) {
		taskUseCase := NewTaskUseCase(mockTaskRepository{})
		data, err := taskUseCase.UpdatedTask(_entities.Task{List: "membuat crud"}, 1)
		assert.Nil(t, err)
		assert.Equal(t, "not completed", data.Status)
	})

	t.Run("TestUpdatedTaskError", func(t *testing.T) {
		taskUseCase := NewTaskUseCase(mockTaskRepositoryError{})
		data, err := taskUseCase.UpdatedTask(_entities.Task{List: "membuat crud"}, 1)
		assert.NotNil(t, err)
		assert.Equal(t, _entities.Task{}, data)
	})
}
func TestReopenTask(t *testing.T) {
	t.Run("TestReopenTaskSuccess", func(t *testing.T) {
		taskUseCase := NewTaskUseCase(mockTaskRepository{})
		data, err := taskUseCase.ReopenTask(_entities.Task{}, 1)
		assert.Nil(t, err)
		assert.Equal(t, "not completed", data.Status)
	})

	t.Run("TestReopenUserError", func(t *testing.T) {
		taskUseCase := NewTaskUseCase(mockTaskRepositoryError{})
		data, err := taskUseCase.ReopenTask(_entities.Task{}, 1)
		assert.NotNil(t, err)
		assert.Equal(t, "completed", data.Status)
	})
}
func TestCompletedTask(t *testing.T) {
	t.Run("TestCompletedTaskSuccess", func(t *testing.T) {
		taskUseCase := NewTaskUseCase(mockTaskRepository{})
		data, err := taskUseCase.CompletedTask(_entities.Task{}, 1)
		assert.Nil(t, err)
		assert.Equal(t, "completed", data.Status)
	})

	t.Run("TestCompletedUserError", func(t *testing.T) {
		taskUseCase := NewTaskUseCase(mockTaskRepositoryError{})
		data, err := taskUseCase.CompletedTask(_entities.Task{}, 1)
		assert.NotNil(t, err)
		assert.Equal(t, "not completed", data.Status)
	})
}

// mock succes
type mockTaskRepository struct{}

func (m mockTaskRepository) GetTask(id int) (_entities.Task, int, error) {
	return _entities.Task{List: "membuat fitur crud"}, 1, nil
}
func (m mockTaskRepository) DeleteTask(id int) (_entities.Task, error) {
	return _entities.Task{Status: "deleted"}, nil
}
func (m mockTaskRepository) UpdatedTask(user _entities.Task) (_entities.Task, error) {
	return _entities.Task{Status: "not completed"}, nil
}
func (m mockTaskRepository) ReopenTask(user _entities.Task) (_entities.Task, error) {
	return _entities.Task{Status: "not completed"}, nil
}
func (m mockTaskRepository) CompletedTask(user _entities.Task) (_entities.Task, error) {
	return _entities.Task{Status: "completed"}, nil
}
func (m mockTaskRepository) CreateTask(UserId int, ProjectId int, List string) (_entities.Task, error) {
	return _entities.Task{Status: "not completed"}, nil
}

//  mock error

type mockTaskRepositoryError struct{}

func (m mockTaskRepositoryError) GetTask(id int) (_entities.Task, int, error) {
	return _entities.Task{}, 0, fmt.Errorf("error get data task")
}
func (m mockTaskRepositoryError) DeleteTask(id int) (_entities.Task, error) {
	return _entities.Task{Status: "completed"}, fmt.Errorf("error delete data task")
}
func (m mockTaskRepositoryError) CreateTask(UserId int, ProjectId int, List string) (_entities.Task, error) {
	return _entities.Task{}, fmt.Errorf("error create data task")
}
func (m mockTaskRepositoryError) ReopenTask(_entities.Task) (_entities.Task, error) {
	return _entities.Task{Status: "completed"}, fmt.Errorf("error updated data task")
}
func (m mockTaskRepositoryError) UpdatedTask(_entities.Task) (_entities.Task, error) {
	return _entities.Task{}, fmt.Errorf("error updated data task")
}
func (m mockTaskRepositoryError) CompletedTask(_entities.Task) (_entities.Task, error) {
	return _entities.Task{Status: "not completed"}, fmt.Errorf("error updated data task")
}
