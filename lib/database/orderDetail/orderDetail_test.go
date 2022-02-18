package orderDetail

// import (
// 	"be/configs"
// 	"be/lib/database/cart"
// 	"be/lib/database/order"
// 	"be/lib/database/paymethod"
// 	"be/lib/database/prodType"
// 	"be/lib/database/product"
// 	"be/lib/database/user"
// 	"be/models"
// 	"be/utils"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestCreate(t *testing.T) {
// 	config := configs.GetConfig()
// 	db := utils.InitDB(config)
// 	repo := New(db)
// 	db.Migrator().DropTable(&models.ProductType{})
// 	db.Migrator().DropTable(&models.PaymentMethod{})
// 	db.Migrator().DropTable(&models.User{})
// 	db.Migrator().DropTable(&models.Product{})
// 	db.Migrator().DropTable(&models.Cart{})
// 	db.Migrator().DropTable(&models.Order{})
// 	db.Migrator().DropTable(&models.OrderDetail{})
// 	db.AutoMigrate(&models.OrderDetail{})

// 	t.Run("success run create", func(t *testing.T) {
// 		mockUser1 := models.User{Name: "anonim1 user", Email: "anonim1 user", Password: "anonim1 user"}
// 		if _, err := user.New(db).Create(mockUser1); err != nil {
// 			t.Fatal()
// 		}
// 		mockProT1 := models.ProductType{Name: "anonim1 prot"}
// 		if _, err := prodType.New(db).Create(mockProT1); err != nil {
// 			t.Fatal()
// 		}
// 		mockProd1 := models.Product{Product_type_id: 1, Name: "anonim1 product", Price: 1000, Qty: 10, Description: "anonim1 Description"}
// 		if _, err := product.New(db).Create(1, mockProd1); err != nil {
// 			t.Fatal()
// 		}
// 		mockCart1 := models.Cart{Product_id: 1, Qty: 10}
// 		if _, err := cart.New(db).Create(1, mockCart1); err != nil {
// 			t.Fatal()
// 		}
// 		mockPm1 := models.PaymentMethod{Name: "anonim1"}
// 		if _, err := paymethod.New(db).Create(mockPm1); err != nil {
// 			t.Fatal()
// 		}
// 		mockOrder1 := models.Order{Payment_method_id: 1}
// 		if _, err := order.New(db).Create(1, mockOrder1); err != nil {
// 			t.Fatal()
// 		}

// 		mockOrderDet1 := models.OrderDetail{Cart_id: 1, Order_id: 1, Qty: mockCart1.Qty, Price: 10000}
// 		res, err := repo.Create(mockOrderDet1)
// 		assert.Nil(t, err)
// 		assert.NotNil(t, res)
// 	})

// 	t.Run("fail run create", func(t *testing.T) {
// 		mockOrderDet1 := models.OrderDetail{Cart_id: 10, Order_id: 10, Qty: 10, Price: 10000}
// 		_, err := repo.Create(mockOrderDet1)
// 		assert.NotNil(t, err)
// 	})
// }

// func TestDeleteById(t *testing.T) {
// 	config := configs.GetConfig()
// 	db := utils.InitDB(config)
// 	repo := New(db)
// 	db.Migrator().DropTable(&models.ProductType{})
// 	db.Migrator().DropTable(&models.PaymentMethod{})
// 	db.Migrator().DropTable(&models.User{})
// 	db.Migrator().DropTable(&models.Product{})
// 	db.Migrator().DropTable(&models.Cart{})
// 	db.Migrator().DropTable(&models.Order{})
// 	db.Migrator().DropTable(&models.OrderDetail{})
// 	db.AutoMigrate(&models.OrderDetail{})

