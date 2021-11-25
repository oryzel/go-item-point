package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/oryzel/go-item-point/src/items"
	"log"
	"strconv"
)

type ItemHandler struct {
	ItemUseCase items.Usecase
}

func ApplyRouter(r *gin.Engine, iuc items.Usecase) {

	items := r.Group("/items")
	handler := ItemHandler{
		ItemUseCase: iuc,
	}
	items.GET("", handler.getItems)
	items.POST("", handler.storeItem)
	items.GET("/:id", handler.getItem)
	items.PUT("/:id", handler.updateItem)
	items.DELETE("/:id", handler.deleteItem)
}

func (handler *ItemHandler) getItems(c *gin.Context) {
	res := handler.ItemUseCase.GetItems()
	c.JSON(200, res)
}

func (handler *ItemHandler) storeItem(c *gin.Context) {
	var item items.Item
	if err := c.BindJSON(&item); err != nil {
		log.Fatal(err)
	}
	handler.ItemUseCase.StoreItem(item)
	c.JSON(200, item)
}

func (handler *ItemHandler) getItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, err.Error())
	}
	res := handler.ItemUseCase.GetItem(id)
	c.JSON(200, res)
}

func (handler *ItemHandler) updateItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}

	var item items.Item
	if err := c.BindJSON(&item); err != nil {
		log.Fatal(err)
	}

	handler.ItemUseCase.UpdateItem(id, item)
}

func (handler *ItemHandler) deleteItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal("wrong value")
	}
	handler.ItemUseCase.DeleteItem(id)
	fmt.Println("Delete Item")
}
