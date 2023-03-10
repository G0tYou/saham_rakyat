package controller

import (
	"encoding/json"
	"log"
	model "saham_rakyat/models"
	"strconv"

	helper "saham_rakyat/pkg/helper"
	repository "saham_rakyat/repository/mysql/order_histories"
	redis "saham_rakyat/repository/redis"

	"net/http"

	"github.com/labstack/echo/v4"
)

func Store(c echo.Context) error {
	log.Println(c.Request().RequestURI)
	order_histories := new(model.OrderHistories)
	if err := c.Bind(order_histories); err != nil {
		return err
	}

	err := repository.Store(order_histories)

	if err != nil {
		return helper.JSONResponse(c, http.StatusInternalServerError, err.Error())
	}
	err = redis.MultipleDelete("order/list")
	if err != nil {
		return helper.JSONResponse(c, http.StatusInternalServerError, err.Error())
	}
	return helper.JSONResponse(c, http.StatusOK, "OK")
}

func Update(c echo.Context) error {
	log.Println(c.Request().RequestURI)
	order_histories := new(model.OrderHistories)
	if err := c.Bind(order_histories); err != nil {
		return helper.JSONResponse(c, http.StatusInternalServerError, err.Error())
	}
	err := repository.Update(order_histories)
	if err != nil {
		return helper.JSONResponse(c, http.StatusInternalServerError, err.Error())
	}
	err = redis.Delete("order/detail/" + strconv.Itoa(order_histories.Id))
	if err != nil {
		return helper.JSONResponse(c, http.StatusInternalServerError, err.Error())
	}
	err = redis.MultipleDelete("order/list")
	if err != nil {
		return helper.JSONResponse(c, http.StatusInternalServerError, err.Error())
	}
	return helper.JSONResponse(c, http.StatusOK, "OK")
}

func Delete(c echo.Context) error {
	log.Println(c.Request().RequestURI)
	order_histories := new(model.OrderHistories)
	id, _ := strconv.Atoi(c.Param("id"))
	order_histories.Id = id
	err := repository.Delete(order_histories)
	if err != nil {
		return helper.JSONResponse(c, http.StatusInternalServerError, err.Error())
	}
	err = redis.Delete("order/detail/" + strconv.Itoa(id))
	if err != nil {
		return helper.JSONResponse(c, http.StatusInternalServerError, err.Error())
	}
	err = redis.MultipleDelete("order/list")
	if err != nil {
		return helper.JSONResponse(c, http.StatusInternalServerError, err.Error())
	}
	return helper.JSONResponse(c, http.StatusOK, "OK")
}

func Detail(c echo.Context) error {
	log.Println(c.Request().RequestURI)
	var order_histories model.OrderHistories
	id, _ := strconv.Atoi(c.Param("id"))
	value, _ := redis.Get("order/detail/" + strconv.Itoa(id))
	if len(value) > 1 {
		err := json.Unmarshal([]byte(value), &order_histories)
		if err != nil {
			return helper.JSONResponse(c, http.StatusInternalServerError, err.Error())
		}
	} else {
		detail, err := repository.GetDataById(id)
		if err != nil {
			return helper.JSONResponse(c, http.StatusInternalServerError, err.Error())
		}
		jsonBytes, err := json.Marshal(detail)
		if err != nil {
			return helper.JSONResponse(c, http.StatusInternalServerError, err.Error())
		}
		err = redis.Set("order/detail/"+strconv.Itoa(id), jsonBytes)
		if err != nil {
			return helper.JSONResponse(c, http.StatusInternalServerError, err.Error())
		}
		order_histories = detail
	}
	return helper.JSONResponse(c, http.StatusOK, order_histories)
}

func List(c echo.Context) error {
	log.Println(c.Request().RequestURI)
	var order_histories []model.OrderHistories
	page, _ := strconv.Atoi(c.Param("page"))
	limit, _ := strconv.Atoi(c.Param("limit"))

	value, _ := redis.Get("order/list/" + strconv.Itoa(page) + "/" + strconv.Itoa(limit))
	if len(value) > 1 {
		err := json.Unmarshal([]byte(value), &order_histories)
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
		err = redis.Set("order/list/"+strconv.Itoa(limit)+"/"+strconv.Itoa(page), jsonBytes)
		if err != nil {
			return helper.JSONResponse(c, http.StatusInternalServerError, err.Error())
		}
		order_histories = list
	}
	return helper.JSONResponse(c, http.StatusOK, order_histories)
}
