package mysql

import (
	"time"

	my "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Config ...
type Config struct {
	User                 string
	Passwd               string
	Host                 string
	Port                 string
	Net                  string
	Addr                 string
	DBName               string
	Collation            string
	InterpolateParams    bool
	AllowNativePasswords bool
	ParseTime            bool
	MaxOpenConns         int
	MaxIdleConns         int
	ConnMaxLifetime      time.Duration
}

func (c Config) build() Config {
	// var (
	// 	host     = os.Getenv("DB_HOST")
	// 	port     = os.Getenv("DB_PORT")
	// 	user     = os.Getenv("DB_USER")
	// 	dbName   = os.Getenv("DB_NAME")
	// 	password = os.Getenv("DB_PASSWORD")
	// )
	if c.User == "" {
		c.User = "root"
	}
	if c.Net == "" {
		c.Net = "tcp"
	}
	if c.Host == "" {
		c.Host = "127.0.0.1"
	}
	if c.Port == "" {
		c.Port = "3306"
	}
	if c.Addr == "" {
		c.Addr = c.Host + ":" + c.Port
	}
	if c.Collation == "" {
		c.Collation = "utf8mb4_bin"
	}
	if c.MaxOpenConns < 0 {
		c.MaxOpenConns = 30
	}
	if c.MaxIdleConns < 0 {
		c.MaxIdleConns = 30
	}
	if c.ConnMaxLifetime < 0 {
		c.ConnMaxLifetime = 60 * time.Second
	}
	return c
}

// Connect .
func Connect(c Config) (*sqlx.DB, error) {
	c = c.build()

	mycfg := my.Config{
		User:                 c.User,
		Passwd:               c.Passwd,
		Net:                  c.Net,
		Addr:                 c.Addr,
		DBName:               c.DBName,
		Collation:            c.Collation,
		InterpolateParams:    c.InterpolateParams,
		AllowNativePasswords: c.AllowNativePasswords,
		ParseTime:            c.ParseTime,
	}

	dbx, err := sqlx.Connect("mysql", mycfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	dbx.SetMaxOpenConns(c.MaxOpenConns)
	dbx.SetMaxIdleConns(c.MaxIdleConns)
	dbx.SetConnMaxLifetime(c.ConnMaxLifetime)

	return dbx, nil
}
