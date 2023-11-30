package sessionmgmt

import (
	"errors"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model

		ID       uint   `json:"id" gorm:"primaryKey,autoIncrement"`
		Username string `json:"username" gorm:"unique"`
		Password string `json:"-"`

		Email  string `json:"email"`
		Mobile string `json:"mobile"`

		Sessions []*Session `json:"sessions" gorm:"-"`
	}
)

func (user *User) BeforeCreate(db *gorm.DB) (err error) {
	var (
		passwordB []byte
	)
	if passwordB, err = bcrypt.GenerateFromPassword([]byte(user.Password), 4); err != nil {
		return
	}
	user.Password = string(passwordB)
	return
}

func (user *User) Validate(reqUser *User) (err error) {
	if user.Username != reqUser.Username {
		err = errors.New("Username doesn't match")
		return
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqUser.Password)); err != nil {
		err = errors.New("Password doesn't match")
		return
	}
	return
}
