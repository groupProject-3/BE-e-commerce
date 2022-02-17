package user

import (
	"be/configs"
	"be/delivery/templates"
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
	db.AutoMigrate(&models.User{})

	t.Run("success run create", func(t *testing.T) {
		mockUser := models.User{Name: "anonim1", Email: "anonim1", Password: "anonim1"}
		res, err := repo.Create(mockUser)
		assert.Nil(t, err)
		assert.NotNil(t, res)
	})

	t.Run("fail run create", func(t *testing.T) {
		mockUser1 := models.User{Name: "anonim2", Email: "anonim2", Password: "anonim2"}
		if _, err := repo.Create(mockUser1); err != nil {
			t.Fatal()
		}
		mockUser := models.User{Name: "anonim2", Email: "anonim2", Password: "anonim2"}
		res, err := repo.Create(mockUser)
		assert.NotNil(t, err)
		assert.Equal(t, models.User{}, res)
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
	db.AutoMigrate(&models.User{})

	t.Run("success run GetById", func(t *testing.T) {
		mocUser := models.User{Name: "anonim123", Email: "anonim@1", Password: "anonim1"}

		if _, err := repo.Create(mocUser); err != nil {
			t.Fatal()
		}
		res, err := repo.GetById(1)
		assert.Nil(t, err)
		assert.Equal(t, 1, int(res.ID))
	})

	t.Run("fail run GetById", func(t *testing.T) {
		_, err := repo.GetById(10)
		assert.NotNil(t, err)
	})
}

func TestUpdateById(t *testing.T) {
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
	db.AutoMigrate(&models.User{})

	t.Run("success run UpdateById", func(t *testing.T) {
		mocUser := models.User{Name: "anonim123", Email: "anonim@1", Password: "anonim1"}

		if _, err := repo.Create(mocUser); err != nil {
			t.Fatal()
		}

		upUser := templates.UserRequest{Name: "anonim2", Email: "anonim2", Password: "anonim2"}
		res, err := repo.UpdateById(1, upUser)
		assert.Nil(t, err)
		assert.NotNil(t, res)
	})

	t.Run("fail run UpdateById", func(t *testing.T) {
		upUser := templates.UserRequest{Name: "anonim2", Email: "anonim2", Password: "anonim2"}
		_, err := repo.UpdateById(10, upUser)
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
	db.AutoMigrate(&models.User{})

	t.Run("success run DeleteById", func(t *testing.T) {
		mocUser := models.User{Name: "anonim123", Email: "anonim@1", Password: "anonim1"}

		if _, err := repo.Create(mocUser); err != nil {
			t.Fatal()
		}

		res, err := repo.DeleteById(1)
		assert.Nil(t, err)
		assert.NotNil(t, res)
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
	db.Migrator().DropTable(&models.PaymentMethod{})
	db.Migrator().DropTable(&models.User{})
	db.Migrator().DropTable(&models.Product{})
	db.Migrator().DropTable(&models.Cart{})
	db.Migrator().DropTable(&models.Order{})
	db.Migrator().DropTable(&models.OrderDetail{})
	db.AutoMigrate(&models.User{})

	t.Run("success run GetAll", func(t *testing.T) {
		mocUser := models.User{Name: "anonim123", Email: "anonim@1", Password: "anonim1"}

		if _, err := repo.Create(mocUser); err != nil {
			t.Fatal()
		}

		res, err := repo.GetAll()
		assert.Nil(t, err)
		assert.NotNil(t, res)
	})

	t.Run("fail run GetAll", func(t *testing.T) {
		if _, err := repo.DeleteById(1); err != nil {
			t.Log()
		}
		_, err := repo.GetAll()
		assert.NotNil(t, err)
	})
}
