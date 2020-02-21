package rds

import (
	"database/sql"
	"errors"
	"github.com/maestre3d/bank-account/internal/application/domain/model"
	"log"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB:db,
	}
}

func (u *UserRepository) Save(model *model.User) error {
	// Store new object into DB
	statement := `INSERT INTO USERS(NAME) VALUES ($1)
	RETURNING id`
	id := 0
	err := u.DB.QueryRow(statement, model.Name).Scan(&id)
	if err != nil {
		log.Fatal(err)
		return err
	} else if id == 0 {
		return errors.New("cannot save user")
	}
	log.Printf("User %d created", id)

	return nil
}

func (u *UserRepository) Update(model *model.User) error {
	return nil
}

func (u *UserRepository) Remove(ID int64) error {
	return nil
}

func (u *UserRepository) FetchAll() ([]*model.User, error) {
	rows, err := u.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*model.User, 0)
	for rows.Next() {
		user := new(model.User)
		err := rows.Scan(&user.ID, &user.Name, &user.Balance)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserRepository) FetchByID(ID int64) (*model.User, error) {
	return nil, nil
}
