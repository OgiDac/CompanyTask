package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGormConnection(env *Env) *gorm.DB {
	// DSN without database name
	// baseDSN := "root:1234@tcp(localhost:3306)/"
	baseDSN := env.BaseDSN
	targetDB := env.TargetDB

	// Connect to MySQL server (no DB yet)
	serverDB, err := gorm.Open(mysql.Open(baseDSN), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to MySQL server:", err)
		return nil
	}

	// Create database if not exists
	err = serverDB.Exec("CREATE DATABASE IF NOT EXISTS " + targetDB).Error
	if err != nil {
		fmt.Println("Error creating database:", err)
		return nil
	}

	// Now connect to the target database
	// finalDSN := "root:1234@tcp(localhost:3306)/company?parseTime=true"
	finalDSN := env.BaseDSN + env.TargetDB
	db, err := gorm.Open(mysql.Open(finalDSN), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return nil
	}

	fmt.Println("Database connected")
	return db
}

func CloseGormConnection(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("Error getting DB from Gorm:", err)
		return
	}

	err = sqlDB.Close()
	if err != nil {
		fmt.Println("Error closing the database:", err)
		return
	}

	fmt.Println("Database connection closed")
}