// 	t.Run("success run DeleteById", func(t *testing.T) {
// 		mockUser1 := models.User{Name: "anonim1 user", Email: "anonim1 user", Password: "anonim1 user"}
// 		if _, err := user.New(db).Create(mockUser1); err != nil {
// 			t.Fatal()
// 		}
// 		mockProT1 := models.ProductType{Name: "anonim1 prot"}
// 		if _, err := prodType.New(db).Create(mockProT1); err != nil {
// 			t.Fatal()
// 		}
// 		mockProd1 := models.Product{Product_type_id: 1, Name: "anonim1 product", Price: 1000, Qty: 10, Description: "anonim1 Description"}
// 		if _, err := product.New(db).Create(1, mockProd1); err != nil {
// 			t.Fatal()
// 		}
// 		mockCart1 := models.Cart{Product_id: 1, Qty: 10}
// 		if _, err := cart.New(db).Create(1, mockCart1); err != nil {
// 			t.Fatal()
// 		}
// 		mockPm1 := models.PaymentMethod{Name: "anonim1"}
// 		if _, err := paymethod.New(db).Create(mockPm1); err != nil {
// 			t.Fatal()
// 		}
// 		mockOrder1 := models.Order{Payment_method_id: 1}
// 		if _, err := order.New(db).Create(1, mockOrder1); err != nil {
// 			t.Fatal()
// 		}

// 		mockOrderDet1 := models.OrderDetail{Cart_id: 1, Order_id: 1, Qty: mockCart1.Qty, Price: 10000}
// 		if _, err := repo.Create(mockOrderDet1); err != nil {
// 			t.Fatal()
// 		}
// 		res, err := repo.DeleteById(1)
// 		assert.Nil(t, err)
// 		assert.Equal(t, true, res.Valid)
// 	})

// 	t.Run("fail run DeleteById", func(t *testing.T) {
// 		_, err := repo.DeleteById(10)
// 		assert.NotNil(t, err)
// 	})
// }

// func TestGetAll(t *testing.T) {
// 	config := configs.GetConfig()
// 	db := utils.InitDB(config)
// 	repo := New(db)
// 	db.Migrator().DropTable(&models.ProductType{})
// 	db.Migrator().DropTable(&models.PaymentMethod{})
// 	db.Migrator().DropTable(&models.User{})
// 	db.Migrator().DropTable(&models.Product{})
// 	db.Migrator().DropTable(&models.Cart{})
// 	db.Migrator().DropTable(&models.Order{})
// 	db.Migrator().DropTable(&models.OrderDetail{})
// 	db.AutoMigrate(&models.OrderDetail{})

// 	t.Run("success run DeleteById", func(t *testing.T) {
// 		mockUser1 := models.User{Name: "anonim1 user", Email: "anonim1 user", Password: "anonim1 user"}
// 		if _, err := user.New(db).Create(mockUser1); err != nil {
// 			t.Fatal()
// 		}
// 		mockProT1 := models.ProductType{Name: "anonim1 prot"}
// 		if _, err := prodType.New(db).Create(mockProT1); err != nil {
// 			t.Fatal()
// 		}
// 		mockProd1 := models.Product{Product_type_id: 1, Name: "anonim1 product", Price: 1000, Qty: 10, Description: "anonim1 Description"}
// 		if _, err := product.New(db).Create(1, mockProd1); err != nil {
// 			t.Fatal()
// 		}
// 		mockCart1 := models.Cart{Product_id: 1, Qty: 10}
// 		if _, err := cart.New(db).Create(1, mockCart1); err != nil {
// 			t.Fatal()
// 		}
// 		mockPm1 := models.PaymentMethod{Name: "anonim1"}
// 		if _, err := paymethod.New(db).Create(mockPm1); err != nil {
// 			t.Fatal()
// 		}
// 		mockOrder1 := models.Order{Payment_method_id: 1}
// 		if _, err := order.New(db).Create(1, mockOrder1); err != nil {
// 			t.Fatal()
// 		}

// 		mockOrderDet1 := models.OrderDetail{Cart_id: 1, Order_id: 1, Qty: mockCart1.Qty, Price: 10000}
// 		if _, err := repo.Create(mockOrderDet1); err != nil {
// 			t.Fatal()
// 		}


// 		res , err := repo.GetAll(1)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("fail run GetAll", func(t *testing.T) {
// 		if _, err := repo.DeleteById(1); err != nil {
// 			t.Log()
// 		}

// 		_, err := repo.GetAll(1)
// 		assert.NotNil(t, err)

// 	})
// }
