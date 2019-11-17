package service3

import (
	"context"
	"fmt"
)

// Provider ...
type Provider struct {
	r Repository
}

// NewProvider ...
func NewProvider(r Repository) *Provider {
	return &Provider{r}
}

// OpenAccount ...
func (p *Provider) OpenAccount(ctx context.Context, initialAmmount int) (Account, error) {
	if initialAmmount <= 0 {
		return Account{}, fmt.Errorf("provider: initial ammount must be greater than 0")
	}

	account, err := p.r.OpenAccount(ctx, initialAmmount)
	if err != nil {
		return Account{}, err
	}
	return account, nil
}

// Transfer ...
func (p *Provider) Transfer(ctx context.Context, ammount int, fromID, toID int64) (from, to Account, err error) {
	if fromID == toID {
		return Account{}, Account{}, fmt.Errorf("provider: cannot transfer money to oneself")
	}

	type Accounts struct {
		from Account
		to   Account
	}

	txFn := func(ctx context.Context) (interface{}, error) {
		// 送金元の残高を確認
		b, err := p.r.IsBalanceSufficient(ctx, fromID, ammount)
		if err != nil {
			return Accounts{}, err
		} else if !b {
			return Accounts{}, fmt.Errorf("provider: balance is not sufficient - accountID: %d", fromID)
		}

		// 送金元の残高を減らす
		from, err := p.r.DecreaseBalance(ctx, fromID, ammount)
		if err != nil {
			return Accounts{}, err
		}

		// 送金先の残高を増やす
		to, err := p.r.IncreaseBalance(ctx, toID, ammount)
		if err != nil {
			return Accounts{}, err
		}

		return Accounts{from: from, to: to}, nil
	}

	v, err := p.r.RunInTransaction(ctx, txFn)
	if err != nil {
		return Account{}, Account{}, err
	}

	val, ok := v.(Accounts)
	if !ok {
		return Account{}, Account{}, fmt.Errorf("provider: an error occurs - transfer")
	}

	return val.from, val.to, nil
}
