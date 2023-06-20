package initializers

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Global DB
var DB *gorm.DB

func ConnectToDB() {
	dsn := "root:2113@tcp(127.0.0.1:3306)/test_api_aquafarm?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed Connect to DATABSE")
	}
}
