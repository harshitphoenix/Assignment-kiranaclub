package configs

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open("DB_URL"), &gorm.Config{})
	fmt.Println("Connection")
	if err != nil {
		fmt.Println("error", err)
		panic(err)
	}
	return db, err
}
