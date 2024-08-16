package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	domain "task8-Testing/Domain"
	"task8-Testing/mocks"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ControllerSuite struct {
	suite.Suite
	taskusecase *mocks.TaskUsecase
	userusecase *mocks.UserUsecase
	controller  Controller
	router      *gin.Engine
}

func (suite *ControllerSuite) SetupSuite() {
	suite.taskusecase = new(mocks.TaskUsecase)
	suite.userusecase = new(mocks.UserUsecase)
	suite.controller = Controller{
		TaskUsecase: suite.taskusecase,
		UserUsecase: suite.userusecase,
	}

	suite.router = gin.Default()

	suite.router.POST("/register", suite.controller.Register)
	suite.router.POST("/login", suite.controller.Login)
	suite.router.POST("/users/:_id", suite.controller.Promote)
	suite.router.POST("/tasks", suite.controller.AddTask)
	suite.router.GET("/tasks", suite.controller.GetTasks)
	suite.router.GET("/tasks/:_id", suite.controller.GetTask)
	suite.router.PUT("/tasks/:_id", suite.controller.UpdateTask)
	suite.router.DELETE("/tasks/:_id", suite.controller.RemoveTask)
}

func (suite *ControllerSuite) TearDownTest() {
	suite.taskusecase.AssertExpectations(suite.T())
	suite.userusecase.AssertExpectations(suite.T())
}

func (suite *ControllerSuite) TestRegister_Positive() {
	user := &domain.User{
		Username: "samrawit",
		Password: "samrawit",
	}
	suite.userusecase.On("Register", mock.Anything, user).Return(nil).Once()

	jsonInput, _ := json.Marshal(user)
	request, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(jsonInput))
	request.Header.Set("Content-Type", "application/json")

	responseWriter := httptest.NewRecorder()
	suite.router.ServeHTTP(responseWriter, request)

	suite.Assert().Equal(responseWriter.Code, http.StatusCreated)
	suite.Assert().Contains(responseWriter.Body.String(), "User registered successfully")
}
func (suite *ControllerSuite) Test_Register_InvaldInput_Negative() {
	user := &domain.User{
		Username: "",
		Password: "",
	}

	jsonInput, _ := json.Marshal(user)
	request, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(jsonInput))
	request.Header.Set("Content-Type", "application/json")

	responseWriter := httptest.NewRecorder()
	suite.router.ServeHTTP(responseWriter, request)

	suite.Assert().Equal(responseWriter.Code, http.StatusBadRequest)
	suite.Assert().Contains(responseWriter.Body.String(), "Username and Password are required fields")

}
func (suite *ControllerSuite) TestRegister_UserName_Not_Available_Negative() {
	user := &domain.User{
		Username: "samrawit",
		Password: "samrawit",
	}
	suite.userusecase.On("Register", mock.Anything, user).Return(errors.New("username already exists")).Once()
	jsonInput, _ := json.Marshal(user)
	request, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(jsonInput))
	request.Header.Set("Content-Type", "application/json")

	responseWriter := httptest.NewRecorder()
	suite.router.ServeHTTP(responseWriter, request)

	suite.Assert().Equal(responseWriter.Code, http.StatusInternalServerError)
	suite.Assert().Contains(responseWriter.Body.String(), "username already exists")
}

func (suite *ControllerSuite) TestLogin_Positive() {
	user := &domain.User{
		Username: "samrawit",
		Password: "samrawit",
	}
	token := "some token"
	suite.userusecase.On("Login", mock.Anything, user).Return(token, nil).Once()
	jsonInput, _ := json.Marshal(user)
	request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonInput))
	request.Header.Set("Content-Type", "application/json")

	responseWriter := httptest.NewRecorder()
	suite.router.ServeHTTP(responseWriter, request)

	suite.Assert().Equal(responseWriter.Code, http.StatusOK)
	suite.Assert().Contains(responseWriter.Body.String(), "User logged in successfully")
	suite.Assert().Contains(responseWriter.Body.String(), token)
}

func (suite *ControllerSuite) TestLogin_InvaldInput_Negative() {
	user := &domain.User{
		Username: "",
		Password: "",
	}
	jsonInput, _ := json.Marshal(user)

	request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonInput))
	request.Header.Set("Content_Type", "application/json")

	responseWriter := httptest.NewRecorder()
	suite.router.ServeHTTP(responseWriter, request)

	suite.Assert().Equal(responseWriter.Code, http.StatusBadRequest)
	suite.Assert().Contains(responseWriter.Body.String(), "Username and Password are required fields")
}

