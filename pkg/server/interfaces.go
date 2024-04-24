package server

import "io"

type Logger interface {
	Debug(string, ...any)
	Error(string, ...any)
	Info(string, ...any)
	Warn(string, ...any)
}

type DataStore interface {
	GetAllUsers() (io.Reader, error)
	GetUserById(int) (io.Reader, error)
	AddUser(string, string) (int, error)
}

type Logic interface {
	GetUser(int) (Person, error)
	GetUsers() ([]Person, error)
	AddUser(Person) error
}

type LogicStream interface {
	Logic
	GetUsersStream() ([]Person, error)
}
