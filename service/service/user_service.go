package service

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/PerforMance308/test1/database"
	"github.com/PerforMance308/test1/options"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// UserService handles rest get user requests
type UserService struct {
	wg    *sync.WaitGroup
	muxer *mux.Router
	db    *database.MockDB
}

// NewUserService returns UserService object
func NewUserService(wg *sync.WaitGroup) (*UserService, error) {
	wg.Add(1)

	s := &UserService{wg: wg}
	s.db = database.NewDbClient()

	s.InitData()
	return s, nil
}

// WithMuxer addes a gorilla http mux router to the service
func (s *UserService) WithMuxer(router *mux.Router) *UserService {
	s.muxer = router

	// adding endpoint
	s.muxer.HandleFunc("/user/{id:[0-9-]+}", s.HandleGetUser)
	s.muxer.HandleFunc("/users", s.HandleGetUsers)
	return s
}

// ServeREST service on a port
func (s *UserService) ServeREST() error {
	logrus.Infof("GetUserServer Start on %s:%d", options.REST.Address, options.REST.Port)
	defer s.wg.Done()
	defer logrus.Infoln("GetUserServer ENd")

	srv := &http.Server{
		Handler:      s.muxer,
		Addr:         fmt.Sprintf("%s:%d", options.REST.Address, options.REST.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv.ListenAndServe()
}
