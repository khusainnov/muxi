package main

import (
	"net/http"
	"time"

	"test/pkg/service"
)

func main() {

	service.NewService()

	s := &http.Server{
		Addr:              ":7001",
		//Handler:           ,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}

	s.ListenAndServe()
}