package cart

import (
	"be/configs"
	"be/delivery/templates"
	"be/lib/database/prodType"
	"be/lib/database/product"
	"be/lib/database/user"
	"be/models"
	"be/utils"
	"testing"

	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
)

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
	db.AutoMigrate(&models.Cart{})

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
		mockProd2 := models.Product{Product_type_id: 1, Name: "anonim1 product", Price: 1000, Qty: 10, Description: "anonim1 Description"}
		if _, err := product.New(db).Create(1, mockProd2); err != nil {
			t.Fatal()
		}
		mockCart1 := models.Cart{Product_id: 1, Qty: 5 /*  Status: "cart" */}
		if _, err := repo.CreateNew(1, mockCart1); err != nil {
			t.Fatal()
		}
		mockCart2 := models.Cart{Product_id: 2, Status: "payed"}
		if _, err := repo.CreateNew(1, mockCart2); err != nil {
			log.Info(err)
			t.Fatal()
		}

		res, err := repo.GetAll(1, "payed")
		assert.Nil(t, err)
		log.Info(res)
	})

	// t.Run("fail run GetAll", func(t *testing.T) {
	// 	if _, err := repo.DeleteNew(1, 1); err != nil {
	// 		t.Log()
	// 	}
	// 	_, err := repo.GetAll(1, "cart")
	// 	assert.NotNil(t, err)
	// })
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
	db.AutoMigrate(&models.Cart{})

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
		mockCart1 := models.Cart{Product_id: 1, Qty: 5}
		if _, err := repo.CreateNew(1, mockCart1); err != nil {
			t.Fatal()
		}

		res, err := repo.GetById(1, 1, "cart")
		assert.Nil(t, err)
		log.Info(res)
	})

	// t.Run("fail run GetAll", func(t *testing.T) {
	// 	if _, err := repo.DeleteById(1, 1); err != nil {
	// 		t.Log()
	// 	}
	// 	_, err := repo.GetAll(1)
	// 	assert.NotNil(t, err)
	// })
}

func TestCreateNew(t *testing.T) {
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
		mockProd1 := models.Product{Product_type_id: 1, Name: "anonim1 product", Price: 1000, Qty: 20, Description: "anonim1 Description"}
		if _, err := product.New(db).Create(1, mockProd1); err != nil {
			t.Fatal()
		}

		mockCart := models.Cart{Product_id: 1, Qty: 20}
		res, err := repo.CreateNew(1, mockCart)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		log.Info(res.Qty)
		log.Info(res.Product_qty)
	})

	// t.Run("success run create update", func(t *testing.T) {

	// 	mockCart := models.Cart{Product_id: 1, Qty: 25}
	// 	res, err := repo.CreateNew(1, mockCart)
	// 	assert.Nil(t, err)
	// 	log.Info(res.Qty)
	// 	log.Info(res.Product_qty)
	// })

}

func TestDeleteNew(t *testing.T) {
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
	db.AutoMigrate(&models.Cart{})

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
		mockCart1 := models.Cart{Product_id: 1, Qty: 5, Status: "cart"}
		if _, err := repo.CreateNew(1, mockCart1); err != nil {
			t.Fatal()
		}

		res, err := repo.DeleteNew(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, true, res.DeletedAt.Valid)
	})

	t.Run("fail run DeleteById", func(t *testing.T) {
		_, err := repo.DeleteNew(10, 1)
		assert.NotNil(t, err)
	})
}

func TestUpdateCart(t *testing.T) {
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
		mockCart1 := models.Cart{Product_id: 1, Qty: 6, Status: "order"}
		if _, err := repo.CreateNew(1, mockCart1); err != nil {
			t.Fatal()
		}

		mockCart := templates.CartRequest{Qty: 10}
		res, err := repo.UpdateCart(1, 1, mockCart)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		log.Info(err)
		log.Info(res.Status)
		log.Info(res.Qty)
		log.Info(res.Product_qty)
	})
}
