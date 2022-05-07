package handler

import (
	"log"
	"net/http"

	"test/pkg/service"

	"github.com/gorilla/mux"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router = &mux.Router{}

	router.HandleFunc("/weather/{city:[aA-zZ]+}", func(w http.ResponseWriter, r *http.Request) {
		err := h.service.GetWeather(w, r)
		if err != nil {
			log.Fatalln()
		}
	})

	return router
}
