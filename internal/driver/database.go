package driver

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

func NewDB(ctx context.Context, cfg *Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", "test:test@tcp(localhost:3306)/zero_system?charset=utf8mb4")
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(30)
	db.SetMaxOpenConns(10)
	return db, nil
}

/*
	sqlx.DBのメソッドのうち、アプリケーションで利用する機能のみをインターフェースとして定義する
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
