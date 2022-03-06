package main

import (
	"app/config"
	"app/infra"
	"app/interface/handler"
	"app/usecase"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
)


func main() {
	itemRepository := infra.NewItemRepository(config.NewDB())
	itemUsecase := usecase.NewItemUsecase(itemRepository)
	itemHandler := handler.NewItemHandler(itemUsecase)

	e := echo.New()
	handler.InitRouting(e, itemHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
