package utils

import (
	"be/configs"
	"be/models"
	"fmt"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config *configs.AppConfig) *gorm.DB {

	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Address,
		config.DB_Port,
		config.Name,
	)
	fmt.Println(connectionString)
	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Info("error in connect database ", err)
		panic(err)
	}

	AutoMigrate(DB)
	return DB
}

func AutoMigrate(DB *gorm.DB) {
	DB.AutoMigrate(&models.ProductType{})
	DB.AutoMigrate(&models.PaymentMethod{})
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Product{})
	DB.AutoMigrate(&models.Cart{})
	DB.AutoMigrate(&models.Order{})
	DB.AutoMigrate(&models.OrderDetail{})
}
