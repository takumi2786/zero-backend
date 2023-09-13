package main

import "fmt"

// インターフェースの定義
type Greeter interface {
	Greet(message string)
}

// インターフェースを実装した構造体
type Japanese struct{}

func (j *Japanese) Greet(message string) {
	fmt.Println("こんにちは, ", message)
}

// インターフェースを実装した構造体
type American struct{}

func (a *American) Greet(message string) {
	fmt.Println("Hello, ", message)
}

func greetMulti(greeters []Greeter, message string) {
	for _, greeter := range greeters {
		greeter.Greet(message)
	}
}

// インターフェースの実装をチェック: https://zenn.dev/seri_k/articles/df3052543f1c62
// nilを struct Japaneseのポインタにキャストして Greeter 型の変数へのキャストをコンパイル時に実行することで，struct A がinterface Iを実装しているかどうかをコンパイル時にチェックできる．
var _ Greeter = (*Japanese)(nil)
var _ Greeter = (*American)(nil)

func main() {
	japanese := &Japanese{}
	// japanese.Greet("世界")

	american := &American{}
	// american.Greet("world")

	greetMulti([]Greeter{japanese, american}, "Jon")
}
