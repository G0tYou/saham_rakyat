package controller

import (
	"encoding/json"
	"log"
	model "saham_rakyat/models"
	"strconv"

	helper "saham_rakyat/pkg/helper"
	repository "saham_rakyat/repository/mysql/users"
	redis "saham_rakyat/repository/redis"

	"net/http"

	"github.com/labstack/echo/v4"
)

func Store(c echo.Context) error {
	log.Println(c.Request().RequestURI)
	user := new(model.Users)
	if err := c.Bind(user); err != nil {
		return err
	}

	err := repository.Store(user)

	if err != nil {
		return helper.JSONResponse(c, http.StatusInternalServerError, err.Error())
	}
	err = redis.MultipleDelete("user/list")
	if err != nil {
		return helper.JSONResponse(c, http.StatusInternalServerError, err.Error())
	}
	return helper.JSONResponse(c, http.StatusOK, "OK")
}

func Update(c echo.Context) error {
	log.Println(c.Request().RequestURI)
	user := new(model.Users)
	if err := c.Bind(user); err != nil {
		return err
	}
	err := repository.Update(user)
	if err != nil {
		return helper.JSONResponse(c, http.StatusInternalServerError, err.Error())
	}
	err = redis.Delete("user/detail/" + strconv.Itoa(user.Id))
	if err != nil {
		return helper.JSONResponse(c, http.StatusInternalServerError, err.Error())
	}
	err = redis.MultipleDelete("user/list")
	if err != nil {
		return helper.JSONResponse(c, http.StatusInternalServerError, err.Error())
	}
	return helper.JSONResponse(c, http.StatusOK, "OK")
}

func Delete(c echo.Context) error {
	log.Println(c.Request().RequestURI)
	user := new(model.Users)
	id, _ := strconv.Atoi(c.Param("id"))
	user.Id = id
	err := repository.Delete(user)
	if err != nil {
		return helper.JSONResponse(c, http.StatusInternalServerError, err.Error())
	}
	err = redis.Delete("user/detail/" + strconv.Itoa(id))
	if err != nil {
		return helper.JSONResponse(c, http.StatusInternalServerError, err.Error())
	}
	err = redis.MultipleDelete("user/list")
	if err != nil {
		return helper.JSONResponse(c, http.StatusInternalServerError, err.Error())
	}
	return helper.JSONResponse(c, http.StatusOK, "OK")
}

func Detail(c echo.Context) error {
	log.Println(c.Request().RequestURI)
	var user model.Users
	id, _ := strconv.Atoi(c.Param("id"))
	value, _ := redis.Get("user/detail/" + strconv.Itoa(id))
	if len(value) > 1 {
		err := json.Unmarshal([]byte(value), &user)
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
		err = redis.Set("user/detail/"+strconv.Itoa(id), jsonBytes)
		if err != nil {
			return helper.JSONResponse(c, http.StatusInternalServerError, err.Error())
		}
		user = detail
	}
	return helper.JSONResponse(c, http.StatusOK, user)
}

func List(c echo.Context) error {
	log.Println(c.Request().RequestURI)
	var user []model.Users
	page, _ := strconv.Atoi(c.Param("page"))
	limit, _ := strconv.Atoi(c.Param("limit"))

	value, _ := redis.Get("user/list/" + strconv.Itoa(page) + "/" + strconv.Itoa(limit))
	if len(value) > 1 {
		err := json.Unmarshal([]byte(value), &user)
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
		err = redis.Set("user/list/"+strconv.Itoa(limit)+"/"+strconv.Itoa(page), jsonBytes)
		if err != nil {
			return helper.JSONResponse(c, http.StatusInternalServerError, err.Error())
		}
		user = list
	}
	return helper.JSONResponse(c, http.StatusOK, user)
}
