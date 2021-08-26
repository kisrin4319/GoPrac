package Stack

import "fmt"

// Eval returns the evaluation result of the given expr.
// The expression can have +, -, *, /, (, ) operators and
// decimal integers. Operators and operands should be
// space delimited

func ExampleEval() {
	fmt.Println(Eval("5"))
	fmt.Println(Eval("1 + 2"))
	fmt.Println(Eval("1 - 2 + 3"))
	fmt.Println(Eval("3 * ( 3 + 1 * 3 ) / 2"))
	fmt.Println(Eval("3 * ( ( 3 + 1 ) * 3 ) / 2"))
	//OutPut:
	//5
	//3
	//2
	//9
	//18
}
