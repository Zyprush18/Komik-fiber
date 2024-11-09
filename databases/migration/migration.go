package migration

import (
	"fmt"

	"github.com/Zyprush18/Komik-fiber/databases"
	"github.com/Zyprush18/Komik-fiber/models/entity"
)

func RunMigrate() {
	err := databases.DB.AutoMigrate(&entity.User{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Succes to Migrate")
}
