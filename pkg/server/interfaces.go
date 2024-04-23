package server

import "io"

type LoggerAdapter func(string)

func (l LoggerAdapter) Log(s string) {
	l(s)
}

type Logger interface {
	Log(string)
}

type DataStore interface {
	GetAllUsers() (io.Reader, error)
	GetUserById(int) (io.Reader, error)
	AddUser(string) (int, error)
}

type Logic interface {
	GetUser(int) (Person, error)
	GetUsers() ([]Person, error)
}

type LogicStream interface {
	Logic
	GetUsersStream() ([]Person, error)
}
