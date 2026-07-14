package schemas

import (
  "gorm.io/driver/sqlite"
  "gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role string
	Company string
	Location string
	Remote bool
	Link string
	Salary int64
}

func CreateDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Opening{})

	return db
}

