package items

type Repository interface {
	GetById(id int) Item
	Get() []Item
	Store(item Item) Item
	Update(id int, item Item) Item
	Delete(id int) error
}
