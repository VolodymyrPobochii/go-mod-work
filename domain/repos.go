package domain

import "github.com/pkg/errors"

var _ Repository = (*usersRepo)(nil)
var _ CrudRepository[User, string] = (*usersRepo)(nil)

const errorEmailNotFound = "User with [email:%s] not found"

type CrudRepository[T any, ID any] interface {
	FindOne(id ID) (*T, error)
	FindAll() []T
	Save(t *T) error
	Delete(t *T) error
	Exists(id ID) (bool, error)
}

type User struct {
	ID    string
	Email string
	Name  string
}

type Repository interface {
	CrudRepository[User, string]
	FindByEmail(email string) (*User, error)
	FindByName(name string) (*User, error)
}

type usersRepo struct {
	db []User
}

func (u *usersRepo) FindByEmail(email string) (*User, error) {
	for _, user := range u.db {
		if user.Email == email {
			return &user, nil
		}
	}
	return nil, errors.Errorf(errorEmailNotFound, email)
}

func (u *usersRepo) FindByName(name string) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *usersRepo) FindOne(id string) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *usersRepo) FindAll() []User {
	//TODO implement me
	panic("implement me")
}

func (u *usersRepo) Save(t *User) error {
	//TODO implement me
	panic("implement me")
}

func (u *usersRepo) Delete(t *User) error {
	//TODO implement me
	panic("implement me")
}

func (u *usersRepo) Exists(id string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func NewUsersRepo() Repository {
	return &usersRepo{}
}

type Account struct {
	ID   string
	Type int
}
