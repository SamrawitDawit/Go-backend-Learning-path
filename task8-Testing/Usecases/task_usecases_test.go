package usecases

import (
	"context"
	"errors"
	domain "task8-Testing/Domain"
	"task8-Testing/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type taskUsecasesTestSuite struct {
	suite.Suite

	repository *mocks.TaskRepository
	usecase    domain.TaskUsecase
}

func (suite *taskUsecasesTestSuite) SetupTest() {
	repository := new(mocks.TaskRepository)

	usecase := NewTaskUsecase(repository, time.Second)

	suite.repository = repository
	suite.usecase = usecase
}

func (suite *taskUsecasesTestSuite) TestGetTasks_Filled_Positive() {
	tasks := []domain.Task{
		{
			ID:          primitive.NewObjectID(),
			Title:       "Task 1",
			Description: "Description 1",
			Status:      "Done",
		},
	}
	suite.repository.On("GetTasks", mock.Anything).Return(tasks, nil)
	result, err := suite.usecase.GetTasks(context.Background())
	suite.NoError(err)
	suite.Equal(tasks, result)
	suite.repository.AssertExpectations(suite.T())
}

func (suite *taskUsecasesTestSuite) TestGetTasks_EmptySlice_Posititve() {
	suite.repository.On("GetTasks", mock.Anything).Return([]domain.Task{}, nil)
	result, err := suite.usecase.GetTasks(context.Background())
	suite.NoError(err, "No error when get task")
	suite.Equal(len(result), 0, "tweets is an empty slice object")
	suite.repository.AssertExpectations(suite.T())

}

func (suite *taskUsecasesTestSuite) TestGetTask_Exists_Positive() {
	id, err := primitive.ObjectIDFromHex("66b74e21d7aa916a022b8479")
	suite.Assert().NoError(err)
	task := domain.Task{
		ID:          id,
		Title:       "Test Task",
		Description: "some description",
		DueDate:     time.Date(2023, time.August, 13, 15, 30, 0, 0, time.UTC),
		Status:      "Not started",
	}
	suite.repository.On("GetTask", mock.Anything, id).Return(task, nil)
	result, err := suite.usecase.GetTask(context.Background(), id)
	suite.NoError(err, "no error when get task with valid input")
	suite.Equal(task, result, "should be equal between result and task")
	suite.repository.AssertExpectations(suite.T())
}

func (suite *taskUsecasesTestSuite) TestGetTask_NotFound_Negative() {
	id, _ := primitive.ObjectIDFromHex("66b74e21d7aa916a022b8479")
	suite.repository.On("GetTask", mock.Anything, id).Return(domain.Task{}, errors.New("task not found"))
	_, err := suite.usecase.GetTask(context.Background(), id)
	suite.Error(err, "error because the task with the id is not found")
	suite.EqualError(err, "task not found", "error message should match")
	suite.repository.AssertExpectations(suite.T())
}

func (suite *taskUsecasesTestSuite) TestAddTask_Positive() {
	task := domain.Task{
		Title:       "Test Task",
		Description: "some description",
		DueDate:     time.Now(),
		Status:      "Not started",
	}
	suite.repository.On("AddTask", mock.Anything, &task).Return(nil)
	err := suite.usecase.AddTask(context.Background(), &task)
	suite.NoError(err, "No error when add task with valid input")
	suite.repository.AssertExpectations(suite.T())
}

func (suite *taskUsecasesTestSuite) TestRemoveTask_ID_Exists_Positive() {
	id, err := primitive.ObjectIDFromHex("66b74e21d7aa916a022b8479")
	suite.Assert().NoError(err)
	suite.repository.On("RemoveTask", mock.Anything, id).Return(nil)
	err = suite.usecase.RemoveTask(context.Background(), id)
	suite.NoError(err, "no error when remove task with valid input")
	suite.repository.AssertExpectations(suite.T())
}
func (suite *taskUsecasesTestSuite) TestRemoveTask_InvalidID_Negative() {
	id, _ := primitive.ObjectIDFromHex("66bab91f6844ea2b60843fea")
	suite.repository.On("RemoveTask", mock.Anything, id).Return(errors.New("mongo: no documents in result"))
	err := suite.usecase.RemoveTask(context.Background(), id)
	suite.Error(err, "error because the task with the id is not found")
	suite.repository.AssertExpectations(suite.T())
}
func (suite *taskUsecasesTestSuite) TestUpdateTask_Positive() {
	id, err := primitive.ObjectIDFromHex("66b74e21d7aa916a022b8479")
	suite.Assert().NoError(err)
	updatedTask := domain.Task{
		Title:       "updated task",
		Description: "updated description",
		DueDate:     time.Date(2023, time.August, 13, 15, 30, 0, 0, time.UTC),
		Status:      "Not started",
	}
	suite.repository.On("UpdateTask", mock.Anything, id, &updatedTask).Return(nil)
	err = suite.usecase.UpdateTask(context.Background(), id, &updatedTask)
	suite.NoError(err, "no error because the task with the id is found")
	suite.repository.AssertExpectations(suite.T())
}
func (suite *taskUsecasesTestSuite) TestUpdateTask_InvalidID_Negative() {
	id, err := primitive.ObjectIDFromHex("66b74e21d7aa916a022b8479")
	suite.Assert().NoError(err)
	updatedTask := domain.Task{
		Title:       "updated task",
		Description: "updated description",
		DueDate:     time.Now(),
		Status:      "Not started",
	}
	suite.repository.On("UpdateTask", mock.Anything, id, &updatedTask).Return(errors.New("task not found"))
	err = suite.usecase.UpdateTask(context.Background(), id, &updatedTask)
	suite.Error(err, "error when update task with invalid input")
	suite.Error(err, "error because the task with the id is not found")
	suite.repository.AssertExpectations(suite.T())
}
func (suite *taskUsecasesTestSuite) TestUpdateTask_EmptyFields_Positive() {
	id, err := primitive.ObjectIDFromHex("66b74e21d7aa916a022b8479")
	suite.Assert().NoError(err)
	updatedTask := domain.Task{}
	suite.repository.On("UpdateTask", mock.Anything, id, &updatedTask).Return(nil)
	err = suite.usecase.UpdateTask(context.Background(), id, &updatedTask)
	suite.NoError(err, "no error when update task with empty fields")
	suite.repository.AssertExpectations(suite.T())
}

func TestTaskUseCase(t *testing.T) {
	suite.Run(t, new(taskUsecasesTestSuite))
}
