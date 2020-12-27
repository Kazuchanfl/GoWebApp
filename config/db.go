package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

/*
Product はDBの動作確認のためのモック構造体
*/
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

var db = make(map[string]string)

/*
SetupDb でDBの初期化
*/
func SetupDb() {
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
}
