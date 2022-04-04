package project

import (
	_entities "Tugas/Project-Test-2/entities"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProject(t *testing.T) {
	t.Run("TestGetProjectSuccess", func(t *testing.T) {
		projectUseCase := NewProjectUseCase(mockProjectRepository{})
		data, rows, err := projectUseCase.GetProject(1)
		assert.Nil(t, err)
		assert.Equal(t, "membuat fitur crud", data.Title)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestGetProjectError", func(t *testing.T) {
		projectUseCase := NewProjectUseCase(mockProjectRepositoryError{})
		data, rows, err := projectUseCase.GetProject(1)
		assert.NotNil(t, err)
		assert.Equal(t, 0, rows)
		assert.Equal(t, _entities.Project{}, data)
	})
}

func TestDeleteProject(t *testing.T) {
	t.Run("TestDeleteProjectSuccess", func(t *testing.T) {
		projectUseCase := NewProjectUseCase(mockProjectRepository{})
		data, err := projectUseCase.DeleteProject(1)
		assert.Nil(t, err)
		assert.Equal(t, "deleted", data.Status)
	})

	t.Run("TestDeleteError", func(t *testing.T) {
		projectUseCase := NewProjectUseCase(mockProjectRepositoryError{})
		data, err := projectUseCase.DeleteProject(1)
		assert.NotNil(t, err)
		assert.Equal(t, "completed", data.Status)
	})
}

func TestCreateProject(t *testing.T) {
	t.Run("TestCreateProjectSuccess", func(t *testing.T) {
		projectUseCase := NewProjectUseCase(mockProjectRepository{})
		data, err := projectUseCase.CreateProject(1, "project")
		assert.Nil(t, err)
		assert.Equal(t, "not completed", data.Status)
	})

	t.Run("TestCreateUserError", func(t *testing.T) {
		projectUseCase := NewProjectUseCase(mockProjectRepositoryError{})
		data, err := projectUseCase.CreateProject(0, "project")
		assert.NotNil(t, err)
		assert.Equal(t, _entities.Project{}, data)
	})
}

func TestUpdatedProject(t *testing.T) {
	t.Run("TestUpdatedProjectSuccess", func(t *testing.T) {
		projectUseCase := NewProjectUseCase(mockProjectRepository{})
		data, err := projectUseCase.UpdatedProject(_entities.Project{Title: "project"}, 1)
		assert.Nil(t, err)
		assert.Equal(t, "not completed", data.Status)
	})

	t.Run("TestUpdatedProjectError", func(t *testing.T) {
		projectUseCase := NewProjectUseCase(mockProjectRepositoryError{})
		data, err := projectUseCase.UpdatedProject(_entities.Project{Title: "project"}, 1)
		assert.NotNil(t, err)
		assert.Equal(t, _entities.Project{}, data)
	})
}
func TestCompletedProject(t *testing.T) {
	t.Run("TestCompletedProjectSuccess", func(t *testing.T) {
		projectUseCase := NewProjectUseCase(mockProjectRepository{})
		data, err := projectUseCase.CompletedProject(_entities.Project{}, 1)
		assert.Nil(t, err)
		assert.Equal(t, "completed", data.Status)
	})

	t.Run("TestCompletedUserError", func(t *testing.T) {
		projectUseCase := NewProjectUseCase(mockProjectRepositoryError{})
		data, err := projectUseCase.CompletedProject(_entities.Project{}, 1)
		assert.NotNil(t, err)
		assert.Equal(t, "not completed", data.Status)
	})
}

// mock succes
type mockProjectRepository struct{}

func (m mockProjectRepository) GetProject(id int) (_entities.Project, int, error) {
	return _entities.Project{Title: "membuat fitur crud"}, 1, nil
}
func (m mockProjectRepository) DeleteProject(id int) (_entities.Project, error) {
	return _entities.Project{Status: "deleted"}, nil
}
func (m mockProjectRepository) UpdatedProject(user _entities.Project) (_entities.Project, error) {
	return _entities.Project{Status: "not completed"}, nil
}
func (m mockProjectRepository) CompletedProject(user _entities.Project) (_entities.Project, error) {
	return _entities.Project{Status: "completed"}, nil
}
func (m mockProjectRepository) CreateProject(UserId int, Title string) (_entities.Project, error) {
	return _entities.Project{Status: "not completed"}, nil
}

//  mock error

type mockProjectRepositoryError struct{}

func (m mockProjectRepositoryError) GetProject(id int) (_entities.Project, int, error) {
	return _entities.Project{}, 0, fmt.Errorf("error get data project")
}
func (m mockProjectRepositoryError) DeleteProject(id int) (_entities.Project, error) {
	return _entities.Project{Status: "completed"}, fmt.Errorf("error delete data project")
}
func (m mockProjectRepositoryError) CreateProject(UserId int, Title string) (_entities.Project, error) {
	return _entities.Project{}, fmt.Errorf("error create data project")
}
func (m mockProjectRepositoryError) UpdatedProject(_entities.Project) (_entities.Project, error) {
	return _entities.Project{}, fmt.Errorf("error updated data project")
}
func (m mockProjectRepositoryError) CompletedProject(_entities.Project) (_entities.Project, error) {
	return _entities.Project{Status: "not completed"}, fmt.Errorf("error updated data project")
}
