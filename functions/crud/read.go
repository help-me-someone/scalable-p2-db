// This file contains the functions related to c[r]ud.

package crud

import (
	"fmt"

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
	err := db.Where("user_id=?", userID).Find(&videos).Error
	return videos, err
}

// GetTopPopularVideos returns the list of videos which are top ranked in terms of views.
// For N resutls queried, it returns [start:start+amount].
func GetTopPopularVideos(db *gorm.DB, page, amount int) ([]video.VideoWithUserEntry, error) {
	query := fmt.Sprintf("SELECT * FROM users, videos WHERE users.id = videos.user_id ORDER BY videos.views LIMIT %d OFFSET %d", amount, page)
	entries := make([]video.VideoWithUserEntry, 0)
	err := db.Raw(query).Find(&entries).Error
	return entries, err
}

func GetUserVideo(db *gorm.DB, videoName string, userID uint) (*video.Video, error) {
	video := &video.Video{}
	err := db.Where("user_id=? AND name=?", userID, videoName).First(video).Error
	return video, err
}

/*----------------------
|  Video
-----------------------*/

func GetVideoByName(db *gorm.DB, name string) (*video.Video, error) {
	vid := &video.Video{}
	err := db.Where(&video.Video{Name: name}).First(vid).Error
	return vid, err
}

func GetVideoByKey(db *gorm.DB, key string) (*video.Video, error) {
	vid := &video.Video{}
	err := db.Where(&video.Video{Key: key}).First(vid).Error
	return vid, err
}

func GetVideo(db *gorm.DB, ID uint) (*video.Video, error) {
	vid := &video.Video{}
	err := db.First(vid, ID).Error
	return vid, err
}
