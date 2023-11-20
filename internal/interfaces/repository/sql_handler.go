package repository

// SQLHandlerは、DBへのSQL実行処理を抽象化するインターフェースです。
// infrastrucure/database/sqlx/database.goで実装しています。
type SQLHandler interface {
	// Getは、クエリ結果をdest構造体に格納します。
	Get(dest interface{}, query string, args ...interface{}) error
}
