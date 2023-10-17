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
	db.Save(&vid)
	return nil
}

func UpdateVideoPrivacy(db *gorm.DB, videoID uint, public bool) error {
	vid, err := GetVideo(db, videoID)
	if err != nil {
		return err
	}
	vid.Public = public
	db.Save(&vid)
	return nil
}
