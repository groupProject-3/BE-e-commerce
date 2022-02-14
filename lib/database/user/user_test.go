package user

import (
	"be/configs"
	"be/models"
	"be/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&models.User{})
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

