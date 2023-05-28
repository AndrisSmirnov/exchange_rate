package controllers

import (
	"exchange_rate/pkg"
	_http "exchange_rate/pkg/controllers/http"
	"log"
	"net/http"
	"os"
)

type Controllers struct {
	HTTP *_http.HTTPController
}

func NewControllers(services *pkg.Services) (*Controllers, error) {
	http, err := _http.NewHttpControllers(services)
	if err != nil {
		return nil, err
	}

	return &Controllers{
		HTTP: http,
	}, nil
}

func (c *Controllers) Start() {
	handlers := c.HTTP.InitControllers()
	url := os.Getenv("SERVER_URL")

	go listenAndServe(url, handlers)
}

func listenAndServe(url string, handlers http.Handler) {
	err := http.ListenAndServe(url, handlers)
	if err != nil {
		log.Panic(err)
	}
}
