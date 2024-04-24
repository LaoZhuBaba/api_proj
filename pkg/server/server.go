package server

import (
	"encoding/json"
)

type SimpleLogic struct {
	l  Logger
	ds DataStore
}

type Person struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func (sl SimpleLogic) GetUser(id int) (person Person, err error) {
	r, err := sl.ds.GetUserById(id)
	if err != nil {
		sl.l.Logf("failed to get user for id %d: %v", id, err)
		return Person{}, err
	}
	err = json.NewDecoder(r).Decode(&person)
	if err != nil {
		sl.l.Logf("failed to decode user for id %d: %v", id, err)
		return person, err
	}
	sl.l.Logf("returning user name: %s for id: %d", person.Name, id)
	return person, nil
}

func (sl SimpleLogic) AddUser(person Person) (err error) {
	sl.ds.AddUser(person.Name, person.Address)
	return nil
}

func (sl SimpleLogic) GetUsers() (persons []Person, err error) {
	r, err := sl.ds.GetAllUsers()
	if err != nil {
		sl.l.Logf("failed to get user list: %v", err)
		return nil, err
	}
	err = json.NewDecoder(r).Decode(&persons)
	if err != nil {
		sl.l.Logf("failed to decode user list: %v", err)
		return persons, err
	}
	sl.l.Logf("returning user name list")
	return persons, nil

}

func (sl SimpleLogic) GetUsersStream() (persons []Person, err error) {
	r, err := sl.ds.GetAllUsers()
	if err != nil {
		sl.l.Logf("failed to get user list: %v", err)
		return nil, err
	}
	err = json.NewDecoder(r).Decode(&persons)
	if err != nil {
		sl.l.Logf("failed to decode user list: %v", err)
		return persons, err
	}
	sl.l.Logf("returning user name list")
	return persons, nil

}

func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic{l: l, ds: ds}
}
