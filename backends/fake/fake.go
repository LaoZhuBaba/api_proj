package fake

import (
	"context"
	"errors"

	backend "github.com/laozhubaba/api_proj/backends"
)

type database struct {
	users []backend.Person
	maxId int
}

func NewFakeDB(ctx context.Context) (*database, error) {
	return &database{}, nil
}
func (db *database) AddUser(n string) (p backend.Person, err error) {
	id := db.maxId + 1
	db.maxId++
	p = backend.Person{Name: n, UserId: id}
	db.users = append(db.users, p)
	return p, nil
}

func (db database) GetUserById(id int) (p backend.Person, err error) {
	for _, p := range db.users {
		if p.UserId == id {
			return p, nil
		}
	}
	return backend.Person{}, errors.New("user not found")
}

func (db database) GetAllUsers() []backend.Person {
	return db.users
}
