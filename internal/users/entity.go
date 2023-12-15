package users

import (
	"crypto/md5"
	"errors"
	"fmt"
	"time"
)

var (
	ErrPasswordRequired = errors.New("password is required and can't be blank")
	ErrPasswordLen      = errors.New("password must have at least 6 characters")
	ErrNameRequired     = errors.New("name is required")
	ErrUserNameRequired = errors.New("username is required")
)

func New(name, username, password string) (*User, error) {
	now := time.Now()
	u := User{Name: name, Username: username, CreatedAt: now, ModifiedAt: now}
	err := u.SetPassword(password)
	if err != nil {
		return nil, err
	}

	err = u.Validate()
	if err != nil {
		return nil, err
	}

	return &u, nil
}

type User struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	Deleted    bool      `json:"-"`
	LastLogin  time.Time `json:"last_login"`
}

func (u *User) SetPassword(password string) error {

	if password == "" {
		return ErrPasswordRequired
	}

	if len(password) < 6 {
		return ErrPasswordLen
	}
	u.Password = fmt.Sprintf("%x", (md5.Sum([]byte(password))))
	return nil
}

func (u *User) Validate() error {

	if u.Name == "" {
		return ErrNameRequired
	}

	if u.Username == "" {
		return ErrUserNameRequired
	}

	if u.Password == fmt.Sprintf("%x", (md5.Sum([]byte("")))) {
		return ErrPasswordRequired
	}

	return nil
}
