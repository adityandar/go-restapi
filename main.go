package main

import (
	"go_restapi/config"
	"go_restapi/domain/item/controllers"

	"github.com/labstack/echo/v4"
)

func main() {
	db := config.InitDB()

	route := echo.New()
	apiV1 := route.Group("api/v1/")

	itemController := controllers.NewItemController(db)
	apiV1.POST("item/create", itemController.Create)
	apiV1.PUT("item/update/:id_item", itemController.Update)
	apiV1.DELETE("item/delete/:id_item", itemController.Delete)
	apiV1.GET("item/all", itemController.GetAll)
	apiV1.GET("item/one", itemController.GetById)

	route.Start(":8000")
}
