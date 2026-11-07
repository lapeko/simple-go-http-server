package api

import (
	"counter/internal/api/handlers"
	"fmt"
	"net/http"

	_ "counter/docs"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type Api interface {
	Setup()
	Start() error
}

type api struct {
	port string
}

func NewApi(port string) Api {
	if port == "" {
		port = "3000"
	}
	return &api{
		port: port,
	}
}

func (a *api) Setup() {
	http.Handle("/counter/", handlers.NewCounter())
	http.Handle("/swagger/", httpSwagger.WrapHandler)
}

func (a *api) Start() error {
	fmt.Printf("Server run on port: %s\n", a.port)
	return http.ListenAndServe(fmt.Sprintf(":%s", a.port), nil)
}
