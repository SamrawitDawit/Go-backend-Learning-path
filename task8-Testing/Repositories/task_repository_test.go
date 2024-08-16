package repositories

import (
	"context"
	"log"
	domain "task8-Testing/Domain"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type taskRepositorySuite struct {
	suite.Suite
	repository domain.TaskRepository
}

func (suite *taskRepositorySuite) SetupSuite() {
	clientOption := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("test_db")
	repo := NewTaskRepository(*db, "tasks")

	suite.repository = repo
}

func (suite *taskRepositorySuite) TearDownTest() {
	suite.repository.(*taskRepository).database.Collection("tasks").Drop(context.Background())
}

func (suite *taskRepositorySuite) TestGetTasks_EmptySlice_Posititve() {
	tasks, err := suite.repository.GetTasks(context.Background())
	suite.NoError(err, "No error when get task")
	suite.Equal(len(tasks), 0, "length of tasks should be 0, since it is empty slice")
	suite.Equal(tasks, []domain.Task(nil), "tasks is an empty slice")
}

func (suite *taskRepositorySuite) TestGetTasks_FilledRecords_Positive() {
	task := domain.Task{
		Title:       "Test Task",
		Description: "some description",
		DueDate:     time.Now(),
		Status:      "Not started",
	}

	err := suite.repository.AddTask(context.Background(), &task)
	suite.NoError(err, "no error when add task with valid input")
	err = suite.repository.AddTask(context.Background(), &task)
	suite.NoError(err, "no error when add task with valid input")
	err = suite.repository.AddTask(context.Background(), &task)
	suite.NoError(err, "no error when add task with valid input")

	tasks, err := suite.repository.GetTasks(context.Background())
	suite.NoError(err, "no error when get tasks when the table is filled")
	suite.Equal(len(tasks), 3, "insert 3 reords before get all data, so it should contain 3 tasks")
}
func (suite *taskRepositorySuite) TestGetTask_Exists_Positive() {
	id, err := primitive.ObjectIDFromHex("66b74e21d7aa916a022b8479")
	suite.Assert().NoError(err)
	task := domain.Task{
		ID:          id,
		Title:       "Test Task",
		Description: "some description",
		DueDate:     time.Date(2023, time.August, 13, 15, 30, 0, 0, time.UTC),
		Status:      "Not started",
	}
	err = suite.repository.AddTask(context.Background(), &task)
	suite.NoError(err, "no error when add task with valid input")
	result, err := suite.repository.GetTask(context.Background(), id)
	suite.NoError(err, "no error because the task with the id is found")
	suite.Equal(task.Title, result.Title, "should be equal between result and task")
	suite.Equal(task.Description, result.Description, "should be equal between result and task")
	suite.Equal(task.DueDate, result.DueDate, "should be equal between result and task")
	suite.Equal(task.Status, result.Status, "should be equal between result and task")
}

func (suite *taskRepositorySuite) TestTask_NotFound_Negative() {
	id, err := primitive.ObjectIDFromHex("66b74e21d7aa916a022b8479")
	suite.Assert().NoError(err)

	_, err = suite.repository.GetTask(context.Background(), id)
	suite.Error(err, "error because the task with the id is not found")
	suite.Equal(err.Error(), "mongo: no documents in result", "error message should be 'mongo: no documents in result'")
}

func (suite *taskRepositorySuite) TestAddTask_Positive() {
	task := domain.Task{
		Title:       "Test Task",
		Description: "some description",
		DueDate:     time.Now(),
		Status:      "Not started",
	}
	err := suite.repository.AddTask(context.Background(), &task)
	suite.NoError(err, "No error when add task with valid input")
}
func (suite *taskRepositorySuite) TestAddTask_NilPointer_Negative() {
	err := suite.repository.AddTask(context.Background(), nil)
	suite.Error(err, "Add task with nil input returns error")
}

func (suite *taskRepositorySuite) TestRemoveTask_ID_Exists_Positive() {
	id, err := primitive.ObjectIDFromHex("66b74e21d7aa916a022b8479")
	suite.Assert().NoError(err)
	task := domain.Task{
		ID:          id,
		Title:       "Test Task",
		Description: "some description",
		DueDate:     time.Now(),
		Status:      "Not started",
	}
	err = suite.repository.AddTask(context.Background(), &task)
	suite.NoError(err, "no error when add task with valid input")
	err = suite.repository.RemoveTask(context.Background(), id)
	suite.NoError(err, "no error when remove task with valid input")
	_, err = suite.repository.GetTask(context.Background(), id)
	suite.Error(err, "error because the task with the id is not found(since it is removed)")
}
func (suite *taskRepositorySuite) TestRemoveTask_InvalidID_Negative() {
	id, _ := primitive.ObjectIDFromHex("66bab91f6844ea2b60843fea")
	err := suite.repository.RemoveTask(context.Background(), id)
	suite.Error(err, "error because the task with the id is not found")
}
func (suite *taskRepositorySuite) TestUpdateTask_Positive() {
	id, err := primitive.ObjectIDFromHex("66b74e21d7aa916a022b8479")
	suite.Assert().NoError(err)
	task := domain.Task{
		ID:          id,
		Title:       "Test Task",
		Description: "some description",
		DueDate:     time.Now(),
		Status:      "Not started",
	}
	err = suite.repository.AddTask(context.Background(), &task)
	suite.NoError(err, "no error when add task with valid input")
	updatedTask := domain.Task{
		Title:       "updated task",
		Description: "updated description",
		DueDate:     time.Date(2023, time.August, 13, 15, 30, 0, 0, time.UTC),
		Status:      "Not started",
	}
	err = suite.repository.UpdateTask(context.Background(), id, &updatedTask)
	suite.NoError(err, "no error when update task with valid input")
	result, err := suite.repository.GetTask(context.Background(), id)
	suite.NoError(err, "no error because the task with the id is found")
	suite.Equal(updatedTask.Title, result.Title, "should be equal between result and task")
	suite.Equal(updatedTask.Description, result.Description, "should be equal between result and task")
	suite.Equal(updatedTask.DueDate, result.DueDate, "should be equal between result and task")
	suite.Equal(updatedTask.Status, result.Status, "should be equal between result and task")
}
func (suite *taskRepositorySuite) TestUpdateTask_InvalidID_Negative() {
	id, err := primitive.ObjectIDFromHex("66b74e21d7aa916a022b8479")
	suite.Assert().NoError(err)
	updatedTask := domain.Task{
		Title:       "updated task",
		Description: "updated description",
		DueDate:     time.Now(),
		Status:      "Not started",
	}
	err = suite.repository.UpdateTask(context.Background(), id, &updatedTask)
	suite.Error(err, "error when update task with invalid input")
	suite.Error(err, "error because the task with the id is not found")
}
func (suite *taskRepositorySuite) TestUpdateTask_NilPointer_Negative() {
	id, err := primitive.ObjectIDFromHex("66b74e21d7aa916a022b8479")
	suite.Assert().NoError(err)
	task := domain.Task{
		ID:          id,
		Title:       "Test Task",
		Description: "some description",
		DueDate:     time.Now(),
		Status:      "Not started",
	}
	err = suite.repository.AddTask(context.Background(), &task)
	suite.NoError(err, "no error when add task with valid input")
	err = suite.repository.UpdateTask(context.Background(), id, nil)
	suite.Error(err, "error when update task with nil input")

}
func (suite *taskRepositorySuite) TestUpdateTask_EmptyFields_Positive() {
	id, err := primitive.ObjectIDFromHex("66b74e21d7aa916a022b8479")
	suite.Assert().NoError(err)
	task := domain.Task{
		ID:          id,
		Title:       "Test Task",
		Description: "some description",
		DueDate:     time.Now(),
		Status:      "Not started",
	}
	err = suite.repository.AddTask(context.Background(), &task)
	suite.NoError(err, "no error when add task with valid input")
	var updatedTask domain.Task
	err = suite.repository.UpdateTask(context.Background(), id, &updatedTask)
	suite.NoError(err, "no error when update task with empty fields")
}

func TestTaskRepository(t *testing.T) {
	suite.Run(t, new(taskRepositorySuite))
}