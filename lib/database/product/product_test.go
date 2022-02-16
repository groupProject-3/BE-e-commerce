package product

import (
	"be/configs"
	"be/delivery/controllers/product"
	prodType "be/lib/database/prodType"
	"be/lib/database/user"
	"be/models"
	"be/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&models.Product{})
	db.Migrator().CreateTable(&models.Product{})

	t.Run("success run create", func(t *testing.T) {
		mockUser1 := models.User{Name: "anonim2", Email: "anonim2", Password: "anonim2"}
		if _, err := user.New(db).Create(mockUser1); err != nil {
			t.Fatal()
		}
		mockProT1 := models.ProductType{Name: "anonim1"}
		if _, err := prodType.New(db).Create(mockProT1); err != nil {
			t.Fatal()
		}
		mockProd := models.Product{Product_type_id: 1, Name: "anonim1 product", Price: 1000, Qty: 10, Description: "anonim1 Description"}

		res, err := repo.Create(1, mockProd)
		assert.Nil(t, err)
		assert.NotNil(t, res)
	})
	mockProd1 := models.Product{Product_type_id: 1, Name: "anonim1 product", Price: 1000, Qty: 10, Description: "anonim1 Description"}
	if _, err := repo.Create(1, mockProd1); err != nil {
		t.Fatal()
	}

	t.Run("fail run create", func(t *testing.T) {
		mockProd2 := models.Product{Model: gorm.Model{ID: 1}, Product_type_id: 1, Name: "anonim1 product", Price: 1000, Qty: 10, Description: "anonim1 Description"}
		_, err := repo.Create(1, mockProd2)
		assert.NotNil(t, err)
	})
}

func TestUpdateById(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&models.Product{})
	db.Migrator().CreateTable(&models.Product{})

	t.Run("success run UpdateById", func(t *testing.T) {
		mockUser1 := models.User{Name: "anonim2", Email: "anonim2", Password: "anonim2"}
		if _, err := user.New(db).Create(mockUser1); err != nil {
			t.Fatal()
		}
		mockProT1 := models.ProductType{Name: "anonim1"}
		if _, err := prodType.New(db).Create(mockProT1); err != nil {
			t.Fatal()
		}
		mockProd1 := models.Product{Product_type_id: 1, Name: "anonim1 product", Price: 1000, Qty: 10, Description: "anonim1 Description"}
		if _, err := repo.Create(1, mockProd1); err != nil {
			t.Fatal()
		}
		mockProd := product.ProductRequest{Product_type_id: 1, Name: "anonim2 product", Price: 1000, Qty: 10, Description: "anonim2 Description"}

		res, err := repo.UpdateById(1, 1, mockProd)
		assert.Nil(t, err)
		assert.Equal(t, "anonim2 product", res.Name)
	})

	t.Run("fail run UpdateById", func(t *testing.T) {
		mockProd := product.ProductRequest{Product_type_id: 1, Name: "anonim2 product", Price: 1000, Qty: 10, Description: "anonim2 Description"}

		_, err := repo.UpdateById(10, 1, mockProd)
		assert.NotNil(t, err)
	})
}

func TestDeleteById(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&models.Product{})
	db.Migrator().CreateTable(&models.Product{})

	t.Run("success run DeleteById", func(t *testing.T) {
		mockUser1 := models.User{Name: "anonim2", Email: "anonim2", Password: "anonim2"}
		if _, err := user.New(db).Create(mockUser1); err != nil {
			t.Fatal()
		}
		mockProT1 := models.ProductType{Name: "anonim1"}
		if _, err := prodType.New(db).Create(mockProT1); err != nil {
			t.Fatal()
		}
		mockProd1 := models.Product{Product_type_id: 1, Name: "anonim1 product", Price: 1000, Qty: 10, Description: "anonim1 Description"}
		if _, err := repo.Create(1, mockProd1); err != nil {
			t.Fatal()
		}

		res, err := repo.DeleteById(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, true, res.Valid)
	})

	t.Run("fail run DeleteById", func(t *testing.T) {
		_, err := repo.DeleteById(10, 1)
		assert.NotNil(t, err)
	})

}

