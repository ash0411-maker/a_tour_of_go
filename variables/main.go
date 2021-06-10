// var ステートメントは変数( variable )を宣言します。
// 関数の引数リストと同様に、複数の変数の最後に型を書くことで、変数のリストを宣言できます。
// package main

// import "fmt"

// var c, python, java bool

// func main() {
// 	var i int
// 	fmt.Println(i, c, python, java)
// }


// ===================================================================
// Variables with initializers
// var 宣言では、変数毎に初期化子( initializer )を与えることができます。
// 初期化子(型の宣言の省略)
// また、変数の宣言時に初期化子を与えることもできます。
// 初期化子を与える場合、型の宣言を省略することができます。
// この場合、変数の型は初期化子が持つ型になります。

// package main

// import "fmt"

// var i, j int = 1, 2

// func main() {
// 	var c, python, java = true, false, "no!"
// 	fmt.Println(i, j, c, python, java)
// }



// ===================================================================
// Short variable declarations
// 関数の中では、 var 宣言の代わりに、短い := の代入文を使い、暗黙的な型宣言ができます。
// なお、関数の外では、キーワードではじまる宣言( var, func, など)が必要で、 := での暗黙的な宣言は利用できません。

// package main

// import "fmt"

// func main() {
// 	var i, j int = 1, 2
// 	k := 3
// 	c, python, java := true, false, "no!"

// 	fmt.Println(i, j, k, c, python, java)
// }


// ===================================================================
// Zero values
// 変数に初期値を与えずに宣言すると、ゼロ値( zero value )が与えられます。

// ゼロ値は型によって以下のように与えられます:

// 数値型(int,floatなど): 0
// bool型: false
// string型: "" (空文字列( empty string ))

// package main

// import "fmt"

// func main() {
// 	var i int
// 	var f float64
// 	var b bool
// 	var s string
// 	fmt.Printf("%v %v %v %q\n", i, f, b, s)
// }



// ===================================================================
// Type inference
// 明示的な型を指定せずに変数を宣言する場合( := や var = のいずれか)、変数の型は右側の変数から型推論されます。

// 右側の変数が型を持っている場合、左側の新しい変数は同じ型になります:

// package main
// import "fmt"
// func main() {
// 	v := "42" // change me!
// 	fmt.Printf("v is of type %T\n", v)
// }



// ===================================================================
// Constants
// 定数( constant )は、 const キーワードを使って変数と同じように宣言します。

// 定数は、文字(character)、文字列(string)、boolean、数値(numeric)のみで使えます。

// なお、定数は := を使って宣言できません。
package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
