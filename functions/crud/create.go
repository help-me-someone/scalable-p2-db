// This file contains the functions related to [c]rud.

package crud

import (
	"github.com/help-me-someone/scalable-p2-db/models/user"
	"github.com/help-me-someone/scalable-p2-db/models/video"
	"gorm.io/gorm"
)

/*----------------------
|  User
-----------------------*/

// CreateUser inserts a new entry into the User table.
// It should be noted that it does NOT expect to have
// to hash the password.
func CreateUser(db *gorm.DB, username, password string) (*user.User, error) {
	usr := &user.User{
		Username:       username,
		HashedPassword: password,
	}
	err := db.Create(usr).Error
	return usr, err
}

/*----------------------
|  Section
-----------------------*/

// CreateVideo inserts a new entry into the video table.
// It should be noted that upon insertion, the video state
// is set to 'UPLOADING'. This is because we expect external
// services to be the one who updates the Status field.
// The Public member is set to be false, denoting private by
// default.
func CreateVideo(db *gorm.DB, name, address string, owner uint) (*video.Video, error) {
	vid := &video.Video{
		Name:    name,
		Address: address,
		Status:  video.VIDEO_UPLOADING,
		Public:  false,
		UserID:  owner,
	}
	err := db.Create(vid).Error
	return vid, err
}
