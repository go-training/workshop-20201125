package model

import (
	"github.com/rs/zerolog/log"
)

// User struct
type User struct {
	ID    int64  `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// CreateUser data
func CreateUser(user *User) error {
	return DB.Create(user).Error
}

// FindAllUsers data
func FindAllUsers() ([]*User, error) {
	users := []*User{}
	return users, DB.Find(&users).Error
}

// FindUserByID ...
func FindUserByID(id int64) *User {
	user := &User{}
	if err := DB.First(user, id).Error; err != nil {
		log.Debug().Err(err).Msg("can't find user")
	}

	return user
}

// DeleteUser ...
func DeleteUser(id int64) *User {
	user := &User{}
	if err := DB.First(user, id).Error; err != nil {
		log.Debug().Err(err).Msg("can't find user")
	}

	DB.Delete(user, id)

	return user
}
