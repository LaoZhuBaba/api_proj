// fakedb provides a simple backend for testing an interface.
// To avoid creating dependencies, we return JSON formatted data as an io.Reader
package fakedb

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type person struct {
	UserId  int    `json:"userid"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type database struct {
	users []person
	maxId int
}

func NewFakeDB(ctx context.Context) *database {
	return &database{
		users: []person{
			{UserId: 1, Name: "David", Address: "David's address"},
			{UserId: 2, Name: "Mary", Address: "Mary's address"},
			{UserId: 3, Name: "Fred", Address: "Fred's address"},
		},
		maxId: 3,
	}
}

func (db *database) AddUser(n string) (id int, err error) {
	id = db.maxId + 1
	db.maxId++
	p := person{Name: n, UserId: id}
	db.users = append(db.users, p)
	return id, nil
}

func (db database) GetUserById(id int) (io.Reader, error) {
	for _, p := range db.users {
		if p.UserId == id {
			var r bytes.Buffer
			err := json.NewEncoder(&r).Encode(p)
			if err != nil {
				return nil, fmt.Errorf("failed to encode for user %d: %v", id, err)
			}
			return &r, nil
		}
	}
	return nil, errors.New("user not found")
}

func (db database) GetAllUsers() (io.Reader, error) {
	var r bytes.Buffer
	err := json.NewEncoder(&r).Encode(db.users)
	if err != nil {
		return nil, fmt.Errorf("failed to encode: %w", err)
	}
	return &r, nil
}
