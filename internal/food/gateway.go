package food

import "github.com/rema424/sqlxx"

// Gateway .
type Gateway struct {
	acsr *sqlxx.Accessor
}

// NewGateway .
func NewGateway(acsr *sqlxx.Accessor) Repository {
	return &Gateway{acsr}
}
