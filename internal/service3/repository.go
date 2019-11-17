package service3

import "context"

// Repository ...
type Repository interface {
	RunInTransaction(context.Context, func(context.Context) (interface{}, error)) (interface{}, error)
	OpenAccount(ctx context.Context, initialAmmount int) (Account, error)
	IsBalanceSufficient(ctx context.Context, accountID int64, ammount int) (bool, error)
	IncreaseBalance(ctx context.Context, accountID int64, ammount int) (Account, error)
	DecreaseBalance(ctx context.Context, accountID int64, ammount int) (Account, error)
}
