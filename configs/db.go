package configs

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open("postgresql://postgres:BmJSabBG9RBYQV9fWNh9@containers-us-west-163.railway.app:5454/railway"), &gorm.Config{})
	fmt.Println("Connection")
	if err != nil {
		fmt.Println("error", err)
		panic(err)
	}
	return db, err
}
