package items

type Usecase interface {
	GetItem(id int) Item
	GetItems() []Item
	StoreItem(item Item) Item
	UpdateItem(id int, item Item) Item
	DeleteItem(id int) error
}
