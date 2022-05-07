package service

import (
	"net/http"
)

type GetWeatherData interface {
	GetWeather(rw http.ResponseWriter, r *http.Request) error
}

type Service struct {
	GetWeatherData
}

func NewService() *Service {
	return &Service{}
}