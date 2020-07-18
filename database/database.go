package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// Datas contains mock datas in database
var Datas = map[interface{}][]byte{}

// MockDB contains a Database client
type MockDB struct {
}

// NewDbClient creates a DB client, could be either kind of Database, just mock here
func NewDbClient() *MockDB {
	db := &MockDB{}
	return db
}

// GetOne handle get operation of database, returns one item
func (db *MockDB) GetOne(tableName string, key, out interface{}) (err error) {
	if v, err := Datas[key]; err {
		err := json.Unmarshal(v, out)
		return err
	}

	return errors.New("undefined record")
}

// GetAll handle get operation of database, returns all items
func (db *MockDB) GetAll(tableName, out interface{}) (err error) {
	var temp []string
	for _, v := range Datas {
		temp = append(temp, fmt.Sprintf("%s", string(v)))
	}

	tempStr := "[" + strings.Join(temp, ",") + "]"
	err = json.Unmarshal([]byte(tempStr), out)
	return
}

// Put handle put operation of database
func (db *MockDB) Put(tableName string, key, item interface{}) (err error) {
	record, err := json.Marshal(item)
	Datas[key] = record
	return
}
