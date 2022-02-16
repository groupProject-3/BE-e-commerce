package orderDetail

import (
	"be/configs"
	"be/lib/database/cart"
	"be/lib/database/order"
	"be/lib/database/paymentmethod"
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
	db.Migrator().DropTable(&models.OrderDetail{})
	db.Migrator().CreateTable(&models.OrderDetail{})

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
		mockCart1 := models.Cart{Product_id: 1, Qty: 10}
		if _, err := cart.New(db).Create(1, mockCart1); err != nil {
			t.Fatal()
		}
		mockPm1 := models.PaymentMethod{Name: "anonim1"}
		if _, err := paymentmethod.New(db).Create(mockPm1); err != nil {
			t.Fatal()
		}
		mockOrder1 := models.Order{Payment_method_id: 1}
		if _, err := order.New(db).Create(1, mockOrder1); err != nil {
			t.Fatal()
		}

		mockOrderDet1 := models.OrderDetail{Cart_id: 1, Order_id: 1, Qty: mockCart1.Qty, Price: 10000}
		res, err := repo.Create(mockOrderDet1)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		log.Info(res)
	})

	t.Run("fail run create", func(t *testing.T) {
		mockOrderDet1 := models.OrderDetail{Cart_id: 10, Order_id: 10, Qty: 10, Price: 10000}
		_, err := repo.Create(mockOrderDet1)
		assert.NotNil(t, err)
		log.Info(err)
	})
}
