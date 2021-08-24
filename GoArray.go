package main

import "fmt"

func main() {

	var a [3]int // 정수형 3개 요소를 가지는 배열 a 선언

	fmt.Println(a[2])

	//2.배열의 초기화
	var a1 = [3]int{1, 2, 3}
	var a3 = [...]int{1, 2, 3} // 배열 크기 자동으로

	fmt.Println(a1[1])
	fmt.Println(a3[1])

	//3. 다배열 배열
	var multiArray [3][4][5]int // 정의
	multiArray[0][1][2] = 10    // 사용용

	//4. 다차원 배열의 초기화
	var multiArray2 = [2][3]int{
		{1, 2, 3},
		{4, 5, 6}, // 끝에 컴마 추가
	}
	println(multiArray2[1][2])

}
