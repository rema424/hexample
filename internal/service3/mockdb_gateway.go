package service3

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type ctxKey string

const txCtxKey ctxKey = "transaction"

var src = rand.NewSource(time.Now().UnixNano())

// MockGateway ...
type MockGateway struct {
	db *MockDB
}

// NewMockGateway ...
func NewMockGateway(db *MockDB) *MockGateway {
	return &MockGateway{db}
}

// MockDB ...
type MockDB struct {
	mu   sync.RWMutex
	data map[int64]Account
}

// NewMockDB ...
func NewMockDB() *MockDB {
	return &MockDB{data: make(map[int64]Account)}
}

// OpenAccount ...
func (g *MockGateway) OpenAccount(ctx context.Context, initialAmmount int) (Account, error) {
	if !isInTx(ctx) {
		g.db.mu.Lock()
		defer g.db.mu.Unlock()
	}

	// 割り当て可能なIDを探す
	var id int64
	for {
		id = src.Int63()
		_, ok := g.db.data[id]
		if !ok {
			break
		}
	}

	a := Account{ID: id, Balance: initialAmmount}
	g.db.data[id] = a
	return a, nil
}

// IsBalanceSufficient ...
func (g *MockGateway) IsBalanceSufficient(ctx context.Context, accountID int64, ammount int) (bool, error) {
	if !isInTx(ctx) {
		g.db.mu.Lock()
		defer g.db.mu.Unlock()
	}

	account, ok := g.db.data[accountID]
	if !ok {
		return false, fmt.Errorf("gateway: account not found - id: %d", accountID)
	}

	return account.Balance >= ammount, nil
}

func (g *MockGateway) getAccountByID(ctx context.Context, accountID int64) (Account, error) {
	if !isInTx(ctx) {
		g.db.mu.Lock()
		defer g.db.mu.Unlock()
	}

	account, ok := g.db.data[accountID]
	if !ok {
		return Account{}, fmt.Errorf("gateway: account not found - id: %d", accountID)
	}

	return account, nil
}

// DecreaseBalance ...
func (g *MockGateway) DecreaseBalance(ctx context.Context, id int64, ammount int) (Account, error) {
	if !isInTx(ctx) {
		g.db.mu.Lock()
		defer g.db.mu.Unlock()
	}

	account, err := g.getAccountByID(ctx, id)
	if err != nil {
		return Account{}, err
	}
	account.Balance -= ammount
	g.db.data[id] = account
	return account, nil
}

// IncreaseBalance ...
func (g *MockGateway) IncreaseBalance(ctx context.Context, id int64, ammount int) (Account, error) {
	if !isInTx(ctx) {
		g.db.mu.Lock()
		defer g.db.mu.Unlock()
	}

	account, err := g.getAccountByID(ctx, id)
	if err != nil {
		return Account{}, err
	}
	account.Balance += ammount
	g.db.data[id] = account
	return account, nil
}

// RunInTransaction ...
func (g *MockGateway) RunInTransaction(ctx context.Context, txFn TxFn) (interface{}, error) {
	// 多重トランザクションはエラーとする
	if isInTx(ctx) {
		return nil, fmt.Errorf("gateway: detect nested transaction")
	}

	// context をデコレートして transaction context を生成する
	txCtx := genTxCtx(ctx)

	// ロックを取得する（ロックの取得から解放までの間がトランザクションとなる）
	g.db.mu.Lock()
	defer g.db.mu.Unlock()

	// transaction 処理を実行する
	return txFn(txCtx)
}

func isInTx(ctx context.Context) bool {
	if val, ok := ctx.Value(txCtxKey).(bool); ok {
		return val
	}
	return false
}

func genTxCtx(ctx context.Context) context.Context {
	if isInTx(ctx) {
		return ctx
	}
	return context.WithValue(ctx, txCtxKey, true)
}
