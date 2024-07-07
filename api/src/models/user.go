package models

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	NickName  string    `json:"nickname,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
}

func (user *User) Prepare(etapa string) error {
	if erro := user.validator(etapa); erro != nil {
		return erro
	}

	if erro := user.format(etapa); erro != nil {
		return erro
	}

	return nil
}

func (user *User) validator(stage string) error {
	if user.Name == "" {
		return errors.New("the name is mandatory and cannot be blank")
	}

	if user.NickName == "" {
		return errors.New("the nickname is mandatory and cannot be blank")
	}

	if user.Email == "" {
		return errors.New("email is mandatory and cannot be blank")
	}

	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return errors.New("the email entered is invalid")
	}

	if stage == "Create" && user.Password == "" {
		return errors.New("the password is mandatory and cannot be blank")
	}

	return nil
}

func (user *User) format(stage string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.NickName = strings.TrimSpace(user.NickName)
	user.Email = strings.TrimSpace(user.Email)

	if stage == "Create" {
		Hash := user.Password

		user.Password = string(Hash)
	}

	return nil
}
