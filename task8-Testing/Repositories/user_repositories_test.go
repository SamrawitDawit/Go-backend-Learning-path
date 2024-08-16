package repositories

import (
	"context"
	"log"
	domain "task8-Testing/Domain"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepositorySuite struct {
	suite.Suite
	repository domain.UserRepository
}

func (suite *userRepositorySuite) SetupSuite() {
	clientOption := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("test_db")
	repo := NewUserRepository(*db, "users")
	suite.repository = repo
}

func (suite *userRepositorySuite) TearDownTest() {
	suite.repository.(*userRepository).database.Collection("users").Drop(context.Background())
}

func (suite *userRepositorySuite) TestRegister_Positive() {
	user := domain.User{
		Username: "samrawit",
		Password: "samrawit",
	}
	err := suite.repository.Register(context.Background(), &user)
	suite.NoError(err, "no error when register with valid input")
}

func (suite *userRepositorySuite) TestRegister_Negative() {
	user := domain.User{
		Username: "samrawit",
		Password: "samrawit",
	}
	err := suite.repository.Register(context.Background(), &user)
	suite.NoError(err, "no error when register with valid input")
	err = suite.repository.Register(context.Background(), &user)
	suite.Error(err, "error when register with the same username")
}

func (suite *userRepositorySuite) TestLogin_Positive() {
	user := domain.User{
		Username: "samrawit",
		Password: "samrawit",
	}
	err := suite.repository.Register(context.Background(), &user)
	suite.NoError(err, "no error when register with valid input")
	login_user := domain.User{
		Username: "samrawit",
		Password: "samrawit",
	}
	_, err = suite.repository.Login(context.Background(), &login_user)
	suite.NoError(err, "no error when login with valid input")
}

func (suite *userRepositorySuite) TestLogin_InvalidPassword_Negative() {
	user := domain.User{
		Username: "samrawit",
		Password: "samrawit",
	}
	err := suite.repository.Register(context.Background(), &user)
	suite.NoError(err, "no error when register with valid input")
	unauthenticatedUser := domain.User{
		Username: "samrawit",
		Password: "something",
	}
	_, err = suite.repository.Login(context.Background(), &unauthenticatedUser)
	suite.Error(err, "error when login with invalid input")
}

func (suite *userRepositorySuite) TestLogin_UserNotRegistered_Negative() {
	user := domain.User{
		Username: "samrawit",
		Password: "samrawit",
	}
	_, err := suite.repository.Login(context.Background(), &user)
	suite.Error(err, "error when login with invalid input")
}

func (suite *userRepositorySuite) TestPromote_Positive() {
	admin := domain.User{
		Username: "samrawit",
		Password: "samrawit",
	}
	err := suite.repository.Register(context.Background(), &admin)
	suite.NoError(err, "no error when register with valid input")
	id, err := primitive.ObjectIDFromHex("66b74e21d7aa916a022b8479")
	suite.Assert().NoError(err)
	regular_user := domain.User{
		ID:       id,
		Username: "rediet",
		Password: "rediet",
	}
	err = suite.repository.Register(context.Background(), &regular_user)
	suite.NoError(err, "no error when register with valid input")
	suite.Assert().Equal("regular", regular_user.Role, "the role of a user, who registered while the database is not empty should be regular")
	err = suite.repository.Promote(context.Background(), regular_user.ID)
	suite.NoError(err, "no error when promoting with valid id")
	promoted_user, err := suite.repository.GetUser(context.Background(), id)
	suite.Assert().NoError(err)
	suite.Assert().Equal("admin", promoted_user.Role, "the role of a user should be admin after promotion.")
}
func (suite *userRepositorySuite) TestPromote_InvalidID_Negative() {
	id, err := primitive.ObjectIDFromHex("66b74e21d7aa916a022b8479")
	suite.Assert().NoError(err)
	err = suite.repository.Promote(context.Background(), id)
	suite.Error(err, "error when promoting with invalid id")
}

func TestUserRepository(t *testing.T) {
	suite.Run(t, new(userRepositorySuite))
}
