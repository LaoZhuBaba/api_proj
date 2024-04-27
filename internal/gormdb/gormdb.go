package gormdb

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"

	"gorm.io/driver/sqlite" // Sqlite driver based on GGO
	"gorm.io/gorm"
)

const (
	dbpath = "file::memory:?cache=shared" // Use a sqllite in-memory DB
	// dbpath = "/Users/david/coding/go_play/api_proj/internal/gormdb/gorm.db"
)

type person struct {
	ID      int    `json:"userid"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type GormDb struct {
	db *gorm.DB
}

func NewGormDb() (db *GormDb, err error) {
	slog.Info("Initialising SqlLite database")
	d, err := gorm.Open(sqlite.Open(dbpath), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	d.AutoMigrate(&person{})
	return &GormDb{db: d}, err
}

func (g GormDb) GetAllUsers() (rdr io.Reader, err error) {
	var p []person
	var b []byte
	result := g.db.Find(&p)
	if result.Error != nil {
		return nil, result.Error
	}
	b, err = json.Marshal(p)
	if err != nil {
		return nil, err
	}
	rdr = bytes.NewReader(b)
	return rdr, nil
}

func (g GormDb) GetUserById(id int) (rdr io.Reader, err error) {
	var p person
	var b []byte
	result := g.db.First(&p, id)
	if result.Error != nil {
		return nil, result.Error
	}
	b, err = json.Marshal(p)
	if err != nil {
		return nil, err
	}
	rdr = bytes.NewReader(b)
	return rdr, nil
}

func (g GormDb) AddUser(name string, address string) (int, error) {
	p := person{Name: name, Address: address}
	result := g.db.Create(&person{Name: name, Address: address})
	if result.Error != nil {
		return 0, result.Error
	}
	return p.ID, nil
}
