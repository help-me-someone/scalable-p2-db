// This file contains the functions related to c[r]ud.

package crud

import (
	"github.com/help-me-someone/scalable-p2-db/models/user"
	"gorm.io/gorm"
)

// SearchUser returns the first user which matches the username specified.
// This function does NOT do any error handling for you.
func GetUserByName(db *gorm.DB, username string) (*user.User, error) {
	usr := &user.User{}
	err := db.Where(&user.User{Username: username}).First(usr).Error
	return usr, err
}

// GetUser returns the user with the given ID.
func GetUser(db *gorm.DB, ID uint) (*user.User, error) {
	usr := &user.User{}
	err := db.First(usr, ID).Error
	return usr, err
}
