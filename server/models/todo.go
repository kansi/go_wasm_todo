package models

import (
	"encoding/json"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"time"
)

type (
	// Todo represents the structure
	Todo struct {
		ID    string    `json:"id"`
		Title string    `json:"gender"`
		Body  string    `json:"age"`
		Date  time.Time `json:"date"`
	}

	// Db Database connection
	Db struct {
		Conn *leveldb.DB
	}
)

// NewDbConnection for storage
func NewDbConnection(path string) *Db {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		panic(fmt.Sprintf("Unable to open db file"))
	}

	return &Db{Conn: db}
}

// GetTodo retrieve Todo with `id` from the database
func (db *Db) GetTodo(id string) (*Todo, error) {
	data, _ := db.Conn.Get([]byte(id), nil)
	todo := Todo{}
	err := json.Unmarshal(data, &todo)
	return &todo, err
}

// StoreTodo retrieve Todo with `id` from the database
func (db *Db) StoreTodo(todo *Todo) error {
	todoJSON, _ := json.Marshal(todo)
	return db.Conn.Put([]byte(todo.ID), todoJSON, nil)
}

// DeleteTodo retrieve Todo with `id` from the database
func (db *Db) DeleteTodo(id string) error {
	return db.Conn.Delete([]byte(id), nil)
}
