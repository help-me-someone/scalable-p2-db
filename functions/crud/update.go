// This file contains the functions related to cr[u]d.

package crud

import "gorm.io/gorm"

/*----------------------
|  Video
-----------------------*/

func UpdateVideoStatus(db *gorm.DB, videoID uint, status uint8) error {
	vid, err := GetVideo(db, videoID)
	if err != nil {
		return err
	}
	vid.Status = status
	return db.Save(&vid).Error
}

func UpdateVideoStatusByKey(db *gorm.DB, videoKey string, status uint8) error {
	vid, err := GetVideoByKey(db, videoKey)
	if err != nil {
		return err
	}
	vid.Status = status
	return db.Save(&vid).Error
}

// I think that this is horrible but this is the only easy way I can
// think of to be able to keep track of the video status with async
// jobs. I could probably improve this if I were to read more into
// asynq docs.
func UpdateVideoStatusIncrementByKey(db *gorm.DB, videoKey string) error {
	vid, err := GetVideoByKey(db, videoKey)
	if err != nil {
		return err
	}
	vid.Status += 1
	return db.Save(&vid).Error
}

func UpdateVideoPrivacy(db *gorm.DB, videoID uint, public bool) error {
	vid, err := GetVideo(db, videoID)
	if err != nil {
		return err
	}
	vid.Public = public
	return db.Save(&vid).Error
}

// Let's just say that views can only go up...
func UpdateVideoViewIncrement(db *gorm.DB, videoID uint) error {
	vid, err := GetVideo(db, videoID)
	if err != nil {
		return err
	}
	vid.Views += 1
	return db.Save(&vid).Error
}

// Toggle like. If on then off, vice versa.
func UpdateToggleVideoLike(db *gorm.DB, user_id, video_id uint) error {
	videoLike, err := GetVideoLike(db, user_id, video_id)
	if err != nil {
		return err
	}
	videoLike.Like = !videoLike.Like
	return db.Save(&videoLike).Error
}

/*----------------------
|  Video Notification
-----------------------*/

func UpdateNotificationAsRead(db *gorm.DB, user_id, video_id uint) error {
	notif, err := GetNotification(db, video_id, user_id)
	if err != nil {
		return err
	}
	notif.Read = true
	return db.Save(&notif).Error
}
