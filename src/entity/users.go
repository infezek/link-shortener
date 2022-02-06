package entity

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type Users struct {
	ID        string    `json:"id,omitempty"`
	Username  string    `json:"username,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func (users *Users) validate() error {
	if users.Username == "" {
		return errors.New("the name is required and cannot be blank")
	}
	if len(users.Password) < 6 {
		return errors.New("username is too weak please create a validate username")
	}
	if len(users.Password) > 64 {
		return errors.New("this does not appear to be a valid username")
	}
	if users.Email == "" {
		return errors.New("the Email is required and cannot be blank")
	}
	if users.Password == "" {
		return errors.New("password is required and cannot be blank")
	}
	if len(users.Password) < 6 {
		return errors.New("password is too weak please create a stronger password")
	}
	if len(users.Password) > 128 {
		return errors.New("this does not appear to be a valid password")
	}
	if erro := checkmail.ValidateFormat(users.Email); erro != nil {
		return errors.New("the e-mail is invalid")
	}
	return nil
}

func (user *Users) format() error {
	user.Username = strings.TrimSpace(user.Username)
	user.Email = strings.TrimSpace(user.Email)

	return nil
}

func (user *Users) Prepare() error {
	if erro := user.validate(); erro != nil {
		return erro
	}

	if erro := user.format(); erro != nil {
		return erro
	}

	return nil
}
