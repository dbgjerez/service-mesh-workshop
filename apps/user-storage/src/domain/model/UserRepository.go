package model

import (
	"errors"
)

var UserNotFound = errors.New("User not found!")

type UserRepository struct {
	users []User
	store *Store
}

func NewUserRepository(store *Store) (*UserRepository, error) {
	users, err := store.ReadStore()
	if err != nil {
		return nil, err
	}
	return &UserRepository{
			users: users,
			store: store},
		nil
}

func (dao *UserRepository) HealthCheck() error {
	if !dao.store.Exists() {
		return errors.New("Database not found!")
	}
	return nil
}

func (dao *UserRepository) FindById(username string) (*User, error) {
	Users := dao.users
	for _, f := range Users {
		if f.Username == username {
			return &f, nil
		}
	}
	return nil, UserNotFound
}

func (dao *UserRepository) GetAll() []User {
	return dao.users
}
