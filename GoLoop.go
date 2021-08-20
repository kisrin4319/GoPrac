package main

func main() {

	//1. for문
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	println(sum)

	//2. for문 - 조건식만 쓰는 for 루프
	n := 1
	for n < 100 {
		n *= 2
		//if n > 90 {
		//break
		//}
		println(n)
	}

	//3. for 문 - 무한루프

	/*
		for {
			println("Infinite loop")
		}*/

	//4. for range 문
	//for 인덱스, 요소값:= range 컬렉션

	//아래 코드는 Collection 타입 관련
	names := []string{"홍길동", "이순신", "강감찬"}
	for index, name := range names {
		println(index, name)
	}

	//5. break, continue, goto 문
	a := 0
	for a < 15 {
		if a == 5 {
			a += a
			continue // for루프 시작으로
		}
		a++
		if a > 10 {
			break
		}
	}
	if a == 11 {
		goto END //goto 사용 ㅇ{
	}
	println(a)
END:
	println("End")

}
