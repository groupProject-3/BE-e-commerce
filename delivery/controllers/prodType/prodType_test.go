package prodType

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

type mockProdTypeLib struct{}

func (m *mockProdTypeLib) Create(proType models.ProductType) (models.ProductType, error) {
	return models.ProductType{}, nil
}

func (m *mockProdTypeLib) UpdateById(id int, upPro templates.ProductTypeRequest) (models.ProductType, error) {
	return models.ProductType{}, nil
}

func (m *mockProdTypeLib) DeleteById(id int) (gorm.DeletedAt, error) {
	prodType := models.ProductType{}

	return prodType.DeletedAt, nil
}

func (m *mockProdTypeLib) GetAll() ([]templates.ProductTypeResponse, error) {
	return []templates.ProductTypeResponse{}, nil
}

type mockFailProdTypeLib struct{}

func (m *mockFailProdTypeLib) Create(proType models.ProductType) (models.ProductType, error) {
	return models.ProductType{}, errors.New("")
}

func (m *mockFailProdTypeLib) UpdateById(id int, upPro templates.ProductTypeRequest) (models.ProductType, error) {
	return models.ProductType{}, errors.New("")
}

func (m *mockFailProdTypeLib) DeleteById(id int) (gorm.DeletedAt, error) {
	prodType := models.ProductType{}

	return prodType.DeletedAt, errors.New("")
}

func (m *mockFailProdTypeLib) GetAll() ([]templates.ProductTypeResponse, error) {
	return []templates.ProductTypeResponse{}, errors.New("")
}

func TestCreate(t *testing.T) {
	var jwtToken string

	t.Run("Success Login", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(map[string]string{
			"name":     "admin",
			"email":    "admin@gmail.com",
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

	t.Run("BadRequest", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(map[string]int{"name": 1})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content=Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/product/type")

		Controller := New(&mockFailProdTypeLib{})
		if err := middlewares.JwtMiddleware()(Controller.Create())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := templates.GetProdTypeResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 400, response.Code)
	})

	t.Run("InternalServerError", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(map[string]string{
			"name":"anonim",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content=Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/product/type")
		log.Info(req)
		log.Info(context)
		Controller := New(&mockFailProdTypeLib{})
		if err := middlewares.JwtMiddleware()(Controller.Create())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := templates.GetProdTypeResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 500, response.Code)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(map[string]string{
			"name":"anonim",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content=Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/product/type")

		Controller := New(&mockProdTypeLib{})
		if err := middlewares.JwtMiddleware()(Controller.Create())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := templates.GetProdTypeResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 201, response.Code)
	})

}
