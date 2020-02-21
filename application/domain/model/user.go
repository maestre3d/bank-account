package model

import "errors"

type User struct {
	ID int64
	Name string
	Balance float32
}

func NewUser(name string) *User {
	return &User{
		ID: 0,
		Name: name,
		Balance: 0.0,
	}
}

func (u *User) IsValid() error {
	if u.Balance < 0 {
		return errors.New("invalid balance")
	} else if u.Name == "" {
		return errors.New("invalid name")
	}

	return nil
}


