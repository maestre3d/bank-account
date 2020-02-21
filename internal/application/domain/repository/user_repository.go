package repository

import "github.com/maestre3d/bank-account/internal/application/domain/model"

type IUserRepository interface {
	Save(model *model.User) error
	Update(model *model.User) error
	Remove(ID int64) error
	FetchAll() ([]*model.User, error)
	FetchByID(ID int64) (*model.User, error)
}
