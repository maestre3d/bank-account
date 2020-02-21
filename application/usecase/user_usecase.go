package usecase

import (
	"github.com/maestre3d/bank-account/application/domain/model"
	"github.com/maestre3d/bank-account/application/domain/repository"
)

type UserUseCase struct {
	userRepository repository.IUserRepository
}

func NewUserUseCase(userRepository repository.IUserRepository) *UserUseCase {
	return &UserUseCase{userRepository}
}

func (u *UserUseCase) CreateUser(name string) error {
	user := model.NewUser(name)
	if err := user.IsValid(); err != nil {
		return err
	}

	return u.userRepository.Save(user)
}

func (u *UserUseCase) GetAllUsers() ([]*model.User, error) {
	return u.userRepository.FetchAll()
}
