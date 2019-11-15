package purchase

// Transaction .
type Transaction struct {
	ID int64
	Order
	Buyer
	Seller
}

// Order .
type Order struct {
	ID      int64
	Details []OrderDetail
}

// OrderDetail .
type OrderDetail struct {
	ID       int64
	SubTotal int
	Quantity int
	Product
}

// Product .
type Product struct {
	ID    int64
	Name  string
	Price int
}
