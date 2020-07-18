package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDbClient(t *testing.T) {
	assert := assert.New(t)
	assert.NotNil(NewDbClient())
}
func TestPut(t *testing.T) {
	db := NewDbClient()
	db.Put("user", "1", "{Id : 1}")

	item := Datas["1"]
	assert.Equal(t, string(item), "\"{Id : 1}\"")

	db.Put("user", "2", "{Id : 2}")
	item = Datas["2"]
	assert.Equal(t, string(item), "\"{Id : 2}\"")
}
