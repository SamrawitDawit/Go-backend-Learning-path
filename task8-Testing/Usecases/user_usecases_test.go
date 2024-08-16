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

type userUsecaseSuite struct {
	suite.Suite
	repository *mocks.UserRepository
	usecase    domain.UserUsecase
}

func (suite *userUsecaseSuite) SetupTest() {
	repo := new(mocks.UserRepository)
	usecase := NewUserUsecase(repo, time.Second)
	suite.repository = repo
	suite.usecase = usecase
}

func (suite *userUsecaseSuite) TestRegister_Positive() {
	user := domain.User{
		Username: "samrawit",
		Password: "samrawit",
	}
	suite.repository.On("Register", mock.Anything, &user).Return(nil)
	err := suite.usecase.Register(context.Background(), &user)
	suite.NoError(err, "no error when registering with valid input")
	suite.repository.AssertExpectations(suite.T())
}

func (suite *userUsecaseSuite) TestRegister_Negative() {
	user := domain.User{
		Username: "samrawit",
		Password: "samrawit",
	}
	suite.repository.On("Register", mock.Anything, &user).Return(errors.New("username already exists"))
	err := suite.usecase.Register(context.Background(), &user)
	suite.Error(err, "error when registering with the same username")
	suite.repository.AssertExpectations(suite.T())
}

func (suite *userUsecaseSuite) TestLogin_Positive() {
	user := domain.User{
		Username: "samrawit",
		Password: "samrawit",
	}
	suite.repository.On("Register", mock.Anything, &user).Return(nil)
	err := suite.usecase.Register(context.Background(), &user)
	suite.NoError(err, "no error when registering with valid input")

	loginUser := domain.User{
		Username: "samrawit",
		Password: "samrawit",
	}
	suite.repository.On("Login", mock.Anything, &loginUser).Return("token", nil)
	_, err = suite.usecase.Login(context.Background(), &loginUser)
	suite.NoError(err, "no error when logging in with valid input")
	suite.repository.AssertExpectations(suite.T())
}

func (suite *userUsecaseSuite) TestLogin_InvalidPassword_Negative() {
	user := domain.User{
		Username: "samrawit",
		Password: "samrawit",
	}
	suite.repository.On("Register", mock.Anything, &user).Return(nil)
	err := suite.usecase.Register(context.Background(), &user)
	suite.NoError(err, "no error when registering with valid input")

	unauthenticatedUser := domain.User{
		Username: "samrawit",
		Password: "something",
	}
	suite.repository.On("Login", mock.Anything, &unauthenticatedUser).Return("", errors.New("invalid password"))
	_, err = suite.usecase.Login(context.Background(), &unauthenticatedUser)
	suite.Error(err, "error when logging in with invalid input")
	suite.repository.AssertExpectations(suite.T())
}

func (suite *userUsecaseSuite) TestLogin_UserNotRegistered_Negative() {
	loginUser := domain.User{
		Username: "samrawit",
		Password: "samrawit",
	}
	suite.repository.On("Login", mock.Anything, &loginUser).Return("", errors.New("user not found"))
	_, err := suite.usecase.Login(context.Background(), &loginUser)
	suite.Error(err, "error when logging in with invalid input")
	suite.repository.AssertExpectations(suite.T())
}

func (suite *userUsecaseSuite) TestPromote_Positive() {
	id, err := primitive.ObjectIDFromHex("66b74e21d7aa916a022b8479")
	suite.Require().NoError(err, "no error when generating valid ObjectID")
	suite.repository.On("Promote", mock.Anything, id).Return(nil)
	err = suite.usecase.Promote(context.Background(), id)
	suite.NoError(err, "no error when promoting with valid id")
	suite.repository.AssertExpectations(suite.T())
}

func (suite *userUsecaseSuite) TestPromote_InvalidID_Negative() {
	id, err := primitive.ObjectIDFromHex("66b74e21d7aa916a022b8479")
	suite.Require().NoError(err, "no error when generating valid ObjectID")
	suite.repository.On("Promote", mock.Anything, id).Return(errors.New("user not found"))
	err = suite.usecase.Promote(context.Background(), id)
	suite.Error(err, "error when promoting with invalid id")
	suite.repository.AssertExpectations(suite.T())
}

func TestUserUsecase(t *testing.T) {
	suite.Run(t, new(userUsecaseSuite))
}
