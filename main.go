package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	server := echo.New()

	server.GET("/status", Handler)

	error := server.Start(":8080")
	if error != nil {
		log.Fatal(error)
	}
}

func Handler(ctx echo.Context) error {
	ctx.String(http.StatusOK, "test")

	return nil
}
