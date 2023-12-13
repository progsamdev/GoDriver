package users

import "time"

func New(name, username, password string) (*User, error) {
	now := time.Now()
	u := User{Name: name, Username: username, CreatedAt: now, ModifiedAt: now}
	return &u, nil
}

type User struct {
	ID         int64
	Name       string
	Username   string
	Password   string
	CreatedAt  time.Time
	ModifiedAt time.Time
	Deleted    bool
	LastLogin  time.Time
}
