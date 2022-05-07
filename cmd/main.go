package main

import (
	"log"

	muxi "test"
	"test/pkg/handler"
	"test/pkg/service"
)

func main() {
	services := service.NewService()
	handlers := handler.NewHandler(services)

	s := muxi.Server{}

	if err := s.Run(handlers.InitRoutes()); err != nil {
		log.Fatalln(err.Error())
	}
}
