package service2

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var src = rand.NewSource(time.Now().UnixNano())

// MockGateway ...
type MockGateway struct {
	db *MockDB
}

// MockDB ...
type MockDB struct {
	mu   sync.RWMutex
	data map[int64]Person
}

// NewMockDB ...s
func NewMockDB() *MockDB {
	return &MockDB{data: make(map[int64]Person)}
}

// NewMockGateway ...
func NewMockGateway(db *MockDB) Repository {
	return &MockGateway{db}
}

// RegisterPerson ...
func (r *MockGateway) RegisterPerson(ctx context.Context, p Person) (Person, error) {
	r.db.mu.Lock()
	defer r.db.mu.Unlock()

	// 割り当て可能なIDを探す
	var id int64
	for {
		id = src.Int63()
		_, ok := r.db.data[id]
		if !ok {
			break
		}
	}

	p.ID = id
	r.db.data[p.ID] = p

	return p, nil
}

// GetPersonByID ...
func (r *MockGateway) GetPersonByID(ctx context.Context, id int64) (Person, error) {
	r.db.mu.RLock()
	defer r.db.mu.RUnlock()

	if p, ok := r.db.data[id]; ok {
		return p, nil
	}
	return Person{}, fmt.Errorf("person not found - id: %d", id)
}
