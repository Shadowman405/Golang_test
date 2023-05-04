package mw

import (
	"log"

	"github.com/labstack/echo"
)

const roleAdmin = "admin"

func RoleCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		value := ctx.Request().Header.Get("User-Role")

		if value == roleAdmin {
			log.Println("admin detected")
		}

		err := next(ctx)
		if err != nil {
			return err
		}

		return nil
	}
}
