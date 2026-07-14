package config

import (
	"os"

	"github.com/blackbarth/goopportunities.git/schemas"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	logger.Info("Initializing SQLite database...")
	
	//Check if the database file exists, if not create it
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		logger.Info("Database file not found, creating a new one...")
		err := os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Errorf("Failed to create database directory: %v", err)
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("Failed to create database file: %v", err)
			return nil, err
		}
		file.Close()
		logger.Info("Database file created successfully")
	} else {
		logger.Info("Database file already exists")
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Failed to initialize SQLite database: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("Failed to auto-migrate database schema: %v", err)
		return nil, err
	}
	logger.Info("SQLite database initialized successfully")
	return db, nil
}
