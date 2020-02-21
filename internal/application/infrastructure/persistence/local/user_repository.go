package local

import (
	"errors"
	"github.com/maestre3d/bank-account/internal/application/domain/model"
)

type LocalUserRepository struct {
	DB *DBServer
}

func NewUserRepository(db *DBServer) *LocalUserRepository {
	return &LocalUserRepository{
		DB:db,
	}
}

func (u *LocalUserRepository) Save(model *model.User) error {
	u.DB.UserDB = append(u.DB.UserDB, model)
	return nil
}

func (u *LocalUserRepository) Update(model *model.User) error {
	return nil
}

func (u *LocalUserRepository) Remove(ID int64) error {
	return nil
}

func (u *LocalUserRepository) FetchAll() ([]*model.User, error) {
	users := u.DB.UserDB
	if len(users) <= 0 {
		return nil, errors.New("users not found")
	}

	return users, nil
}

func (u *LocalUserRepository) FetchByID(ID int64) (*model.User, error) {
	return nil, nil
}
