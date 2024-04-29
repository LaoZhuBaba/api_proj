package gormdb

import (
	"bytes"
	"context"
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
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type GormDb struct {
	db  *gorm.DB
	ctx context.Context
}

func NewGormDb(ctx context.Context) (db *GormDb, err error) {
	slog.Info("Initialising SqlLite database")
	d, err := gorm.Open(sqlite.Open(dbpath), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// wrap our db in a context
	d = d.WithContext(ctx)
	err = d.AutoMigrate(&person{})
	if err != nil {
		return nil, err
	}
	return &GormDb{ctx: ctx, db: d}, err
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

func (g GormDb) AddUser(rdr io.Reader) (int, error) {
	p := &person{}
	if err := json.NewDecoder(rdr).Decode(p); err != nil {
		return -1, err
	}
	result := g.db.Create(p)
	if result.Error != nil {
		return 0, result.Error
	}
	return p.ID, nil
}
