package local

import (
	"errors"
	"github.com/maestre3d/bank-account/application/domain/model"
)

type UserRepository struct {
	DB *DBServer
}

var(
	userIDCount int64 = 0
)

func NewUserRepository(db *DBServer) *UserRepository {
	return &UserRepository{
		DB:db,
	}
}

func (u *UserRepository) Save(model *model.User) error {
	u.DB.UserDB = append(u.DB.UserDB, model)
	return nil
}

func (u *UserRepository) Update(model *model.User) error {
	return nil
}

func (u *UserRepository) Remove(ID int64) error {
	return nil
}

func (u *UserRepository) FetchAll() ([]*model.User, error) {
	users := u.DB.UserDB
	if len(users) <= 0 {
		return nil, errors.New("users not found")
	}

	return users, nil
}

func (u *UserRepository) FetchByID(ID int64) (*model.User, error) {
	return nil, nil
}
