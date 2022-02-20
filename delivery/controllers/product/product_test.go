package product

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

type MockProdLib struct{}

func (m *MockProdLib) Create(user_id uint, newPro models.Product) (models.Product, error) {
	return models.Product{}, nil
}

func (m *MockProdLib) UpdateById(id int, user_id uint, upPro templates.ProductRequest) (models.Product, error) {
	return models.Product{}, nil
}

func (m *MockProdLib) DeleteById(id int, user_id uint) (gorm.DeletedAt, error) {
	prod := models.Product{}
	return prod.DeletedAt, nil
}

func (m *MockProdLib) GetAllMe(user_id uint) ([]templates.ProductResponse, error) {
	return []templates.ProductResponse{}, nil
}

func (m *MockProdLib) GetByIdMe(id int, user_id uint) (templates.ProductResponse, error) {
	return templates.ProductResponse{}, nil
}

func (m *MockProdLib) GetAll() ([]templates.ProductResponse, error) {
	return []templates.ProductResponse{}, nil
}

func (m *MockProdLib) GetById(id int) (templates.ProductResponse, error) {
	return templates.ProductResponse{}, nil
}

func (m *MockProdLib) UpdateByIdAll(id int, upPro templates.ProductRequest) (models.Product, error) {
	return models.Product{}, nil
}

type MockFailProdLib struct{}

func (m *MockFailProdLib) Create(user_id uint, newPro models.Product) (models.Product, error) {
	return models.Product{}, errors.New("")
}

func (m *MockFailProdLib) UpdateById(id int, user_id uint, upPro templates.ProductRequest) (models.Product, error) {
	return models.Product{}, errors.New("")
}

func (m *MockFailProdLib) DeleteById(id int, user_id uint) (gorm.DeletedAt, error) {
	prod := models.Product{}
	return prod.DeletedAt, errors.New("")
}

func (m *MockFailProdLib) GetAllMe(user_id uint) ([]templates.ProductResponse, error) {
	return []templates.ProductResponse{}, errors.New("")
}

func (m *MockFailProdLib) GetByIdMe(id int, user_id uint) (templates.ProductResponse, error) {
	return templates.ProductResponse{}, errors.New("")
}

func (m *MockFailProdLib) GetAll() ([]templates.ProductResponse, error) {
	return []templates.ProductResponse{}, errors.New("")
}

func (m *MockFailProdLib) GetById(id int) (templates.ProductResponse, error) {
	return templates.ProductResponse{}, errors.New("")
}

func (m *MockFailProdLib) UpdateByIdAll(id int, upPro templates.ProductRequest) (models.Product, error) {
	return models.Product{}, errors.New("")
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

		Controller := New(&MockFailProdLib{})
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
		reqBody, _ := json.Marshal(map[string]interface{}{
			"name":         "anonim",
			"product_type": 1,
			"price":        1000,
			"qty":          10,
			"description":  "anonim description",
		})
		log.Info(reqBody)
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/product/type")
		log.Info(req)
		log.Info(context)
		Controller := New(&MockFailProdLib{})
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
		reqBody, _ := json.Marshal(map[string]interface{}{
			"name":         "anonim",
			"product_type": 1,
			"price":        1000,
			"qty":          10,
			"description":  "anonim description",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/product/type")

		Controller := New(&MockProdLib{})
		if err := middlewares.JwtMiddleware()(Controller.Create())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := templates.GetProdTypeResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 201, response.Code)
	})
}

func TestUpdateById(t *testing.T) {
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

		Controller := New(&MockFailProdLib{})
		if err := middlewares.JwtMiddleware()(Controller.UpdateById())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := templates.GetProdTypeResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 400, response.Code)
	})

	t.Run("InternalServerError", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(map[string]interface{}{
			"name":         "anonim",
			"product_type": 1,
			"price":        1000,
			"qty":          10,
			"description":  "anonim description",
		})
		log.Info(reqBody)
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/product/type")
		log.Info(req)
		log.Info(context)
		Controller := New(&MockFailProdLib{})
		if err := middlewares.JwtMiddleware()(Controller.UpdateById())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := templates.GetProdTypeResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 500, response.Code)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(map[string]interface{}{
			"name":         "anonim",
			"product_type": 1,
			"price":        1000,
			"qty":          10,
			"description":  "anonim description",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/product/type")

		Controller := New(&MockProdLib{})
		if err := middlewares.JwtMiddleware()(Controller.UpdateById())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := templates.GetProdTypeResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 202, response.Code)
	})
}

func TestDeleteById(t *testing.T) {
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

	t.Run("InternalServerError", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/product/type")
		log.Info(req)
		log.Info(context)
		Controller := New(&MockFailProdLib{})
		if err := middlewares.JwtMiddleware()(Controller.DeleteById())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := templates.GetProdTypeResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 500, response.Code)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/product/type")

		Controller := New(&MockProdLib{})
		if err := middlewares.JwtMiddleware()(Controller.DeleteById())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := templates.GetProdTypeResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 200, response.Code)
	})
}

func TestGetAllMe(t *testing.T) {
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

	t.Run("InternalServerError", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/product/type")
		log.Info(req)
		log.Info(context)
		Controller := New(&MockFailProdLib{})
		if err := middlewares.JwtMiddleware()(Controller.GetAllMe())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := templates.GetProdTypeResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 500, response.Code)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/product/type")

		Controller := New(&MockProdLib{})
		if err := middlewares.JwtMiddleware()(Controller.GetAllMe())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := templates.GetProdTypeResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 200, response.Code)
	})
}

func TestGetByIdMe(t *testing.T) {
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

	t.Run("InternalServerError", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/product/type")
		log.Info(req)
		log.Info(context)
		Controller := New(&MockFailProdLib{})
		if err := middlewares.JwtMiddleware()(Controller.GetByIdMe())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := templates.GetProdTypeResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 500, response.Code)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/product/type")

		Controller := New(&MockProdLib{})
		if err := middlewares.JwtMiddleware()(Controller.GetByIdMe())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := templates.GetProdTypeResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 200, response.Code)
	})
}

func TestGetAll(t *testing.T) {
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

	t.Run("InternalServerError", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/product/type")
		log.Info(req)
		log.Info(context)
		Controller := New(&MockFailProdLib{})
		if err := middlewares.JwtMiddleware()(Controller.GetAll())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := templates.GetProdTypeResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 500, response.Code)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/product/type")

		Controller := New(&MockProdLib{})
		if err := middlewares.JwtMiddleware()(Controller.GetAll())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := templates.GetProdTypeResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 200, response.Code)
	})
}

func TestGetById(t *testing.T) {
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

	t.Run("InternalServerError", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/product/type")
		log.Info(req)
		log.Info(context)
		Controller := New(&MockFailProdLib{})
		if err := middlewares.JwtMiddleware()(Controller.GetById())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := templates.GetProdTypeResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 500, response.Code)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/product/type")

		Controller := New(&MockProdLib{})
		if err := middlewares.JwtMiddleware()(Controller.GetById())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := templates.GetProdTypeResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 200, response.Code)
	})
}