func TestGetAllMe(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&models.Product{})
	db.Migrator().CreateTable(&models.Product{})

	t.Run("success run GetAllMe", func(t *testing.T) {
		mockUser1 := models.User{Name: "anonim2", Email: "anonim2", Password: "anonim2"}
		if _, err := user.New(db).Create(mockUser1); err != nil {
			t.Fatal()
		}
		mockProT1 := models.ProductType{Name: "anonim1"}
		if _, err := prodType.New(db).Create(mockProT1); err != nil {
			t.Fatal()
		}
		mockProd1 := models.Product{Product_type_id: 1, Name: "anonim1 product", Price: 1000, Qty: 10, Description: "anonim1 Description"}
		if _, err := repo.Create(1, mockProd1); err != nil {
			t.Fatal()
		}

		_, err := repo.GetAllMe(1)
		assert.Nil(t, err)
	})

	t.Run("fail run GetAllMe", func(t *testing.T) {
		if _, err := repo.DeleteById(1, 1); err != nil {
			t.Log()
		}

		_, err := repo.GetAllMe(1)
		assert.NotNil(t, err)
	})
}

func TestGetByIdMe(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&models.Product{})
	db.Migrator().CreateTable(&models.Product{})

	t.Run("success run GetAllMe", func(t *testing.T) {
		mockUser1 := models.User{Name: "anonim2", Email: "anonim2", Password: "anonim2"}
		if _, err := user.New(db).Create(mockUser1); err != nil {
			t.Fatal()
		}
		mockProT1 := models.ProductType{Name: "anonim1"}
		if _, err := prodType.New(db).Create(mockProT1); err != nil {
			t.Fatal()
		}
		mockProd1 := models.Product{Product_type_id: 1, Name: "anonim1 product", Price: 1000, Qty: 10, Description: "anonim1 Description"}
		if _, err := repo.Create(1, mockProd1); err != nil {
			t.Fatal()
		}

		_, err := repo.GetByIdMe(1, 1)
		assert.Nil(t, err)
	})

	t.Run("fail run GetByIdMe", func(t *testing.T) {
		_, err := repo.GetByIdMe(10, 1)
		assert.NotNil(t, err)
	})
}

func TestGetAll(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&models.Product{})
	db.Migrator().CreateTable(&models.Product{})

	t.Run("success run GetAll", func(t *testing.T) {
		mockUser1 := models.User{Name: "anonim2", Email: "anonim2", Password: "anonim2"}
		if _, err := user.New(db).Create(mockUser1); err != nil {
			t.Fatal()
		}
		mockProT1 := models.ProductType{Name: "anonim1"}
		if _, err := prodType.New(db).Create(mockProT1); err != nil {
			t.Fatal()
		}
		mockProd1 := models.Product{Product_type_id: 1, Name: "anonim1 product", Price: 1000, Qty: 10, Description: "anonim1 Description"}
		if _, err := repo.Create(1, mockProd1); err != nil {
			t.Fatal()
		}

		_, err := repo.GetAll()
		assert.Nil(t, err)
	})

	t.Run("fail run GetAll", func(t *testing.T) {
		if _, err := repo.DeleteById(1, 1); err != nil {
			t.Log()
		}

		_, err := repo.GetAll()
		assert.NotNil(t, err)
	})
}

func TestGetById(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&models.Product{})
	db.Migrator().CreateTable(&models.Product{})

	t.Run("success run GetAll", func(t *testing.T) {
		mockUser1 := models.User{Name: "anonim2", Email: "anonim2", Password: "anonim2"}
		if _, err := user.New(db).Create(mockUser1); err != nil {
			t.Fatal()
		}
		mockProT1 := models.ProductType{Name: "anonim1"}
		if _, err := prodType.New(db).Create(mockProT1); err != nil {
			t.Fatal()
		}
		mockProd1 := models.Product{Product_type_id: 1, Name: "anonim1 product", Price: 1000, Qty: 10, Description: "anonim1 Description"}
		if _, err := repo.Create(1, mockProd1); err != nil {
			t.Fatal()
		}

		_, err := repo.GetById(1)
		assert.Nil(t, err)
	})

	t.Run("fail run GetAll", func(t *testing.T) {
		_, err := repo.GetById(10)
		assert.NotNil(t, err)
	})
}
