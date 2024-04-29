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
	"sync"
	"time"
)

type person struct {
	UserId  int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type database struct {
	users       []person
	maxId       int
	fakeLatency time.Duration
	mutex       *sync.Mutex
	ctx         context.Context
}

func NewFakeDB(ctx context.Context, fakeLatency time.Duration) *database {
	return &database{
		users:       []person{},
		maxId:       0,
		fakeLatency: fakeLatency,
		mutex:       &sync.Mutex{},
		ctx:         ctx,
	}
}

func (db *database) housekeeping() func() {
	time.Sleep(db.fakeLatency)
	db.mutex.Lock()
	return func() {
		db.mutex.Unlock()
	}
}

func (db *database) AddUser(rdr io.Reader) (id int, err error) {
	defer db.housekeeping()()
	p := person{}
	if err := json.NewDecoder(rdr).Decode(&p); err != nil {
		return -1, err
	}
	id = db.maxId + 1
	db.maxId++
	db.users = append(db.users, p)
	return id, nil
}

func (db database) GetUserById(id int) (io.Reader, error) {
	defer db.housekeeping()()
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
	defer db.housekeeping()()
	var r bytes.Buffer
	err := json.NewEncoder(&r).Encode(db.users)
	if err != nil {
		return nil, fmt.Errorf("failed to encode: %w", err)
	}
	return &r, nil
}
