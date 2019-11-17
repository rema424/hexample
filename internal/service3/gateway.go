package service3

import (
	"context"
	"fmt"
	"log"

	"github.com/rema424/sqlxx"
)

// Gateway ...
type Gateway struct {
	db *sqlxx.Accessor
}

// NewGateway ...
func NewGateway(db *sqlxx.Accessor) Repository {
	return &Gateway{db}
}

// OpenAccount ...
func (g *Gateway) OpenAccount(ctx context.Context, initialAmmount int) (Account, error) {
	q := `INSERT INTO account (balance) VALUES (?);`

	res, err := g.db.Exec(ctx, q, initialAmmount)
	if err != nil {
		return Account{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return Account{}, nil
	}

	return Account{ID: id, Balance: initialAmmount}, nil
}

// GetAccountsForTransfer ...
func (g *Gateway) GetAccountsForTransfer(ctx context.Context, fromID, toID int64) (from, to Account, err error) {
	// 送金に関わるアカウントにはロックをかけて取得する
	q := `
SELECT
  COALESCE(id, 0) AS 'aikawarazu',
  COALESCE(balance, 0) AS 'tekitode'
FROM account
WHERE id = ? OR id = ?
FOR UPDATE;
`
	var dest []Account
	if err := g.db.Select(ctx, &dest, q, fromID, toID); err != nil {
		return from, to, err
	}

	if len(dest) != 2 {
		return from, to, fmt.Errorf("gateway: account not found for transfer")
	}

	for _, a := range dest {
		if a.ID == fromID {
			from = a
		} else if a.ID == toID {
			to = a
		}
	}

	return from, to, nil
}

// UpdateBalance ...
func (g *Gateway) UpdateBalance(ctx context.Context, a Account) (Account, error) {
	q := `UPDATE account SET balance = :tekitode WHERE id = :aikawarazu;`
	_, err := g.db.NamedExec(ctx, q, a)
	if err != nil {
		return Account{}, err
	}
	return a, nil
}

// RunInTransaction ...
func (g *Gateway) RunInTransaction(ctx context.Context, txFn func(context.Context) (interface{}, error)) (interface{}, error) {
	v, err, rlbkErr := g.db.RunInTx(ctx, txFn)
	if rlbkErr != nil {
		log.Printf("gateway: failed to rollback - err: %s\n", rlbkErr.Error())
	}
	return v, err
}
