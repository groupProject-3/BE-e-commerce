package paymentmethod

import (
	"be/configs"
	paymentmethod "be/delivery/controllers/paymentMethod"
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
	db.Migrator().DropTable(&models.PaymentMethod{})
	db.AutoMigrate(&models.PaymentMethod{})

	t.Run("success run create", func(t *testing.T) {
		mockPm := models.PaymentMethod{Name: "anonim1"}
		res, err := repo.Create(mockPm)
		assert.Nil(t, err)
		assert.NotNil(t, res)
	})

	t.Run("fail run create", func(t *testing.T) {
		mockPm1 := models.PaymentMethod{Name: "anonim1"}
		if _, err := repo.Create(mockPm1); err != nil {
			t.Fatal()
		}
		mockPm2 := models.PaymentMethod{Model: gorm.Model{ID: 1}, Name: "anonim2"}
		_, err := repo.Create(mockPm2)
		assert.NotNil(t, err)
	})
}

func TestUpdateById(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&models.PaymentMethod{})
	db.AutoMigrate(&models.PaymentMethod{})

	t.Run("success run UpdateById", func(t *testing.T) {
		mockPm1 := models.PaymentMethod{Name: "anonim1"}
		if _, err := repo.Create(mockPm1); err != nil {
			t.Fatal()
		}
		upPm := paymentmethod.PaymentMethodRequest{Name: "anonim2"}
		res, err := repo.UpdateById(1, upPm)
		assert.Nil(t, err)
		assert.Equal(t, "anonim2", res.Name)
	})

	t.Run("fail run UpdateById", func(t *testing.T) {
		upPm := paymentmethod.PaymentMethodRequest{Name: "anonim2"}
		_, err := repo.UpdateById(10, upPm)
		assert.NotNil(t, err)
	})
}

func TestDeleteById(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&models.PaymentMethod{})
	db.AutoMigrate(&models.PaymentMethod{})

	t.Run("success run DeleteById", func(t *testing.T) {
		mockPm1 := models.PaymentMethod{Name: "anonim1"}
		if _, err := repo.Create(mockPm1); err != nil {
			t.Fatal()
		}

		res, err := repo.DeleteById(1)
		assert.Nil(t, err)
		assert.Equal(t, true, res.Valid)
	})

	t.Run("fail run DeleteById", func(t *testing.T) {
		_, err := repo.DeleteById(10)
		assert.NotNil(t, err)
	})
}

func TestGetAll(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&models.PaymentMethod{})
	db.AutoMigrate(&models.PaymentMethod{})

	t.Run("success run GetAll", func(t *testing.T) {
		mockPm1 := models.PaymentMethod{Name: "anonim1"}
		if _, err := repo.Create(mockPm1); err != nil {
			t.Fatal()
		}

		_, err := repo.GetAll()
		assert.Nil(t, err)
	})

	t.Run("fail run GetAll", func(t *testing.T) {
		if _, err := repo.DeleteById(1); err != nil {
			t.Fatal()
		}
		_, err := repo.GetAll()
		assert.NotNil(t, err)
	})
}
