package helper

import (
	"log"

	"github.com/labstack/echo/v4"
)

func JSONResponse(c echo.Context, code int, i interface{}) error {
	log.Println(code)
	log.Println(i)
	return c.JSON(code, i)
}
