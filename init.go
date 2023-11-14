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
		return
	}

	var err error = nil

	if !db.Migrator().HasTable(&user.User{}) {
		dberr := db.AutoMigrate(&user.User{})
		if err == nil {
			err = dberr
		}
	}

	if !db.Migrator().HasTable(&video.Video{}) {
		dberr := db.AutoMigrate(&video.Video{})
		if err == nil {
			err = dberr
		}
	}

	if !db.Migrator().HasTable(&video.VideoLikes{}) {
		dberr := db.AutoMigrate(&video.VideoLikes{})
		if err == nil {
			err = dberr
		}
	}

	if !db.Migrator().HasTable(&video.VideoComments{}) {
		dberr := db.AutoMigrate(&video.VideoComments{})
		if err == nil {
			err = dberr
		}
	}

	if !db.Migrator().HasTable(&video.VideoNotifications{}) {
		dberr := db.AutoMigrate(&video.VideoNotifications{})
		if err == nil {
			err = dberr
		}
	}

	if err != nil {
		log.Panic("Failed to migrate User table.")
	}
}
