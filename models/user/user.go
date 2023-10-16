// This file contains the schema for the user table.

package user

import "gorm.io/gorm"

// The User model.
type User struct {
	// ID, CreatedAt, UpdatedAt, DeletedAt.
	gorm.Model

	// The user's username.
	Username string

	// The user's hashed password.
	HashedPassword string
}
