package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/oryzel/go-item-point/src/config"
	itemHttp "github.com/oryzel/go-item-point/src/items/handler/http"
	itemRepository "github.com/oryzel/go-item-point/src/items/repository"
	itemUsecase "github.com/oryzel/go-item-point/src/items/usecase"
)

func main() {
	conf := config.Init()
	db := conf.DB.Connect()

	app := gin.Default()
	app.GET("/health", func(context *gin.Context) {
		fmt.Println("Health")
	})

	ipr := itemRepository.NewPersist(db)
	ir := itemRepository.New(ipr)
	iu := itemUsecase.New(ir)
	itemHttp.ApplyRouter(app, iu)

	app.Run(":8181")

}
