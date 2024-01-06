package users

import "time"

func (h *handler) autenticate(login, password string) (*User, error) {
	user, err := GetByUsernameAndEncPassword(h.db, login, encPass(password))
	return user, err
}

func (h *handler) updateLastLogin(u *User) error {
	u.LastLogin = time.Now()
	return Update(h.db, int(u.ID), u)
}

func Autenticate(login, password string) (u *User, err error) {

	u, err = gh.autenticate(login, password)
	if err != nil {
		return nil, err
	}
	err = gh.updateLastLogin(u)
	if err != nil {
		return nil, err
	}
	return u, err
}
