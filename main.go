package main

import "fmt"

func main() {
	fmt.Print("Hello World!")

	// 변수는 Go 키워드 var를 사용하여 선언한다.
	// var 키워드 뒤에 변수명을 적고, 그 뒤에 변수타입을 적는다.

	var a int

	// 변수 초기값을 할당할 수 도 있다. // float32 타입의 변수 f에 11.0이라는 초기값을 할당
	var f float32 = 11.

	//만약 선언된 변수가 Go 프로그램 내에서 사용되지 않는다면, 에러를 발생시킨다.

	// 상수는 Go 키워드 const를 사용하여 선언한다.
	const c int = 10
	// 아래와 같이 초기 값을 지정할 경우 타입을 추론하는 기능이 자주 사용 된다.
	const s = "Hi"

	// 여러개의 상수들 묶어서 지정할 수 있는데 , 아래와 같이 괄호 안에 상수들을 나열하여 사용할 수 있다.
	const ( // 상수값을 0부터 순차적으로 부여하기 위해 iota라는 identifier를 사용할 수 있다.
		Visa   = "Visa"             // = itoa // 0
		Master = "MasterCard"       // = itoa // 1
		Amex   = "American Express" // = itoa // 2
	)
}
