package service3

import "context"

// Repository ...
type Repository interface {
	RunInTransaction(context.Context, func(context.Context) (interface{}, error)) (interface{}, error)
	OpenAccount(ctx context.Context, initialAmmount int) (Account, error)
	GetAccountsForTransfer(ctx context.Context, fromID, toID int64) (from, to Account, err error)
	UpdateBalance(ctx context.Context, a Account) (Account, error)
}
