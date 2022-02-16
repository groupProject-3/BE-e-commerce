package cart

import (
	"be/configs"
	"be/delivery/controllers/cart"
	"be/lib/database/prodType"
	"be/lib/database/product"
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

	db.Migrator().DropTable(&models.Cart{})
	db.Migrator().CreateTable(&models.Cart{})

	t.Run("success run create", func(t *testing.T) {
		mockUser1 := models.User{Name: "anonim1 user", Email: "anonim1 user", Password: "anonim1 user"}
		if _, err := user.New(db).Create(mockUser1); err != nil {
			t.Fatal()
		}
		mockProT1 := models.ProductType{Name: "anonim1 prot"}
		if _, err := prodType.New(db).Create(mockProT1); err != nil {
			t.Fatal()
		}
		mockProd1 := models.Product{Product_type_id: 1, Name: "anonim1 product", Price: 1000, Qty: 10, Description: "anonim1 Description"}
		if _, err := product.New(db).Create(1, mockProd1); err != nil {
			t.Fatal()
		}

		mockCart := models.Cart{Product_id: 1, Qty: 10}
		res, err := repo.Create(1, mockCart)
		assert.Nil(t, err)
		assert.NotNil(t, res)
	})

	t.Run("fail run create", func(t *testing.T) {
		mockCart := models.Cart{Model: gorm.Model{ID: 1}, Product_id: 10, Qty: 10}
		_, err := repo.Create(1, mockCart)
		assert.NotNil(t, err)
	})

}

func TestDeleteById(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	db.Migrator().DropTable(&models.Cart{})
	db.Migrator().CreateTable(&models.Cart{})

	t.Run("success run DeleteById", func(t *testing.T) {
		mockUser1 := models.User{Name: "anonim1 user", Email: "anonim1 user", Password: "anonim1 user"}
		if _, err := user.New(db).Create(mockUser1); err != nil {
			t.Fatal()
		}
		mockProT1 := models.ProductType{Name: "anonim1 prot"}
		if _, err := prodType.New(db).Create(mockProT1); err != nil {
			t.Fatal()
		}
		mockProd1 := models.Product{Product_type_id: 1, Name: "anonim1 product", Price: 1000, Qty: 10, Description: "anonim1 Description"}
		if _, err := product.New(db).Create(1, mockProd1); err != nil {
			t.Fatal()
		}
		mockCart1 := models.Cart{Product_id: 1, Qty: 10}
		if _, err := repo.Create(1, mockCart1); err != nil {
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

func TestUpdateById(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	db.Migrator().DropTable(&models.Cart{})
	db.Migrator().CreateTable(&models.Cart{})

	t.Run("success run UpdateById", func(t *testing.T) {
		mockUser1 := models.User{Name: "anonim1 user", Email: "anonim1 user", Password: "anonim1 user"}
		if _, err := user.New(db).Create(mockUser1); err != nil {
			t.Fatal()
		}
		mockProT1 := models.ProductType{Name: "anonim1 prot"}
		if _, err := prodType.New(db).Create(mockProT1); err != nil {
			t.Fatal()
		}
		mockProd1 := models.Product{Product_type_id: 1, Name: "anonim1 product", Price: 1000, Qty: 10, Description: "anonim1 Description"}
		if _, err := product.New(db).Create(1, mockProd1); err != nil {
			t.Fatal()
		}
		mockCart1 := models.Cart{Product_id: 1, Qty: 10}
		if _, err := repo.Create(1, mockCart1); err != nil {
			t.Fatal()
		}
		UpCart := cart.CartRequest{Product_id: 1, Qty: 20}
		// log.Info(repo.GetAll(1))
		res, err := repo.UpdateById(1, 1, UpCart)
		assert.Nil(t, err)
		// log.Info(repo.GetAll(1))
		assert.Equal(t, 20, int(res.Qty))
	})

	t.Run("fail run UpdateById", func(t *testing.T) {
		UpCart := cart.CartRequest{Product_id: 1, Qty: 20}
		_, err := repo.UpdateById(10, 1, UpCart)
		assert.NotNil(t, err)
	})
}

func TestGetAll(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	db.Migrator().DropTable(&models.Cart{})
	db.Migrator().CreateTable(&models.Cart{})

	t.Run("success run GetAll", func(t *testing.T) {
		mockUser1 := models.User{Name: "anonim1 user", Email: "anonim1 user", Password: "anonim1 user"}
		if _, err := user.New(db).Create(mockUser1); err != nil {
			t.Fatal()
		}
		mockProT1 := models.ProductType{Name: "anonim1 prot"}
		if _, err := prodType.New(db).Create(mockProT1); err != nil {
			t.Fatal()
		}
		mockProd1 := models.Product{Product_type_id: 1, Name: "anonim1 product", Price: 1000, Qty: 10, Description: "anonim1 Description"}
		if _, err := product.New(db).Create(1, mockProd1); err != nil {
			t.Fatal()
		}
		mockCart1 := models.Cart{Product_id: 1, Qty: 10}
		if _, err := repo.Create(1, mockCart1); err != nil {
			t.Fatal()
		}

		_, err := repo.GetAll(1)
		assert.Nil(t, err)
	})

	t.Run("fail run GetAll", func(t *testing.T) {
		if _, err := repo.DeleteById(1, 1); err != nil {
			t.Log()
		}
		_, err := repo.GetAll(1)
		assert.NotNil(t, err)
	})
}
