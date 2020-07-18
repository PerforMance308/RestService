package main

import (
	"sync"

	"github.com/PerforMance308/test1/options"
	"github.com/PerforMance308/test1/service/service"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	var wg sync.WaitGroup

	// init go flags
	options.Init("rest", "logging")

	// create get user service
	s, err := service.NewUserService(&wg)

	if err != nil {
		logrus.Panicf("error creating REST service, error: %s", err)
	}

	// serve rest service in goroutin
	go s.WithMuxer(mux.NewRouter()).ServeREST()

	wg.Wait()
}
