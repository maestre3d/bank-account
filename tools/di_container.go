package tools

import (
	"github.com/maestre3d/bank-account/internal/application/domain/model"
	"github.com/maestre3d/bank-account/internal/application/infrastructure/persistence/local"
	"github.com/maestre3d/bank-account/internal/application/infrastructure/persistence/rds"
	"github.com/maestre3d/bank-account/internal/application/infrastructure/service"
	"github.com/maestre3d/bank-account/internal/application/usecase"
	"github.com/maestre3d/bank-account/internal/view/cli/presenter"
	"github.com/sarulabs/di"
)

type Container struct {
	// Self container
	ctn di.Container
}

func NewContainer() (*Container, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return nil, err
	}

	err = builder.Add([]di.Def{
		{
			Name: "user.repository.local",
			Build: buildUserRepository,
		},
		{
			Name: "user.repository.rds",
			Build: buildUserRepository,
		},
		{
			Name: "user.usecase",
			Build: buildUserUseCase,
		},
		{
			Name: "user.presenter",
			Build: buildUserPresenter,
		},
	}...)

	return &Container{ctn:builder.Build()}, nil
}

func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

func (c *Container) Clean() error {
	return c.ctn.Clean()
}

func buildUserRepository(ctn di.Container) (interface{}, error) {

	db, err := service.ConnectPool("postgres://postgres:caca123@localhost:5432/bank-test?sslmode=disable")
	if err != nil {
		return nil, err
	}

	return rds.NewUserRepository(db), nil
}

func buildLocalUserRepository(ctn di.Container) (interface{}, error) {
	users := make([]*model.User, 0)
	db := local.DBServer{UserDB: users}
	return local.NewUserRepository(&db), nil
}

func buildUserUseCase(ctn di.Container) (interface{}, error) {
	return usecase.NewUserUseCase(ctn.Get("user.repository.rds").(*rds.UserRepository)), nil
}

func buildUserPresenter(ctn di.Container) (interface{}, error) {
	return presenter.NewUserPresenter(ctn.Get("user.usecase").(*usecase.UserUseCase)), nil
}