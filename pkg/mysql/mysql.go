package mysql

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Database() {

	DB_PASSWORD := os.Getenv("MYSQLPASSWORD")
	DB_USER := os.Getenv("MYSQLUSER")
	DB_HOST := os.Getenv("MYSQLHOST")
	DB_PORT := os.Getenv("MYSQLPORT")
	DB_NAME := os.Getenv("MYSQLDATABASE")


	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("mysql error")
	}

	fmt.Println("Connected to Database")
}
