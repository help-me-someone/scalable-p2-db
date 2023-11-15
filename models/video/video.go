// This file contains the schema for the video table.

package video

import (
	"time"

	"gorm.io/gorm"
)

// Enums representing video status
const (
	VIDEO_CONVERTING uint8 = iota
	VIDEO_THUMBNAILING
	VIDEO_CHUNKING
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

type VideoLikes struct {
	// ID, CreatedAt, UpdatedAt, DeletedAt.
	gorm.Model

	// VideoID foreign key.
	VideoID uint `json:"video_id"`

	// UserID foreign key.
	UserID uint `json:"user_id"`

	// Dictates whether the user liked
	// the video or not.
	Like bool `json:"like"`
}

type VideoComments struct {
	// ID, CreatedAt, UpdatedAt, DeletedAt.
	ID uint `gorm:"primarykey" json:"id"`

	// VideoID foreign key.
	VideoID uint `json:"video_id"`

	// UserID foreign key.
	UserID uint `json:"user_id"`

	// The comment body.
	Comment string `json:"comment"`

	// When the comment was created.
	Date time.Time `json:"date"`
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

type NotificationType uint

const (
	Like NotificationType = iota
	Comment
)

// Note(Appy): The message is expected to be constructed, not stored (waste of space).
type VideoNotifications struct {
	// ID, CreatedAt, UpdatedAt, DeletedAt.
	ID uint `gorm:"primarykey" json:"id"`

	// VideoID foreign key.
	VideoID uint `json:"video_id"`

	// The actor who caused the notification.
	ActorID uint `json:"actor_id"`

	// UserID foreign key.
	UserID uint `json:"user_id"`

	// Whether the notification has been read.
	Read bool `json:"read"`

	// A message describing what caused the notification.
	Type NotificationType `json:"type"`

	// When the notification was made.
	Date time.Time `json:"date"`
}

type Recipient struct {
	UserID uint `json:"user_id"`
}
