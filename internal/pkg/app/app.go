package app

import (
	"log"

	"github.com/Shadowman405/Golang_test/internal/app/endpoint"
	"github.com/Shadowman405/Golang_test/internal/app/mw"
	"github.com/Shadowman405/Golang_test/internal/app/service"
	"github.com/labstack/echo"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.s = service.New()
	a.e = endpoint.New(a.s)

	a.echo = echo.New()

	a.echo.Use(mw.RoleCheck)

	a.echo.GET("/status", a.e.Handler)

	return a, nil
}

func (a *App) Run() error {
	println("Server started")

	error := a.echo.Start(":8080")
	if error != nil {
		log.Fatal(error)
	}

	return nil
}
