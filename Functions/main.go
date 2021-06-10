
// Functions
// 関数は、0個以上の引数を取ることができます。
// この例では、 add 関数は、 int 型の２つのパラメータを取ります。
// 変数名の 後ろ に型名を書くことに注意してください。
// https://blog.golang.org/declaration-syntax


// package main
// import "fmt"

// func add(x int, y int) int {
// 	return x + y
// }

// func main() {
// 	fmt.Println(add(42, 13))
// }


// ===================================================================
// 関数の２つ以上の引数が同じ型である場合には、最後の型を残して省略して記述できます。

// package main

// import "fmt"

// func add(x, y int) int {
// 	return x + y
// }

// func main() {
// 	fmt.Println(add(42, 13))
// }

// ===================================================================
// 関数は複数の戻り値を返すことができます。

// package main
// import "fmt"

// func swap(x, y string) (string, string) {
// 	return y, x
// }

// func main() {
// 	a, b := swap("hello", "world")
// 	fmt.Println(a, b)
// }


// ===================================================================
// Named return values
// Goでの戻り値となる変数に名前をつける( named return value )ことができます。戻り値に名前をつけると、
// 関数の最初で定義した変数名として扱われます。
// この戻り値の名前は、戻り値の意味を示す名前とすることで、関数のドキュメントとして表現するようにしましょう。
// 名前をつけた戻り値の変数を使うと、 return ステートメントに何も書かずに戻すことができます。これを "naked" return と呼びます。

package main
import "fmt"

func split(sum int) (w, v int) {
	w = sum * 4 / 9
	v = sum - w
	return
}

func main() {
	fmt.Println(split(17))
}
