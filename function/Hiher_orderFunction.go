package main

import "fmt"

//函数作为参数传递
func simple(f func(a int, b int) int) {
	fmt.Println(f(60, 7))
}
func complex() func(a, b int) int {
	f := func(a int, b int) int {
		return a + b
	}
	return f
}

func main() {
	f := func(a int, b int) int {
		return a + b
	}
	simple(f)
	c := complex()
	fmt.Println(c(60, 2))
}
