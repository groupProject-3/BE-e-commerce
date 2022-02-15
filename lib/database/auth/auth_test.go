package auth

import (
	"be/configs"
	"be/delivery/controllers/auth"
	libUser "be/lib/database/user"
	"be/models"
	"be/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&models.Cart{}, &models.Order{}, &models.OrderDetail{}, &models.PaymentMethod{}, &models.Product{}, &models.ProductType{}, &models.User{})
	db.AutoMigrate(&models.User{})

	t.Run("success run login", func(t *testing.T) {
		mockUser := models.User{Name: "anonim123", Email: "anonim@123", Password: "anonim123"}
		_, err := libUser.New(db).Create(mockUser)
		if err != nil {
			t.Fail()
		}
		mockLogin := auth.Userlogin{Email: "anonim@123", Password: "anonim123"}
		res, err := repo.Login(mockLogin)
		assert.Nil(t, err)
		assert.Equal(t, "anonim@123", res.Email)
		assert.Equal(t, "anonim123", res.Password)
	})

	t.Run("fail run login", func(t *testing.T) {
		mockLogin := auth.Userlogin{Email: "anonim@456", Password: "anonim456"}
		_, err := repo.Login(mockLogin)
		assert.NotNil(t, err)
	})

}
