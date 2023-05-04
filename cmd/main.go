package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	server := echo.New()

	server.Use(MiddleWare)

	server.GET("/status", Handler)

	error := server.Start(":8080")
	if error != nil {
		log.Fatal(error)
	}
}

func Handler(ctx echo.Context) error {
	date := time.Date(2025, time.April, 1, 0, 0, 0, 0, time.UTC)
	duration := time.Until(date)

	strStatus := fmt.Sprintf("Number of days: %d", int64(duration.Hours())/24)

	ctx.String(http.StatusOK, strStatus)

	return nil
}

func MiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		value := ctx.Request().Header.Get("User-Role")

		if value == "admin" {
			log.Println("admin detected")
		}

		err := next(ctx)
		if err != nil {
			return err
		}

		return nil
	}
}