package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

//Rect 정의
type Rect struct {
	width, height float64
}

// Circle 정의
type Circle struct {
	radius float64
}

func Marshal(v interface{}) ([]byte, error)
func Println(a ...interface{}) (n int, err error)

//Rect 타입에 대한 Shape 인터페이스 구현
func (r Rect) area() float64 {
	return r.width * r.height
}
func (r Rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

//Circle 타입에 대한 Shape 인퍼페이스 구현
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {

	r := Rect{10., 20.}
	c := Circle{10}

	showArea(r, c)

	var x interface{}
	x = 1
	x = "Tom"

	printIt(x)

	var a interface{} = 1

	i := a       // a와 i는 dynamic type, 값은 1
	j := a.(int) //j는 ㅌint 타입 , 값은 1

	println(i) //포인터주소 출력
	println(j) //1 출력
}

func showArea(shapes ...Shape) {

	for _, s := range shapes {
		a := s.area() // 인터페이스 메서드 호출
		println(a)
	}

}

func printIt(v interface{}) {
	fmt.Println(v) // Tom
}
