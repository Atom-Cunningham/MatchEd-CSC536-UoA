package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// TODO initialize gorm with postgres using real dbname
func InitDB() {
	var err error
	// Skip DB connection if it's not available yet
	if isDatabaseAvailable() {
		_, err = gorm.Open("postgres", "user=username dbname=mydb sslmode=disable")
		if err != nil {
			panic("failed to connect to database")
		}
		fmt.Println("Database connected")
	} else {
		fmt.Println("Database is not available, skipping initialization")
	}
}

// TODO ping the database
func isDatabaseAvailable() bool {
	// Check if the database is accessible, e.g., via pinging or a network check
	return false // Simulate a condition where the DB is not yet available
}
