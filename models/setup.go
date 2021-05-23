package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase(db_name string) {
	// Open DB
	db, err := gorm.Open(sqlite.Open(db_name), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	// Auto Migrate
	db.AutoMigrate(&Article{})

	DB = db
}
