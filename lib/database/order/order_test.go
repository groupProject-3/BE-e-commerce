package order

import (
	"be/configs"
	paymentmethod "be/lib/database/paymentMethod"
	"be/lib/database/user"
	"be/models"
	"be/utils"
	"testing"

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
		res, err := repo.DeleteById(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, true, res.Valid)
	})

	t.Run("fail run create", func(t *testing.T) {
		_, err := repo.DeleteById(1, 1)
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
	db.AutoMigrate(&models.OrderDetail{})

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
		res, err := repo.GetById(1, 1)
		assert.Nil(t, err)
		assert.NotNil(t, res)
	})

	t.Run("fail run create", func(t *testing.T) {
		_, err := repo.GetById(10, 10)
		assert.NotNil(t, err)
	})
}
