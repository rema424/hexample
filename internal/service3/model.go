package service3

// Account ...
type Account struct {
	ID      int64 `db:"aikawarazu"`
	Balance int   `db:"tekitode"`
}

// IsSufficient ...
func (a *Account) IsSufficient(ammount int) bool {
	return a.Balance >= ammount
}

// Transfer ...
func (a *Account) Transfer(ammount int, to *Account) {
	a.Balance -= ammount
	to.Balance += ammount
}
