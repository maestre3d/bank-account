package cli

import (
	"errors"
	"fmt"
	helper "github.com/maestre3d/bank-account/internal/view/cli/helper"
	presenter "github.com/maestre3d/bank-account/internal/view/cli/presenter"
	"log"
	"os"
)

var (
	userPresenter *presenter.UserPresenter
)

func Start() {
	userPresenter = presenter.NewUserPresenter()
	showMainMenu()
}

func showMainMenu() {
	var selectedOption string
	fmt.Println("Welcome to Foo Bank")
	fmt.Println("1. User\n2. Product\n3. Store\nPress any key to exit")
	_, err := fmt.Scan(&selectedOption)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	switch selectedOption {
	case "1":
		showUserMenu()
		fmt.Println("User created")
		break
	default:
		os.Exit(1)
	}
}

func showUserMenu() {
	var selectedOption string
	fmt.Println("User Operations\n1. Create\n2. Delete\n3. Show all\nPress any key to exit")
	_, err := fmt.Scan(&selectedOption)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	switch selectedOption {
	case "1":
		createUser()
		showUserMenu()
		break
	case "3":
		showAllUsers()
		showUserMenu()
		break
	default:
		showMainMenu()
		break
	}
}

func createUser() {
	var name string
	fmt.Println("Enter your full name:")
	_, err := fmt.Scan(&name)
	helper.IsValid(err)

	if name == "" {
		helper.IsValid(errors.New("invalid name"))
	}


	helper.IsValid(userPresenter.CreateUser(name))
}

func showAllUsers() {
	users := userPresenter.GetAll()
	for i := range users {
		fmt.Printf("User %d: \n%v\n", i, users[i])
	}
}


