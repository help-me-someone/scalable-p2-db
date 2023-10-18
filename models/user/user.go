// This file contains the schema for the user table.

package user

import (
	"github.com/help-me-someone/scalable-p2-db/models/video"
	"gorm.io/gorm"
)

// The User model.
type User struct {
	// ID, CreatedAt, UpdatedAt, DeletedAt.
	gorm.Model

	// The user's username.
	Username string `json:"username"`

	// The user's hashed password.
	HashedPassword string `json:"hashed_password,omitempty"`

	// Videos owned by the user.
	Videos []video.Video `json:"videos,omitempty"`
}

type UserAPI struct {
	Username string `json:"username"`
	ID       uint   `json:"id,omitempty"`
}
