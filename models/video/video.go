// This file contains the schema for the video table.

package video

import "gorm.io/gorm"

// Enums representing video status
const (
	VIDEO_UPLOADING uint8 = iota
	VIDEO_CONVERTING
	VIDEO_READY
)

// The Video model.
type Video struct {
	// ID, CreatedAt, UpdatedAt, DeletedAt.
	gorm.Model

	// The video's name.
	Name string `json:"name"`

	// The video's key in s3.
	Key string `json:"key"`

	// The video's status.
	Status uint8 `json:"status"`

	// Public status.
	Public bool `json:"public"`

	// Owner ID.
	UserID uint `json:"user_id"`

	// Views.
	Views uint `json:"views"`
}

type VideoWithUserEntry struct {
	// The video's name.
	Name string `json:"name"`

	// The video's key in s3.
	Key string `json:"key"`

	// The owner's name.
	Username string `json:"username"`

	// Views.
	Views uint `json:"views"`
}
