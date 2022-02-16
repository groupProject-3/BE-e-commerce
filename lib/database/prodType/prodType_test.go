package prodType

import (
	"be/configs"
	"be/delivery/templates"
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
	db.Migrator().DropTable(&models.ProductType{})
	db.Migrator().CreateTable(&models.ProductType{})

	t.Run("success run create", func(t *testing.T) {
		mockPro := models.ProductType{Name: "anonim1"}
		res, err := repo.Create(mockPro)
		assert.Nil(t, err)
		assert.NotNil(t, res)
	})

	t.Run("fail run create", func(t *testing.T) {
		mockPro1 := models.ProductType{Name: "anonim1"}
		if _, err := repo.Create(mockPro1); err != nil {
			t.Fatal()
		}
		mockPro2 := models.ProductType{Model: gorm.Model{ID: 1}, Name: "anonim2"}
		_, err := repo.Create(mockPro2)
		assert.NotNil(t, err)
	})
}

func TestUpdateById(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&models.ProductType{})
	db.Migrator().CreateTable(&models.ProductType{})

	t.Run("success run UpdateById", func(t *testing.T) {
		mockPro1 := models.ProductType{Name: "anonim1"}
		if _, err := repo.Create(mockPro1); err != nil {
			t.Fatal()
		}
		upPro := templates.ProductTypeRequest{Name: "anonim2"}
		res, err := repo.UpdateById(1, upPro)
		assert.Nil(t, err)
		assert.Equal(t, "anonim2", res.Name)
	})

	t.Run("fail run UpdateById", func(t *testing.T) {
		upPro := templates.ProductTypeRequest{Name: "anonim2"}
		_, err := repo.UpdateById(10, upPro)
		assert.NotNil(t, err)
	})
}

func TestDeleteById(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&models.ProductType{})
	db.Migrator().CreateTable(&models.ProductType{})

	t.Run("success run DeleteById", func(t *testing.T) {
		mockPro1 := models.ProductType{Name: "anonim1"}
		if _, err := repo.Create(mockPro1); err != nil {
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
	db.Migrator().DropTable(&models.ProductType{})
	db.Migrator().CreateTable(&models.ProductType{})

	t.Run("success run GetAll", func(t *testing.T) {
		mockPro1 := models.ProductType{Name: "anonim1"}
		if _, err := repo.Create(mockPro1); err != nil {
			t.Fatal()
		}

		_, err := repo.GetAll()
		assert.Nil(t, err)
	})

	t.Run("fail run GetAll", func(t *testing.T) {
		if _, err := repo.DeleteById(1); err != nil {
			t.Log()
		}
		_, err := repo.GetAll()
		assert.NotNil(t, err)
	})
}
