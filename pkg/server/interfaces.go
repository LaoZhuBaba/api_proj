package server

import "io"

type LoggerAdapter func(string, ...any)

func (l LoggerAdapter) Logf(s string, other ...any) {
	l(s, other...)
}

type Logger interface {
	Logf(string, ...any)
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
