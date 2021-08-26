package main

import "fmt"

func main() {
	Hanoi(3)

	//3. fibonacci
	fmt.Println(Fib(90))
}

func Move(n, from, to, via int) {
	if n <= 0 {
		return
	}
	Move(n-1, from, via, to)
	fmt.Println(from, "->", to)
	Move(n-1, via, to, from)
}
func Hanoi(n int) {
	fmt.Println("Number of disks:", n)
	Move(n, 1, 2, 3)
}

func Fib(n int) int {
	p, q := 0, 1
	for i := 0; i < n; i++ {
		p, q = q, p+q
	}
	return p

}
