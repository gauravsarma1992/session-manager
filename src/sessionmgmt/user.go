package sessionmgmt

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model

		ID       string `json:"id"`
		Username string `json:"username"`
		Password string `json:"-"`

		Email  string `json:"email"`
		Mobile string `json:"mobile"`

		Sessions []*Session `json:"sessions",gorm:"-"'`
	}
)

func (user *User) BeforeCreate() (err error) {
	var (
		passwordB []byte
	)
	if passwordB, err = bcrypt.GenerateFromPassword([]byte(user.Password), 256); err != nil {
		return
	}
	user.Password = string(passwordB)
	return
}
