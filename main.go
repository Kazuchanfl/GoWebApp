package main

import (
	"github.com/Kazuchanfl/GoWebApp/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

var db = make(map[string]string)

func main() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Delete - delete product
	db.Delete(&Product{}, 1)

	// gin web server を起動
	config.SetupRouter()

}
