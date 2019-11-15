package catalog

// Product .
type Product struct {
	ID       int64
	Name     string
	Price    int
	ImageURL string
	Category Category
	Maker    Maker
}

// Category .
type Category struct {
	ID   int64
	Name string
}

// Maker .
type Maker struct {
	ID   int64
	Name string
}
