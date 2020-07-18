package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// User struct
type User struct {
	ID   int
	Name string
	Age  int
	City string
}

// InitData creates mock data into database
func (s *UserService) InitData() {
	user1 := &User{
		ID:   1,
		Name: "John",
		Age:  31,
		City: "New York",
	}

	user2 := &User{
		ID:   2,
		Name: "Doe",
		Age:  22,
		City: "Vancouver",
	}
	s.db.Put("user", user1.ID, user1)
	s.db.Put("user", user2.ID, user2)
}

// HandleGetUser returns mutiple user objects
func (s *UserService) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	logrus.Tracef("handling incoming request: %v\n", r)
	logrus.Infoln("HandleGetUser Start")
	defer logrus.Infoln("HandleGetUser End")

	// Get id and convert to int
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("failed to convert id to int, id: %s, err: %s", vars["id"], err)
		return
	}

	// read from db
	user := &User{}
	err = s.db.GetOne("user", userID, user)
	if err != nil {
		w.Write([]byte(err.Error()))
		logrus.Errorf("failed to get an item from table user, with id: %s, err: %s", vars["id"], err)
		return
	}

	// marshal response
	buf, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("failed to marshal return value: %v", err)
		return
	}

	// response
	if _, err := w.Write(buf); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("error writing response: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// HandleGetUsers returns mutiple user objects
func (s *UserService) HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	logrus.Tracef("handling incoming request: %v\n", r)
	logrus.Infoln("HandleGetUsers Start")
	defer logrus.Infoln("HandleGetUsers End")

	users := []User{}
	err := s.db.GetAll("user", &users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Fatalf("failed to get an item from table uses, error: %s", err)
		return
	}

	buf, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("failed to marshal return value: %v", err)
		return
	}

	if _, err := w.Write(buf); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("error writing response: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
