package presenter

import (
	"github.com/maestre3d/bank-account/internal/application/domain/model"
	"github.com/maestre3d/bank-account/internal/application/usecase"
)

type UserPresenter struct {
	useCase *usecase.UserUseCase
}

func NewUserPresenter(useCase *usecase.UserUseCase) *UserPresenter {

	return &UserPresenter{useCase: useCase}
}

// CreateUser creates a user
func (u *UserPresenter) CreateUser(name string) error {
	return u.useCase.CreateUser(name)
}

func (u *UserPresenter) GetAll() []*model.User {
	users, err := u.useCase.GetAllUsers()
	if err != nil || len(users) <= 0 {
		return nil
	}
	return users
}