package routes

import (
	migrate "saham_rakyat/controller/migrate"
	order_histories "saham_rakyat/controller/order_histories"
	order_item "saham_rakyat/controller/orders_item"
	user "saham_rakyat/controller/users"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	e.GET("/migrate", migrate.Migrate)

	// start endpoint user
	e.POST("/user/store", user.Store)
	e.PUT("/user/update", user.Update)
	e.DELETE("/user/delete/:id", user.Delete)
	e.GET("/user/detail/:id", user.Detail)
	e.GET("/user/list/:limit/:page", user.List)
	// end endpoint user

	// start endpoint order_item
	e.POST("/item/store", order_item.Store)
	e.PUT("/item/update", order_item.Update)
	e.DELETE("/item/delete/:id", order_item.Delete)
	e.GET("/item/detail/:id", order_item.Detail)
	e.GET("/item/list/:limit/:page", order_item.List)
	// end endpoint order_item

	// start endpoint order_histories
	e.POST("/order/store", order_histories.Store)
	e.PUT("/order/update", order_histories.Update)
	e.DELETE("/order/delete/:id", order_histories.Delete)
	e.GET("/order/detail/:id", order_histories.Detail)
	e.GET("/order/list/:limit/:page", order_histories.List)
	// end endpoint order_histories

}
