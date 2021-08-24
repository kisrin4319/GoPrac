package main

func main() {

	//1. 익명함수
	sum := func(n ...int) int { // 익명함수 정의
		s := 0
		for _, i := range n {
			s += i
		}
		return s
	}

	result := sum(1, 2, 3, 4, 5) // 익명함수 호출
	println(result)

	//2.일급함수

	//변수 add 에 익명함수 할당
	add := func(i int, j int) int {
		return i + j
	}

	// add 함수 전달
	r1 := calc(add, 10, 20)
	println(r1)

	// 직접 첫번째 파라미터에 익명함수를 정의함
	r2 := calc(func(x int, y int) int { return x - y }, 10, 20)
	println(r2)

}

func calc(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}

//3. type문을 사용한 함수 원형 정의

//원형 정의
type calculator func(int, int) int

// calculator 원형 사용
func calc2(f calculator, a int, b int) int {
	result := f(a, b)
	return result
}
