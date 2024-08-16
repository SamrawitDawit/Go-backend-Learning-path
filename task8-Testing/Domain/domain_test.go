package domain

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestTaskStruct(t *testing.T) {
	id := primitive.NewObjectID()
	dueDate := time.Now().Add(24 * time.Hour)
	task := Task{
		ID:          id,
		Title:       "New Task",
		Description: "This is a new task.",
		DueDate:     dueDate,
		Status:      "Pending",
	}

	assert.Equal(t, id, task.ID)
	assert.Equal(t, "New Task", task.Title)
	assert.Equal(t, "This is a new task.", task.Description)
	assert.Equal(t, dueDate, task.DueDate)
	assert.Equal(t, "Pending", task.Status)
}

func TestUserStruct(t *testing.T) {
	id := primitive.NewObjectID()
	user := User{
		ID:       id,
		Username: "testuser",
		Password: "hashedpassword",
		Role:     "admin",
	}

	assert.Equal(t, id, user.ID)
	assert.Equal(t, "testuser", user.Username)
	assert.Equal(t, "hashedpassword", user.Password)
	assert.Equal(t, "admin", user.Role)
}

func TestTaskJSONSerialization(t *testing.T) {
	id := primitive.NewObjectID()
	dueDate := time.Date(2023, time.August, 13, 15, 30, 0, 0, time.UTC)
	task := Task{
		ID:          id,
		Title:       "New Task",
		Description: "This is a new task.",
		DueDate:     dueDate,
		Status:      "Pending",
	}

	jsonData, err := json.Marshal(task)
	assert.NoError(t, err)

	var deserializedTask Task
	err = json.Unmarshal(jsonData, &deserializedTask)
	assert.NoError(t, err)

	assert.Equal(t, task, deserializedTask)
}

func TestUserJSONSerialization(t *testing.T) {
	id := primitive.NewObjectID()
	user := User{
		ID:       id,
		Username: "testuser",
		Password: "hashedpassword",
		Role:     "admin",
	}

	jsonData, err := json.Marshal(user)
	assert.NoError(t, err)

	var deserializedUser User
	err = json.Unmarshal(jsonData, &deserializedUser)
	assert.NoError(t, err)

	assert.Equal(t, user, deserializedUser)
}
