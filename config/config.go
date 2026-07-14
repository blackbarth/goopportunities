package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB

	logger *Logger

)

func Init() error {

	// Initialize the SQLite database
	var err error
	db, err = InitializeSQLite()
	if err != nil {
		return fmt.Errorf("failed to initialize SQLite database: %v", err)
	}

	return nil
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	sqlDB, err := db.DB()
	if err != nil {
		logger.Error("Failed to get database instance:", err)
		return
	}
	sqlDB.Close()
}

func GetLogger(p string ) *Logger {
	// Initialize the logger if it hasn't been created yet
	logger = NewLogger(p)
	return logger
}