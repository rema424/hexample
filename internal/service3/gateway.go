package service3

import (
	"context"
	"log"

	"github.com/rema424/sqlxx"
)

// Gateway ...
type Gateway struct {
	db *sqlxx.Accessor
}

// NewGateway ...
func NewGateway(db *sqlxx.Accessor) *Gateway {
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

// IsBalanceSufficient ...
func (g *Gateway) IsBalanceSufficient(ctx context.Context, accountID int64, ammount int) (bool, error) {
	// 行ロックを取得する（FOR UPDATE）
	q := `SELECT COALESCE(balance, 0) AS 'tekitode' FROM account WHERE id = ? FOR UPDATE;`

	// 残高を取得
	var b int
	if err := g.db.Get(ctx, &b, q, accountID); err != nil {
		return false, err
	}

	// 残高を確認
	return b >= ammount, nil
}

// DecreaseBalance ...
func (g *Gateway) DecreaseBalance(ctx context.Context, id int64, ammount int) (Account, error) {
	q := `UPDATE account SET balance = balance - ? WHERE id = ?;`
	_, err := g.db.Exec(ctx, q, ammount, id)
	if err != nil {
		return Account{}, err
	}

	return g.getAccountByID(ctx, id)
}

// IncreaseBalance ...
func (g *Gateway) IncreaseBalance(ctx context.Context, id int64, ammount int) (Account, error) {
	q := `UPDATE account SET balance = balance + ? WHERE id = ?;`
	_, err := g.db.Exec(ctx, q, ammount, id)
	if err != nil {
		return Account{}, err
	}

	return g.getAccountByID(ctx, id)
}

func (g *Gateway) getAccountByID(ctx context.Context, id int64) (Account, error) {
	q := `
SELECT
  COALESCE(id, 0) AS 'aikawarazu',
  COALESCE(balance, 0) AS 'tekitode'
FROM account
WHERE id = ?;
`
	var a Account
	if err := g.db.Get(ctx, &a, q, id); err != nil {
		return a, err
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
