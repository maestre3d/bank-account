package presenter

import (
	"github.com/maestre3d/bank-account/application/domain/model"
	"github.com/maestre3d/bank-account/application/infrastructure/persistence/local"
	"github.com/maestre3d/bank-account/application/usecase"
	"log"
)

type UserPresenter struct {
	useCase *usecase.UserUseCase
}

func NewUserPresenter() *UserPresenter {
	users := make([]*model.User, 0)
	db := local.DBServer{UserDB:users}
	useCase := usecase.NewUserUseCase(local.NewUserRepository(&db))

	return &UserPresenter{useCase: useCase}
}

func (u *UserPresenter) CreateUser(name string) error {
	return u.useCase.CreateUser(name)
}

func (u *UserPresenter) GetAll() []*model.User {
	users, err := u.useCase.GetAllUsers()
	if err != nil || len(users) <= 0 {
		return nil
	}

	log.Println("Que pedo")
	log.Printf("%v", users)
	return users
}