// This file is responsible for creating and managing the user table in the database.
//
// On initial startup, this fill will also initialize the part of the database which it requires.
// This isn't exactly idea but it will have to work for now.

package db

import (
	"log"

	"github.com/help-me-someone/scalable-p2-db/models/user"
	"gorm.io/gorm"
)

// InitUserTable creates the user table in the database. This function will check for the existence
// of the table first, only if the table DOES NOT EXIST, it will create it, otherwise it will do nothing.
func InitUserTable(db *gorm.DB) {
	if db == nil {
		log.Panic("Database is invalid.")
	}

	// Migrate the schema
	err := db.AutoMigrate(&user.User{})
	if err != nil {
		log.Panic("Failed to migrate User table.")
	}
}
