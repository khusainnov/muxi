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

func (h *Handler) InitRoutes(router mux.Router) {
	//router = mux.NewRouter()
	router.HandleFunc("/weather/{city:[aA-zZ]+}", func(w http.ResponseWriter, r *http.Request) {
		err := h.service.GetWeather(w, r)
		if err != nil {
			log.Fatalln()
		}
	})
}