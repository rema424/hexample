package inventory

// Status .
type Status int

const (
	// OnSale .
	OnSale Status = 1
	// Reserved .
	Reserved Status = 2
	// SoldOut .
	SoldOut Status = 3
)

// Stock .
type Stock struct {
	ID      int64
	Status  Status
	Product Product
}

// Product .
type Product struct {
	ID     int64
	Name   string
	Price  int
	Stocks []Stock
}
