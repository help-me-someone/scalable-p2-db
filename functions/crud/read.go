// This file contains the functions related to c[r]ud.

package crud

import (
	"github.com/help-me-someone/scalable-p2-db/models/user"
	"github.com/help-me-someone/scalable-p2-db/models/video"
	"gorm.io/gorm"
)

/*----------------------
|  User
-----------------------*/

// SearchUser returns the first user which matches the username specified.
// This function does NOT do any error handling for you.
func GetUserByName(db *gorm.DB, username string) (*user.User, error) {
	usr := &user.User{}
	err := db.Where(&user.User{Username: username}).First(usr).Error
	return usr, err
}

// GetUser returns the user with the given ID.
func GetUser(db *gorm.DB, ID uint) (*user.User, error) {
	usr := &user.User{}
	err := db.First(usr, ID).Error
	return usr, err
}

// GetUserVideos returns the list of videos the user owns.
func GetUserVideos(db *gorm.DB, userID uint) ([]video.Video, error) {
	videos := make([]video.Video, 0)
	db.Where("user_id=?", userID).Find(&video.Video{}).Scan(&videos)
	return videos, nil
}

/*----------------------
|  Video
-----------------------*/

func GetVideoByName(db *gorm.DB, name string) (*video.Video, error) {
	vid := &video.Video{}
	err := db.Where(&video.Video{Name: name}).First(vid).Error
	return vid, err
}

func GetVideo(db *gorm.DB, ID uint) (*video.Video, error) {
	vid := &video.Video{}
	err := db.First(vid, ID).Error
	return vid, err
}
