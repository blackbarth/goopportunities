package schemas

import (
	"time"

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



type OpeningResponse struct {
	Id uint `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	Role string `json:"role"`
	Company string `json:"company"`
	Location string `json:"location"`
	Remote bool `json:"remote"`
	Link string `json:"link"`
	Salary int64 `json:"salary"`
}