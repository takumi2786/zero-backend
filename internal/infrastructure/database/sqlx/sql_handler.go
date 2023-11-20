package sqlx

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/takumi2786/zero-backend/internal/interfaces/repository"
	"github.com/takumi2786/zero-backend/internal/util"
)

var retryCount = 0

const maxRetryCount = 10

func connect(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		if retryCount < maxRetryCount {
			time.Sleep(time.Second)
			retryCount++
			return connect(dsn)
		} else {
			return nil, err
		}
	}
	db.SetMaxIdleConns(30)
	db.SetMaxOpenConns(10)
	return db, nil
}

func NewDB(cfg *util.Config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)
	return connect(dsn)
}

type SQLHandler struct {
	db *sqlx.DB
}

func NewSQLHandler(db *sqlx.DB) repository.SQLHandler {
	return &SQLHandler{db: db}
}

// Get using this DB.
// Any placeholder parameters are replaced with supplied args.
// An error is returned if the result set is empty.
func (h *SQLHandler) Get(dest interface{}, query string, args ...interface{}) error {
	return h.db.Get(dest, query, args...)
}

// TODO: 他のメソッドも必要に応じて実装する
