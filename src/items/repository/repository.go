package repository

import (
	"github.com/oryzel/go-item-point/src/items"
)

type itemRepository struct {
	persist items.PersistRepository
}

func New(persist items.PersistRepository) items.Repository {
	return &itemRepository{
		persist: persist,
	}
}
func (ir *itemRepository) GetById(id int) items.Item {
	res := ir.persist.GetById(id)
	return res
}

func (ir *itemRepository) Get() []items.Item {
	res := ir.persist.Get()
	return res
}

func (ir *itemRepository) Store(item items.Item) items.Item {
	res := ir.persist.Insert(item)
	return res
}

func (ir *itemRepository) Update(id int, item items.Item) items.Item {
	res := ir.persist.Update(id, item)
	return res

}

func (ir *itemRepository) Delete(id int) error {
	res := ir.persist.Delete(id)
	return res
}
