package handler

import (
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

func (h *Handler) InitRoutes() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/weather/{city:[aA-zZ]+}", h.service.GetWeather).Methods("GET")

	return router
}
