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
