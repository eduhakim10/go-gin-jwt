package testpackage

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDB() {
	database, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres password=Wafe2024 dbname=evendo port=5432 sslmode=disable TimeZone=Asia/Jakarta",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {

		panic(err)

	} else {
		fmt.Println("lah kok konek")
	}
	//fmt.Println("connect coy wkwwk")
	// }
	DB = database
}
