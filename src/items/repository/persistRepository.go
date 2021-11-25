package repository

import (
	"fmt"
	"github.com/oryzel/go-item-point/src/items"
	"gorm.io/gorm"
)

type ItemPersist struct {
	db *gorm.DB
}

func NewPersist(db *gorm.DB) items.PersistRepository {
	return &ItemPersist{
		db: db,
	}
}

func (persist *ItemPersist) GetById(id int) items.Item {
	var item items.Item
	persist.db.First(&item, id)
	return item
}

func (persist *ItemPersist) Get() []items.Item {
	var itemsData []items.Item
	persist.db.Find(&itemsData)
	return itemsData
}

func (persist *ItemPersist) Insert(item items.Item) items.Item {
	persist.db.Create(&item)
	return item
}

func (persist *ItemPersist) Update(id int, item items.Item) items.Item {
	item.Id = id
	persist.db.Updates(&item)
	return item
}

func (persist *ItemPersist) Delete(id int) error {
	item := items.Item{}
	item.Id = id
	res := persist.db.Delete(&item)
	fmt.Println(res)
	if res.Error != nil {
		return res.Error
	}
	return nil

}
