package items

type PersistRepository interface {
	GetById(id int) Item
	Get() []Item
	Insert(item Item) Item
	Update(id int, item Item) Item
	Delete(id int) error
}
