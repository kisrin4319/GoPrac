package main

func main() {

	msg := "Hello"
	say(&msg)
	println(msg) //변경된 메시지 출력

	say2("This", "is", "a", "book")
	say2("Hi")

	total := sum(1, 7, 3, 5, 9)
	println(total)

	count, total := sum2(1, 7, 3, 5, 9)
	println(count, total)

}

func say(msg *string) {
	println(*msg)
	*msg = "Changed" // 메시지 번경
}

func say2(msg ...string) {
	for _, s := range msg {
		print(s)
	}
}
func sum(nums ...int) int {
	s := 0
	for _, n := range nums {
		s += n
	}
	return s
}
func sum2(nums ...int) (int, int) {
	s := 0     // 합계
	count := 0 // 요소 갯수
	for _, n := range nums {
		s += n
		count++
	}
	return count, s
}
