package main

import "fmt"

func main() {

	var a []int        //슬라이스 변수 선언
	a = []int{1, 2, 3} // 슬라이스에 리터럴값 지정
	a[1] = 10
	fmt.Println(a) // [1,10,3]

	s := make([]int, 5, 10)
	println(len(s), cap(s)) // len 5, cap 10

	//닐 슬라이스 (Nil Slice)
	var s1 []int
	if s1 == nil {
		println("Nil Slice")
	}
	print(len(s1), cap(s1)) // 모두 0

	//2. 부분 슬라이스(Sub-Slice)
	s2 := []int{0, 1, 2, 3, 4, 5}
	s2 = s2[2:5]
	fmt.Println(s2) // 2,3,4 출력

	//3. 슬라이스 추가, 병합(append)과 복사(Copy)

	s3 := []int{0, 1}

	//하나 확장
	s3 = append(s3, 2) // 0,1,2
	// 복수 요소들 확장
	s3 = append(s3, 3, 4, 5) // 0,1,2,3,4,5

	fmt.Println(s3)

	//len = 0 , cap = 3 인 슬라이스
	sliceA := make([]int, 0, 3)

	// 계속 한 요소씩 추가
	for i := 1; i <= 15; i++ {
		sliceA = append(sliceA, i)
		// 슬라이스 길이와 용량 확인
		fmt.Println(len(sliceA), cap(sliceA))
	}

	fmt.Println(sliceA)

	sliceB := []int{1, 2, 3}
	sliceC := []int{4, 5, 6}

	sliceB = append(sliceB, sliceC...)

	fmt.Println(sliceB)

	source := []int{0, 1, 2}
	target := make([]int, len(source), cap(source)*2)
	copy(target, source)
	fmt.Println(target)               //[0,1,2] 출력
	println(len(target), cap(target)) // 3,6
}
