package main

import "fmt"

type speaker interface {
	speak()
}

type english struct{}

type chinese struct{}

func (english) speak() {
	fmt.Println("Hello World")

}

func (chinese) speak() {
	fmt.Println("你好世界")
}

func main() {

	var sp speaker
	var e english
	sp = e
	sp.speak()

	var c chinese
	sp = c
	sp.speak()

	greet(new(english))
	greet(new(chinese))

}

func greet(sp speaker) {
	sp.speak()
}
