package usecase

import (
	"github.com/oryzel/go-item-point/src/items"
)

type itemUsecase struct {
	ItemRepository items.Repository
}

func New(repository items.Repository) items.Usecase {
	return &itemUsecase{
		ItemRepository: repository,
	}
}

func (iu itemUsecase) GetItem(id int) items.Item {
	res := iu.ItemRepository.GetById(id)
	return res
}

func (iu itemUsecase) GetItems() []items.Item {
	res := iu.ItemRepository.Get()
	return res
}

func (iu itemUsecase) StoreItem(item items.Item) items.Item {
	res := iu.ItemRepository.Store(item)
	return res
}

func (iu itemUsecase) UpdateItem(id int, item items.Item) items.Item {
	res := iu.ItemRepository.Update(id, item)
	return res
}

func (iu itemUsecase) DeleteItem(id int) error {
	res := iu.ItemRepository.Delete(id)
	return res
}
