// This file is responsible for creating and managing the user table in the database.
//
// On initial startup, this fill will also initialize the part of the database which it requires.
// This isn't exactly idea but it will have to work for now.

package db

import (
	"log"

	"github.com/help-me-someone/scalable-p2-db/models/user"
	"github.com/help-me-someone/scalable-p2-db/models/video"
	"gorm.io/gorm"
)

// InitTables creates the user and video table in the database.
func InitTables(db *gorm.DB) {
	if db == nil {
		log.Panic("Database is invalid.")
	}

	// Migrate the models.
	err := db.AutoMigrate(
		&user.User{},
		&video.Video{},
		&video.VideoLikes{},
		&video.VideoComments{})
	if err != nil {
		log.Panic("Failed to migrate User table.")
	}
}
