package controller

import (
	"log"
	helper "saham_rakyat/pkg/helper"
	history_model "saham_rakyat/repository/mysql/order_histories"
	item_model "saham_rakyat/repository/mysql/orders_item"
	user_model "saham_rakyat/repository/mysql/users"

	"net/http"

	"github.com/labstack/echo/v4"
)

func Migrate(c echo.Context) error {
	log.Println(c.Request().RequestURI)
	err := user_model.Migrate()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	err = item_model.Migrate()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	err = history_model.Migrate()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return helper.JSONResponse(c, http.StatusOK, "OK")
}
