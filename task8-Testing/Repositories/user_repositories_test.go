package repositories

import (
	"context"
	"fmt"
	domain "task8-Testing/Domain"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type userRepositorySuite struct {
	suite.Suite
	repository domain.UserRepository
	mongoC     testcontainers.Container
	client     *mongo.Client
}

func (suite *userRepositorySuite) setupMongoContainer() (testcontainers.Container, *mongo.Client) {
	ctx := context.Background()

	mongoC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "mongo:latest",
			ExposedPorts: []string{"27017/tcp"},
			WaitingFor:   wait.ForListeningPort("27017/tcp"),
		},
		Started: true,
	})
	if err != nil {
		suite.T().Fatalf("could not start mongo container: %v", err)
	}
	host, err := mongoC.Host(ctx)
	if err != nil {
		suite.T().Fatalf("could not get mongo container host: %v", err)
	}
	port, err := mongoC.MappedPort(ctx, "27017")
	if err != nil {
		suite.T().Fatalf("could not get mongo container port : %v", err)
	}

	mongoURI := fmt.Sprintf("mongodb://%s:%s", host, port.Port())
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))

	if err != nil {
		suite.T().Fatalf("couldn't connect to mongo: %v", err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		suite.T().Fatalf("could not ping mongo: %v", err)
	}
	return mongoC, client
}
func (suite *userRepositorySuite) SetupSuite() {
	mongoC, client := suite.setupMongoContainer()
	db := client.Database("test_db")
	repo := NewUserRepository(*db, "tasks")
	suite.repository = repo
	suite.mongoC = mongoC
	suite.client = client
}

func (suite *userRepositorySuite) TearDownTest() {
	err := suite.client.Disconnect(context.Background())
	if err != nil {
		suite.T().Fatalf("couldn't disconnect from mongo: %v", err)
	}
	err = suite.mongoC.Terminate(context.Background())
	if err != nil {
		suite.T().Fatalf("couldn't terminate mongo container: %v", err)
	}
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
