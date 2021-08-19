package main

import "fmt"

func main() {

	var k int

	// if
	if k == 1 { // 같은 라인
		println("One")
	} else if k == 2 {
		println("Two")
	} else {
		println("other")
	}

	//Switch

	var name string
	var category = 1

	// Go만의 Switch 특징
	// 1. switch 뒤에 expression이 없을 수 있음
	// 2. case문에 expression을 쓸 수 있음
	// 3. No default fall through ( 다른 언어의 case문은 break를 쓰지 않는 한 다음 case로 이동하지만 Go는 다음 case로 가지 않는다.
	// 4. Type Switch
	switch category {
	case 1:
		name = "Paper Book"
	case 2:
		name = "eBook"
	case 3, 4:
		name = "Blog"
	default:
		name = "Other"
	}
	println(name)

	var score int = 85
	grade(score)

	/*var v int = 100

	switch v.(type) {

	case int:
		println("int")
	case bool:
		println("bool")
	case string:
		println("string")
	default:
		println("unknown")
	}*/

	check(2)

}

func check(val int) {
	switch val {
	case 1:
		fmt.Println("1 이하")
		fallthrough
	case 2:
		fmt.Println("2 이하")
		fallthrough
	case 3:
		fmt.Println("3 이하")
		fallthrough
	default:
		fmt.Println("default 도달")
	}
}
func grade(score int) {
	switch {
	case score >= 90:
		println("A")
	case score >= 80:
		println("B")
	case score >= 70:
		println("C")
	case score >= 60:
		println("D")
	default:
		println("No Hope")
	}
}
