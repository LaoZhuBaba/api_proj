package interfaces

import (
	backend "github.com/laozhubaba/api_proj/backends"
)

type Backend interface {
	AddUser(string) (backend.Person, error)
	GetUserById(int) (backend.Person, error)
	GetAllUsers() []backend.Person
}

type Frontend interface {
	Run() error
}
