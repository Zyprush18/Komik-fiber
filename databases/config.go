package databases

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB()  {
	var err error
	const mysqls = "root:root@tcp(127.0.0.1:3306)/db_crud_fiber?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := mysqls

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Can't connect to database")
	}
	fmt.Println("Succes Connected to Database")
}