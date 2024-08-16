// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "task8-Testing/Domain"

	mock "github.com/stretchr/testify/mock"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// TaskUsecase is an autogenerated mock type for the TaskUsecase type
type TaskUsecase struct {
	mock.Mock
}

// AddTask provides a mock function with given fields: c, task
func (_m *TaskUsecase) AddTask(c context.Context, task *domain.Task) error {
	ret := _m.Called(c, task)

	if len(ret) == 0 {
		panic("no return value specified for AddTask")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Task) error); ok {
		r0 = rf(c, task)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTask provides a mock function with given fields: c, taskID
func (_m *TaskUsecase) GetTask(c context.Context, taskID primitive.ObjectID) (domain.Task, error) {
	ret := _m.Called(c, taskID)

	if len(ret) == 0 {
		panic("no return value specified for GetTask")
	}

	var r0 domain.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) (domain.Task, error)); ok {
		return rf(c, taskID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) domain.Task); ok {
		r0 = rf(c, taskID)
	} else {
		r0 = ret.Get(0).(domain.Task)
	}

	if rf, ok := ret.Get(1).(func(context.Context, primitive.ObjectID) error); ok {
		r1 = rf(c, taskID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTasks provides a mock function with given fields: c
func (_m *TaskUsecase) GetTasks(c context.Context) ([]domain.Task, error) {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for GetTasks")
	}

	var r0 []domain.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.Task, error)); ok {
		return rf(c)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.Task); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveTask provides a mock function with given fields: c, taskID
func (_m *TaskUsecase) RemoveTask(c context.Context, taskID primitive.ObjectID) error {
	ret := _m.Called(c, taskID)

	if len(ret) == 0 {
		panic("no return value specified for RemoveTask")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) error); ok {
		r0 = rf(c, taskID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTask provides a mock function with given fields: c, taskID, updatedTask
func (_m *TaskUsecase) UpdateTask(c context.Context, taskID primitive.ObjectID, updatedTask *domain.Task) error {
	ret := _m.Called(c, taskID, updatedTask)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTask")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID, *domain.Task) error); ok {
		r0 = rf(c, taskID, updatedTask)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewTaskUsecase creates a new instance of TaskUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTaskUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *TaskUsecase {
	mock := &TaskUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}