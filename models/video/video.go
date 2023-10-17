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
	Name string

	// The video's key in s3.
	Address string

	// The video's status.
	Status uint8

	// Public status.
	Public bool

	// Owner ID.
	UserID uint
}
