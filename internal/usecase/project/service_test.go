package project_test

import (
	"context"
	"errors"
	"golang-testing/domain"
	"golang-testing/internal/repository/mocks"
	"golang-testing/internal/usecase/project"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	mockProjectRepo := new(mocks.ProjectRepository)
	mockProject := domain.Project{
		Title:          "Hello",
		Description:    "Test description",
		Image:          "image.jpg",
		NumberOfTester: 2,
	}

	t.Run("success", func(t *testing.T) {
		tempMockProject := mockProject
		tempMockProject.ID = 0
		mockProjectRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.Project")).Return(nil).Once()

		u := project.NewProjectService(mockProjectRepo, time.Second*2)

		err := u.Create(context.TODO(), &tempMockProject)

		assert.NoError(t, err)
		assert.Equal(t, mockProject.Title, tempMockProject.Title)
		mockProjectRepo.AssertExpectations(t)
	})

}

func TestGetByID(t *testing.T) {
	mockProjectRepo := new(mocks.ProjectRepository)
	mockProject := domain.Project{
		Title:          "Hello",
		Description:    "Test description",
		Image:          "image.jpg",
		NumberOfTester: 2,
	}

	t.Run("success", func(t *testing.T) {
		mockProjectRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(&mockProject, nil).Once()

		u := project.NewProjectService(mockProjectRepo, time.Second*2)

		a, err := u.GetByID(context.TODO(), mockProject.ID)

		assert.NoError(t, err)
		assert.NotNil(t, a)

		mockProjectRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockProjectRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(nil, errors.New("Unexpected")).Once()

		u := project.NewProjectService(mockProjectRepo, time.Second*2)

		a, err := u.GetByID(context.TODO(), mockProject.ID)

		assert.Error(t, err)
		assert.Nil(t, a)

		mockProjectRepo.AssertExpectations(t)
	})

}
