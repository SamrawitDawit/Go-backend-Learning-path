package infrastructure

import (
	"net/http"
	"net/http/httptest"
	domain "task8-Testing/Domain"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockJwtService struct {
	mock.Mock
}

func (m *MockJwtService) GenerateToken(user domain.User) (string, error) {
	args := m.Called(user)
	return args.String(0), args.Error(1)
}

func (m *MockJwtService) CheckToken(authPart string) (*jwt.Token, error) {
	args := m.Called(authPart)
	return args.Get(0).(*jwt.Token), args.Error(1)
}

func (m *MockJwtService) FindClaim(token *jwt.Token) (jwt.MapClaims, bool) {
	args := m.Called(token)
	return args.Get(0).(jwt.MapClaims), args.Bool(1)
}

func TestAuthMiddleWare(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockJWTService := new(MockJwtService)
	router := gin.New()
	router.Use(AuthMiddleWare(mockJWTService))
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	})

	t.Run("Missing Authorization Header", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusUnauthorized, recorder.Code)
		assert.Contains(t, recorder.Body.String(), "Authorization header is required")
	})

	t.Run("Invalid Authorization Header", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		req.Header.Set("Authorization", "InvalidToken")
		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusUnauthorized, recorder.Code)
		assert.Contains(t, recorder.Body.String(), "Invalid Authorization header")
	})

	t.Run("Invalid Token", func(t *testing.T) {
		mockToken := &jwt.Token{Valid: false}
		mockJWTService.On("CheckToken", mock.Anything).Return(mockToken, nil).Once()

		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		req.Header.Set("Authorization", "Bearer InvalidToken")
		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusUnauthorized, recorder.Code)
		assert.Contains(t, recorder.Body.String(), "Invalid or expired token")
	})

	t.Run("Valid Token with Valid CLaims", func(t *testing.T) {
		mockJWTService.On("CheckToken", "ValidToken").Return(&jwt.Token{Valid: true}, nil)
		mockJWTService.On("FindClaim", mock.Anything).Return(jwt.MapClaims{"role": "user"}, true)

		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		req.Header.Set("Authorization", "Bearer ValidToken")
		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Contains(t, recorder.Body.String(), "Success")
	})
}

func TestAdminMiddlWare(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.Use(func(c *gin.Context) {
		c.Set("role", "user")
	})
	router.Use(AdminMiddleWare())
	router.GET("/admin", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome Admin"})
	})

	t.Run("Non-Admin Role", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/admin", nil)
		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusForbidden, recorder.Code)
		assert.Contains(t, recorder.Body.String(), "Sorry, you are not eligible to do this.")
	})
	t.Run("Admin Role", func(t *testing.T) {
		router := gin.New()
		router.Use(func(c *gin.Context) {
			c.Set("role", "admin")
		})
		router.Use(AdminMiddleWare())
		router.GET("/admin", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Welcome Admin"})
		})

		req, _ := http.NewRequest(http.MethodGet, "/admin", nil)
		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Contains(t, recorder.Body.String(), "Welcome Admin")
	})
}
