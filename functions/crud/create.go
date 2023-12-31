// This file contains the functions related to [c]rud.

package crud

import (
	"time"

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
|  Video
-----------------------*/

// CreateVideo inserts a new entry into the video table.
// It should be noted that upon insertion, the video state
// is set to 'UPLOADING'. This is because we expect external
// services to be the one who updates the Status field.
// The Public member is set to be false, denoting private by
// default.
func CreateVideo(db *gorm.DB, name, key string, owner uint) (*video.Video, error) {
	vid := &video.Video{
		Name:   name,
		Key:    key,
		Status: video.VIDEO_CONVERTING,
		Public: false,
		UserID: owner,
	}
	err := db.Create(vid).Error
	return vid, err
}

/*----------------------
|  Video Likes
-----------------------*/

func CreateVideoLike(db *gorm.DB, video_id, user_id uint, like bool) (*video.VideoLikes, error) {
	vidLike := &video.VideoLikes{
		VideoID: video_id,
		UserID:  user_id,
		Like:    like,
	}
	err := db.Create(vidLike).Error
	return vidLike, err
}

/*----------------------
|  Video Comments
-----------------------*/

func CreateVideoComment(db *gorm.DB, video_id, user_id uint, comment string) (*video.VideoComments, error) {
	vidComment := &video.VideoComments{
		VideoID: video_id,
		UserID:  user_id,
		Comment: comment,
		Date:    time.Now(),
	}
	err := db.Create(vidComment).Error
	return vidComment, err
}

/*----------------------
|  Video Notification
-----------------------*/

func CreateVideoNotification(db *gorm.DB, video_id, actor_id, user_id uint, notification_type video.NotificationType) (*video.VideoNotifications, error) {
	vidNotif := &video.VideoNotifications{
		VideoID: video_id,
		ActorID: actor_id,
		UserID:  user_id,
		Read:    false,
		Type:    notification_type,
		Date:    time.Now(),
	}
	err := db.Create(vidNotif).Error
	return vidNotif, err
}
