// This file contains the functions related to [c]rud.

package crud

import (
	"github.com/help-me-someone/scalable-p2-db/models/user"
	"gorm.io/gorm"
)

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
