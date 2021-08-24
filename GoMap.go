package main

import "fmt"

func main() {

	var idMap map[int]string // 선언 map[key타입]Value타입

	//make 함수의 첫번째 파라미터 map키워드와 [키타입]값타입을 지정
	idMap = make(map[int]string) // map 초기화

	println(idMap)

	var m map[int]string

	m = make(map[int]string)

	m[901] = "Apple"
	m[134] = "Grape"
	m[777] = "Tomato"

	//키에 대한 값 읽기
	str := m[134]
	println(str)

	noData := m[999] // 값이 없으면 nil 혹은 zero 리턴
	println(noData)

	//삭제
	delete(m, 777)

	//Map 키 체크
	//리터럴을 사용한 초기화
	tickers := map[string]string{
		"GOOG": "Google Inc",
		"MSFT": "Microsoft",
		"FB":   "FaceBook",
		"AMZN": "Amazon",
	}

	//map 키 체크
	val, exists := tickers["MSFT"]
	if !exists {
		print("No MSFT ticker")
	} else {
		print(val)
	}

	myMap := map[string]string{
		"A": "Apple",
		"B": "Banana",
		"C": "Charlie",
	}

	// for range 문을 사용하여 모든 맵 요소 출력
	// map은 unordered 이므로 순서는 무작위
	for key, val := range myMap {
		fmt.Println(key, val)
	}
}
