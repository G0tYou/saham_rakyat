package controller

import (
	"encoding/json"
	"log"
	model "saham_rakyat/models"
	"strconv"

	helper "saham_rakyat/pkg/helper"
	repository "saham_rakyat/repository/mysql/orders_item"
	redis "saham_rakyat/repository/redis"

	"net/http"

	"github.com/labstack/echo/v4"
)

func Store(c echo.Context) error {
	log.Println(c.Request().RequestURI)
	orders_item := new(model.OrdersItem)
	if err := c.Bind(orders_item); err != nil {
		return helper.JSONResponse(c, http.StatusInternalServerError, err.Error())
	}
	err := repository.Store(orders_item)

	if err != nil {
		return helper.JSONResponse(c, http.StatusInternalServerError, err.Error())
	}
	err = redis.MultipleDelete("item/list")
	if err != nil {
		return helper.JSONResponse(c, http.StatusInternalServerError, err.Error())
	}
	return helper.JSONResponse(c, http.StatusOK, "OK")

}

func Update(c echo.Context) error {
	log.Println(c.Request().RequestURI)
	orders_item := new(model.OrdersItem)
	if err := c.Bind(orders_item); err != nil {
		return helper.JSONResponse(c, http.StatusInternalServerError, err.Error())
	}
	err := repository.Update(orders_item)
	if err != nil {
		return helper.JSONResponse(c, http.StatusInternalServerError, err.Error())
	}
	err = redis.Delete("item/detail/" + strconv.Itoa(orders_item.Id))
	if err != nil {
		return helper.JSONResponse(c, http.StatusInternalServerError, err.Error())
	}
	err = redis.MultipleDelete("item/list")
	if err != nil {
		return helper.JSONResponse(c, http.StatusInternalServerError, err.Error())
	}
	return helper.JSONResponse(c, http.StatusOK, "OK")
}

func Delete(c echo.Context) error {
	log.Println(c.Request().RequestURI)
	orders_item := new(model.OrdersItem)
	id, _ := strconv.Atoi(c.Param("id"))
	orders_item.Id = id
	err := repository.Delete(orders_item)
	if err != nil {
		return helper.JSONResponse(c, http.StatusInternalServerError, err.Error())
	}
	err = redis.Delete("item/detail/" + strconv.Itoa(id))
	if err != nil {
		return helper.JSONResponse(c, http.StatusInternalServerError, err.Error())
	}
	err = redis.MultipleDelete("item/list")
	if err != nil {
		return helper.JSONResponse(c, http.StatusInternalServerError, err.Error())
	}
	return helper.JSONResponse(c, http.StatusOK, "OK")
}

func Detail(c echo.Context) error {
	log.Println(c.Request().RequestURI)
	var orders_item model.OrdersItem
	id, _ := strconv.Atoi(c.Param("id"))
	value, _ := redis.Get("item/detail/" + strconv.Itoa(id))
	if len(value) > 1 {
		err := json.Unmarshal([]byte(value), &orders_item)
		if err != nil {
			return helper.JSONResponse(c, http.StatusInternalServerError, err.Error())
		}
	} else {
		detail, err := repository.GetDataById(id)
		if err != nil {
			return helper.JSONResponse(c, http.StatusInternalServerError, err.Error())
		}
		jsonBytes, err := json.Marshal(orders_item)
		if err != nil {
			return helper.JSONResponse(c, http.StatusInternalServerError, err.Error())
		}
		err = redis.Set("item/detail/"+strconv.Itoa(id), jsonBytes)
		if err != nil {
			return helper.JSONResponse(c, http.StatusInternalServerError, err.Error())
		}
		orders_item = detail
	}
	return helper.JSONResponse(c, http.StatusOK, orders_item)
}

func List(c echo.Context) error {
	log.Println(c.Request().RequestURI)
	var orders_item []model.OrdersItem
	page, _ := strconv.Atoi(c.Param("page"))
	limit, _ := strconv.Atoi(c.Param("limit"))

	value, _ := redis.Get("item/list/" + strconv.Itoa(page) + "/" + strconv.Itoa(limit))
	if len(value) > 1 {
		err := json.Unmarshal([]byte(value), &orders_item)
		if err != nil {
			return helper.JSONResponse(c, http.StatusInternalServerError, err.Error())
		}
	} else {
		list, err := repository.GetAllData(page, limit)
		if err != nil {
			return helper.JSONResponse(c, http.StatusInternalServerError, err.Error())
		}
		jsonBytes, err := json.Marshal(list)
		if err != nil {
			return helper.JSONResponse(c, http.StatusInternalServerError, err.Error())
		}
		err = redis.Set("item/list/"+strconv.Itoa(limit)+"/"+strconv.Itoa(page), jsonBytes)
		if err != nil {
			return helper.JSONResponse(c, http.StatusInternalServerError, err.Error())
		}
		orders_item = list
	}
	return helper.JSONResponse(c, http.StatusOK, orders_item)
}
