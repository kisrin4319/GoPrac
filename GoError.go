package main

import (
	"fmt"
	"log"
	"os"
)

type error interface {
	Error() string
}

func main() {
	f, err := os.Open("C:\\temp\\1.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	println(f.Name())

	_, err := otherFunc()
	switch err.(type) {
	default:
		println("ok")
	case MyError:
		log.Print("Log my error")
	case error:
		log.Fatal(err.Error())
	}

}
