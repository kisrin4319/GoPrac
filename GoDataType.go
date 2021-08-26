package main

import (
	"fmt"
)

func main() {

	/*	 Go 데이타 타입
		 1. 부울린 타입
		 bool
		 2. 문자열 타입
		 string : string은 한번 생성되면 수정될 수 없는 Immutable 타입
		 3. 정수형 타입
		 int , int8 , int16 , int32 , int 64
		 uint , uint8, uint16, uint32, uint64 , uintptr // uinitptr은 uint와 크기가 동일하며 포인터를 저장할 때 사용
		 4. Float 및 복소수 타입
		 float32 , float64, complex, complex64 , complex128 // complex 타입은 float 크기의 실수부와 허수부로 된 복소수..
		 5. 기타 타입
		 byte : uint8과 동일하며 바이트 코드에 사용
		 rune : int32과 동일하여 유니코드 코드포인트에 사용한다.
	*/
	//Raw String Literal. 복수라인 가능
	rawLiteral := `아리랑\n
					아리랑\n
  					아라리요`
	// Interpreted String Literal
	interLiteral := "아리랑아리랑\n아라리요"
	// 아래와 같이 +를 사용하여 두 라인에 걸쳐 사용할 수도 있다.
	// interLiteral := "아리랑아리랑\n" +
	//                 "아라리요"

	fmt.Println(rawLiteral)
	fmt.Println()
	fmt.Println(interLiteral)

	//3.데이터 타입 변환 (Type Conversion)
	/*var i int = 100
	var u uint = uint(i)
	var f float32 = float32(i)
	println(f, u)*/

	str := "ABC"
	bytes := []byte(str)
	str2 := string(bytes)
	println(bytes, str2)

	//	var i int
	//	var k int64
	//	var f float64
	//	var s string
	//	var err error
	//	i, err = strconv.Atoi("350") //i == 350
	//	k, err = strconv.ParseInt**("cc7fdd",16,32) // k == 13402077
	//	k, err = strconv.ParseInt("0xcc7fdd",0,32) //k == 13402077
	//	f, err = strconv.ParseFloat***("3.14",64) // f == 3.14
	//	s = strconv.Itoa****(340)	// s == "340"
	//	s = strconv.FormatInt*****(13402077,16)

}
