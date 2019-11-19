package main

import "fmt"

//用户自定义函数
type add func(a int, b int) int

func main() {
	a := func() {
		fmt.Println("first class function!")
	}
	a()
	fmt.Printf("%T", a)
	//匿名函数
	func(a string) {
		fmt.Printf("Hello,%s", a)
	}("hali")

	var sum add = func(a int, b int) int {
		return a + b
	}
	s := sum(4, 5)
	fmt.Println("Sum:", s)
}
