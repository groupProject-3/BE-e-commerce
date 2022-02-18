package user

import (
	"be/delivery/controllers/auth"
	"be/delivery/middlewares"
	"be/delivery/templates"
	"be/models"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type MockAuthLib struct{}

func (m *MockAuthLib) Login(UserLogin templates.Userlogin) (models.User, error) {
	return models.User{Model: gorm.Model{ID: 1}, Email: UserLogin.Email, Password: UserLogin.Password}, nil
}

type MockUserLib struct{}

func (m *MockUserLib) Create(user models.User) (models.User, error) {
	return models.User{}, nil
}

func (m *MockUserLib) GetById(id int) (models.User, error) {
	return models.User{}, nil
}

func (m *MockUserLib) UpdateById(id int, userUp templates.UserRequest) (models.User, error) {
	return models.User{}, nil
}

func (m *MockUserLib) DeleteById(id int) (gorm.DeletedAt, error) {
	user := models.User{}
	return user.DeletedAt, nil
}

func (m *MockUserLib) GetAll() ([]templates.UserResponse, error) {
	return []templates.UserResponse{}, nil
}

type MockFailUserLib struct{}

func (m *MockFailUserLib) Create(user models.User) (models.User, error) {
	return models.User{}, errors.New("")
}

func (m *MockFailUserLib) GetById(id int) (models.User, error) {
	return models.User{}, errors.New("")
}

func (m *MockFailUserLib) UpdateById(id int, userUp templates.UserRequest) (models.User, error) {
	return models.User{}, errors.New("")
}

func (m *MockFailUserLib) DeleteById(id int) (gorm.DeletedAt, error) {
	user := models.User{}
	return user.DeletedAt, errors.New("")
}

func (m *MockFailUserLib) GetAll() ([]templates.UserResponse, error) {
	return []templates.UserResponse{}, errors.New("")
}

func TestCreate(t *testing.T) {
	t.Run("BadRequest", func(t *testing.T) {
		e := echo.New()

		reqBody, _ := json.Marshal(map[string]interface{}{
			"name":     1,
			"email":    "",
			"password": "anonim123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := New(&MockUserLib{})
		userController.Create()(context)

		response := templates.GetUserResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 400, response.Code)
		assert.Equal(t, "error in request for create new user", response.Message)
	})

	t.Run("InternalServerError", func(t *testing.T) {
		e := echo.New()

		reqBody, _ := json.Marshal(map[string]string{
			"name":     "",
			"email":    "anonim@123",
			"password": "anonim123",
		})
		log.Info(reqBody)
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users")
		log.Info(req)
		log.Info(context)
		userController := New(&MockFailUserLib{})
		userController.Create()(context)

		response := templates.GetUserResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 500, response.Code)
		assert.Equal(t, "error internal server error fo create new user", response.Message)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()

		reqBody, _ := json.Marshal(map[string]string{
			"name":     "anonim123",
			"email":    "anonim@123",
			"password": "anonim123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := New(&MockUserLib{})
		userController.Create()(context)

		response := templates.GetUserResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 201, response.Code)
		assert.Equal(t, "Success create new user", response.Message)
	})
}

func TestGetAll(t *testing.T) {
	var jwtToken string

	t.Run("Success Login", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(map[string]string{
			"name":     "admin",
			"email":    "admin@admin.com",
			"password": "admin",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		authController := auth.New(&MockAuthLib{})
		authController.Login()(context)

		response := templates.LoginRespFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		jwtToken = response.Data["token"].(string)

		assert.Equal(t, response.Message, "success login")
		assert.NotNil(t, response.Data["token"])
	})

	t.Run("Success", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := New(&MockUserLib{})
		if err := middlewares.JwtMiddleware()(userController.GetAll())(context); err != nil {
			return
		}

		response := templates.GetUserResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "Success get all user", response.Message)
	})

	t.Run("InternalServerError", func(t *testing.T) {
		// token := string(jwtToken)
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		// userController := New(&MockFalseLib{})
		userController := New(&MockFailUserLib{})

		if err := middlewares.JwtMiddleware()(userController.GetAll())(context); err != nil {
			return
		}

		response := templates.GetUserResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 400, response.Code)
		assert.Equal(t, "error internal server error request get all user", response.Message)
	})
}


