package items

type Item struct {
	Id        int
	Name      string
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
}
