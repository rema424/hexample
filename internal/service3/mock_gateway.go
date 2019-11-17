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
func NewMockGateway(db *MockDB) Repository {
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

// GetAccountsForTransfer ...
func (g *MockGateway) GetAccountsForTransfer(ctx context.Context, fromID, toID int64) (from, to Account, err error) {
	if !isInTx(ctx) {
		g.db.mu.Lock()
		defer g.db.mu.Unlock()
	}

	var ok bool
	from, ok = g.db.data[fromID]
	if !ok {
		return Account{}, Account{}, fmt.Errorf("gateway: account not found - accoutID: %d", fromID)
	}

	to, ok = g.db.data[toID]
	if !ok {
		return Account{}, Account{}, fmt.Errorf("gateway: account not found - accoutID: %d", toID)
	}

	return from, to, nil
}

// UpdateBalance ...
func (g *MockGateway) UpdateBalance(ctx context.Context, a Account) (Account, error) {
	if !isInTx(ctx) {
		g.db.mu.Lock()
		defer g.db.mu.Unlock()
	}

	g.db.data[a.ID] = a
	return a, nil
}

// RunInTransaction ...
func (g *MockGateway) RunInTransaction(ctx context.Context, txFn func(context.Context) (interface{}, error)) (interface{}, error) {
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
