package service

import (
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/PerforMance308/test1/database"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestHandleGetUsers(t *testing.T) {
	assert := assert.New(t)
	db := database.NewDbClient()

	assert.NotNil(db, "database should have values")

	var wg sync.WaitGroup
	svc, _ := NewUserService(&wg)
	svc.WithMuxer(mux.NewRouter())
	assert.NotNil(svc, "service should be initialized")

	wr := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/users", nil)

	svc.HandleGetUsers(wr, r)

	r = httptest.NewRequest("GET", "/user/2", nil)
	svc.HandleGetUser(wr, r)
}