func (suite *ControllerSuite) TestLogin_NotMatching_Credentials_Negative() {
	user := &domain.User{
		Username: "samrawit",
		Password: "samrawit",
	}

	suite.userusecase.On("Login", mock.Anything, user).Return("", errors.New("invalid credentials")).Once()
	jsonInput, _ := json.Marshal(user)
	request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonInput))
	request.Header.Set("Content-Type", "application/json")

	responseWriter := httptest.NewRecorder()
	suite.router.ServeHTTP(responseWriter, request)

	suite.Assert().Equal(responseWriter.Code, http.StatusUnauthorized)
	suite.Assert().Contains(responseWriter.Body.String(), "invalid credentials")
}

func (suite *ControllerSuite) TestPromote_Positive() {
	id := primitive.NewObjectID()

	suite.userusecase.On("Promote", mock.Anything, id).Return(nil).Once()

	req, _ := http.NewRequest(http.MethodPost, "/users/"+id.Hex(), nil)
	recorder := httptest.NewRecorder()
	suite.router.ServeHTTP(recorder, req)

	suite.Assert().Equal(http.StatusOK, recorder.Code)
	suite.Assert().Contains(recorder.Body.String(), "User promoted successfully")
}

func (suite *ControllerSuite) TestPromote_InvalidID_Negative() {
	id := primitive.NewObjectID()

	suite.userusecase.On("Promote", mock.Anything, id).Return(errors.New("Invalid ID")).Once()

	req, _ := http.NewRequest(http.MethodPost, "/users/"+id.Hex(), nil)
	recorder := httptest.NewRecorder()
	suite.router.ServeHTTP(recorder, req)

	suite.Assert().Equal(http.StatusUnauthorized, recorder.Code)
	suite.Assert().Contains(recorder.Body.String(), "Invalid ID")
}

func (suite *ControllerSuite) TestGetTasks_Positive() {
	tasks := []domain.Task{
		{
			ID:          primitive.NewObjectID(),
			Title:       "Task 1",
			Description: "Description 1",
			Status:      "In Progress",
		},
	}

	suite.taskusecase.On("GetTasks", mock.Anything).Return(tasks, nil).Once()

	req, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
	recorder := httptest.NewRecorder()
	suite.router.ServeHTTP(recorder, req)

	suite.Assert().Equal(http.StatusOK, recorder.Code)
	suite.Assert().Contains(recorder.Body.String(), "Task 1")
}

func (suite *ControllerSuite) TestGetTasks_Negative() {
	suite.taskusecase.On("GetTasks", mock.Anything).Return(domain.Task{}, errors.New("something went wrong while retrieving the tasks")).Once()

	req, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
	recorder := httptest.NewRecorder()
	suite.router.ServeHTTP(recorder, req)

	suite.Assert().Equal(http.StatusInternalServerError, recorder.Code)
	suite.Assert().Equal(len(recorder.Body.String()), 0)
}

func (suite *ControllerSuite) TestGetTask_Positive() {
	id, _ := primitive.ObjectIDFromHex("66b74e21d7aa916a022b8479")

	task := domain.Task{
		ID:          id,
		Title:       "Task 1",
		Description: "Description 1",
		Status:      "In Progress",
	}

	suite.taskusecase.On("GetTask", mock.Anything, id).Return(task, nil).Once()

	req, _ := http.NewRequest(http.MethodGet, "/tasks/"+id.Hex(), nil)
	recorder := httptest.NewRecorder()
	suite.router.ServeHTTP(recorder, req)

	suite.Assert().Equal(http.StatusOK, recorder.Code)
	suite.Assert().Contains(recorder.Body.String(), "Task 1")
}

func (suite *ControllerSuite) TestGetTask_TaskNotFound_Negative() {
	id := primitive.NewObjectID()

	suite.taskusecase.On("GetTask", mock.Anything, id).Return(domain.Task{}, errors.New("Task not found")).Once()

	req, _ := http.NewRequest(http.MethodGet, "/tasks/"+id.Hex(), nil)
	recorder := httptest.NewRecorder()
	suite.router.ServeHTTP(recorder, req)

	suite.Assert().Equal(http.StatusNotFound, recorder.Code)
	suite.Assert().Contains(recorder.Body.String(), "Task Not Found")
}

// func (suite *ControllerSuite) TestAddTask_Positive() {
// 	task := &domain.Task{
// 		Title:       "New Task",
// 		Description: "New Description",
// 		Status:      "Not started",
// 		DueDate:     time.Now(),
// 	}

// 	suite.taskusecase.On("AddTask", mock.Anything, task).Return(nil).Once()

// 	jsonInput, _ := json.Marshal(task)
// 	req, _ := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(jsonInput))
// 	req.Header.Set("Content-Type", "application/json")

