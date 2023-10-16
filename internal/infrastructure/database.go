package infrastructure

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
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

func NewDB(ctx context.Context, cfg *util.Config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)
	return connect(dsn)
}

/*
	sqlx.DBのメソッドのうち、アプリケーションで利用する機能のみをインターフェースとして定義する
	これらは使わなくても良いかも
*/
// トランザクションメソッド系類
type Beginner interface {
	BeginTxx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error)
}

// 比較系メソッド類
type Preparer interface {
	PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error)
}

// 書き込み系メソッド類
type Execer interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
}

func GetExecer(db *sqlx.DB) Execer {
	return db
}

// 読み取り系メソッド類
type Queryer interface {
	Preparer
	QueryxContext(ctx context.Context, query string, args ...any) (*sqlx.Rows, error)
	QueryRowxContext(ctx context.Context, query string, args ...any) *sqlx.Row
	GetContext(ctx context.Context, dest interface{}, query string, args ...any) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...any) error
}

func GetQueryer(db *sqlx.DB) Queryer {
	return db
}

// トランザクション内で扱うメソッド群
type Transacter interface {
	DB() *sqlx.Tx
	Begin(ctx context.Context) error
	Commit() error
	Rollback() error
}

var (
	// インターフェースが期待通りに宣言されているか確認
	_ Beginner = (*sqlx.DB)(nil)
	_ Preparer = (*sqlx.DB)(nil)
	_ Queryer  = (*sqlx.DB)(nil)
	_ Execer   = (*sqlx.DB)(nil)
	_ Execer   = (*sqlx.Tx)(nil)
)
