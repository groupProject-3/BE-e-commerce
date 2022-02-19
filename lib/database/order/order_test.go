package order

import (
	"be/configs"
	"be/delivery/templates"
	"be/lib/database/cart"
	paymentmethod "be/lib/database/paymethod"
	"be/lib/database/prodType"
	"be/lib/database/product"
	"be/lib/database/user"
	"be/models"
	"be/utils"
	"testing"

	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&models.ProductType{})
	db.Migrator().DropTable(&models.PaymentMethod{})
	db.Migrator().DropTable(&models.User{})
	db.Migrator().DropTable(&models.Product{})
	db.Migrator().DropTable(&models.Cart{})
	db.Migrator().DropTable(&models.Order{})
	db.Migrator().DropTable(&models.OrderDetail{})
	db.AutoMigrate(&models.Order{})

	t.Run("success run create", func(t *testing.T) {
		mockUser1 := models.User{Name: "anonim1 user", Email: "anonim1 user", Password: "anonim1 user"}
		if _, err := user.New(db).Create(mockUser1); err != nil {
			t.Fatal()
		}
		mockPm1 := models.PaymentMethod{Name: "anonim1"}
		if _, err := paymentmethod.New(db).Create(mockPm1); err != nil {
			t.Fatal()
		}

		mockOrder1 := models.Order{Payment_method_id: 1}
		res, err := repo.Create(1, mockOrder1)
		assert.Nil(t, err)
		assert.NotNil(t, res)
	})

	t.Run("fail run create", func(t *testing.T) {
		mockOrder1 := models.Order{Payment_method_id: 10}
		_, err := repo.Create(10, mockOrder1)
		assert.NotNil(t, err)
	})
}

func TestDeleteById(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&models.ProductType{})
	db.Migrator().DropTable(&models.PaymentMethod{})
	db.Migrator().DropTable(&models.User{})
	db.Migrator().DropTable(&models.Product{})
	db.Migrator().DropTable(&models.Cart{})
	db.Migrator().DropTable(&models.Order{})
	db.Migrator().DropTable(&models.OrderDetail{})
	db.AutoMigrate(&models.Order{})

	t.Run("success run create", func(t *testing.T) {
		mockUser1 := models.User{Name: "anonim1 user", Email: "anonim1 user", Password: "anonim1 user"}
		if _, err := user.New(db).Create(mockUser1); err != nil {
			t.Fatal()
		}
		mockPm1 := models.PaymentMethod{Name: "anonim1"}
		if _, err := paymentmethod.New(db).Create(mockPm1); err != nil {
			t.Fatal()
		}

		mockOrder1 := models.Order{Payment_method_id: 1}
		if _, err := repo.Create(1, mockOrder1); err != nil {
			t.Fatal()
		}
		res, err := repo.DeleteById(1)
		assert.Nil(t, err)
		assert.Equal(t, true, res.Valid)
	})

	t.Run("fail run create", func(t *testing.T) {
		_, err := repo.DeleteById(1)
		assert.NotNil(t, err)
	})
}

func TestGetById(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&models.ProductType{})
	db.Migrator().DropTable(&models.PaymentMethod{})
	db.Migrator().DropTable(&models.User{})
	db.Migrator().DropTable(&models.Product{})
	db.Migrator().DropTable(&models.Cart{})
	db.Migrator().DropTable(&models.Order{})
	db.Migrator().DropTable(&models.OrderDetail{})
	db.AutoMigrate(&models.Order{})
	db.AutoMigrate(&models.Cart{})

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
		mockCart1 := models.Cart{Product_id: 1, Qty: 10, Status: "order"}
		if _, err := cart.New(db).Create(1, mockCart1); err != nil {
			t.Fatal()
		}
		mockPm1 := models.PaymentMethod{Name: "anonim1"}
		if _, err := paymentmethod.New(db).Create(mockPm1); err != nil {
			t.Fatal()
		}

		mockOrder1 := models.Order{Payment_method_id: 1}
		if _, err := repo.Create(1, mockOrder1); err != nil {
			t.Fatal()
		}
		res, err := repo.GetById(1, 1)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		log.Info(res)

	})

	t.Run("fail run create", func(t *testing.T) {
		_, err := repo.GetById(10, 10)
		assert.NotNil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&models.ProductType{})
	db.Migrator().DropTable(&models.PaymentMethod{})
	db.Migrator().DropTable(&models.User{})
	db.Migrator().DropTable(&models.Product{})
	db.Migrator().DropTable(&models.Cart{})
	db.Migrator().DropTable(&models.Order{})
	db.Migrator().DropTable(&models.OrderDetail{})
	db.AutoMigrate(&models.Order{})

	t.Run("success run create", func(t *testing.T) {
		mockUser1 := models.User{Name: "anonim1 user", Email: "anonim1 user", Password: "anonim1 user"}
		if _, err := user.New(db).Create(mockUser1); err != nil {
			t.Fatal()
		}
		mockPm1 := models.PaymentMethod{Name: "anonim1"}
		if _, err := paymentmethod.New(db).Create(mockPm1); err != nil {
			t.Fatal()
		}

		mockOrder1 := models.Order{Payment_method_id: 1}
		if _, err := repo.Create(1, mockOrder1); err != nil {
			t.Fatal()
		}

		upPro := templates.OrderRequest{Status: "payed"}

		res, err := repo.Update(1, upPro)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		log.Info(res)
	})

	// t.Run("fail run create", func(t *testing.T) {
	// 	mockOrder1 := models.Order{Payment_method_id: 10}
	// 	_, err := repo.Create(10, mockOrder1)
	// 	assert.NotNil(t, err)
	// })
}