// 	recorder := httptest.NewRecorder()
// 	suite.router.ServeHTTP(recorder, req)

// 	suite.Assert().Equal(http.StatusCreated, recorder.Code)
// 	suite.Assert().Contains(recorder.Body.String(), "Task Created")
// }

func (suite *ControllerSuite) TestAddTask_InvalidInput_Negative() {
	task := &domain.Task{
		Title:       "",
		Description: "",
		Status:      "",
	}

	jsonInput, _ := json.Marshal(task)
	req, _ := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(jsonInput))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	suite.router.ServeHTTP(recorder, req)

	suite.Assert().Equal(http.StatusBadRequest, recorder.Code)
	suite.Assert().Contains(recorder.Body.String(), "Title, Description, and Status are required fields")
}

func (suite *ControllerSuite) TestRemoveTask_Positive() {
	id := primitive.NewObjectID()

	suite.taskusecase.On("RemoveTask", mock.Anything, id).Return(nil).Once()

	req, _ := http.NewRequest(http.MethodDelete, "/tasks/"+id.Hex(), nil)
	recorder := httptest.NewRecorder()
	suite.router.ServeHTTP(recorder, req)

	suite.Assert().Equal(http.StatusOK, recorder.Code)
	suite.Assert().Contains(recorder.Body.String(), "Task removed")
}

func (suite *ControllerSuite) TestRemoveTask_TaskNotFound_Negative() {
	id := primitive.NewObjectID()

	suite.taskusecase.On("RemoveTask", mock.Anything, id).Return(errors.New("Task not found")).Once()

	req, _ := http.NewRequest(http.MethodDelete, "/tasks/"+id.Hex(), nil)
	recorder := httptest.NewRecorder()
	suite.router.ServeHTTP(recorder, req)

	suite.Assert().Equal(http.StatusNotFound, recorder.Code)
	suite.Assert().Contains(recorder.Body.String(), "Task Not Found")
}

func (suite *ControllerSuite) TestUpdateTask_Positive() {
	id := primitive.NewObjectID()
	originalTask := &domain.Task{
		Title:       "Old Task",
		Description: "Old Description",
		Status:      "Not started",
	}
	updatedTask := &domain.Task{
		Title:       "Updated Task",
		Description: "Updated Description",
		Status:      "In Progress",
	}

	suite.taskusecase.On("GetTask", mock.Anything, id).Return(*originalTask, nil).Once()
	suite.taskusecase.On("UpdateTask", mock.Anything, id, updatedTask).Return(nil).Once()

	jsonInput, _ := json.Marshal(updatedTask)
	req, _ := http.NewRequest(http.MethodPut, "/tasks/"+id.Hex(), bytes.NewBuffer(jsonInput))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	suite.router.ServeHTTP(recorder, req)

	suite.Assert().Equal(http.StatusOK, recorder.Code)
	suite.Assert().Contains(recorder.Body.String(), "Task Updated")
}

// func (suite *ControllerSuite) TestUpdateTask_EmptyFields_Positive() {
// 	id := primitive.NewObjectID()
// 	updatedTask := domain.Task{}
// 	suite.taskusecase.On("UpdateTask", mock.Anything, id, &updatedTask).Return(nil).Once()

// 	jsonInput, _ := json.Marshal(updatedTask)
// 	req, _ := http.NewRequest(http.MethodPut, "/tasks/"+id.Hex(), bytes.NewBuffer(jsonInput))
// 	req.Header.Set("Content-Type", "application/json")

// 	recorder := httptest.NewRecorder()
// 	suite.router.ServeHTTP(recorder, req)

// 	suite.Assert().Equal(http.StatusOK, recorder.Code)
// 	suite.Assert().Contains(recorder.Body.String(), "Task updated successfully")
// }

// func (suite *ControllerSuite) TestUpdateTask_TaskNotFound_Negative() {
// 	id := primitive.NewObjectID()

// 	updatedTask := domain.Task{}

// 	suite.taskusecase.On("GetTask", mock.Anything, id).Return(domain.Task{}, errors.New("Task not found")).Once()

// 	jsonInput, _ := json.Marshal(updatedTask)
// 	req, _ := http.NewRequest(http.MethodPut, "/tasks/"+id.Hex(), bytes.NewBuffer(jsonInput))
// 	req.Header.Set("Content-Type", "application/json")

// 	recorder := httptest.NewRecorder()
// 	suite.router.ServeHTTP(recorder, req)

// 	suite.Assert().Equal(http.StatusBadRequest, recorder.Code)
// 	suite.Assert().Contains(recorder.Body.String(), "Task not found")
// }

func TestControllerSuite(t *testing.T) {
	suite.Run(t, new(ControllerSuite))
}
