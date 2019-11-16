package service2

// Person ...
type Person struct {
	ID    int64  `db:"kokoha"`
	Name  string `db:"tekitode"` // sql.NullString はインフラに結合するので使わない
	Email string `db:"yoiyo"`
}
