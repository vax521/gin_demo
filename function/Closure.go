package main

import "fmt"

func appendStr() func(string) string {
	t := "hello"
	c := func(b string) string {
		return t + " " + b
	}
	return c
}

func main() {
	a := appendStr()
	b := appendStr()
	fmt.Println(a("World"))
	fmt.Println(b("Everyone"))

	fmt.Println(a("Gopher"))
	fmt.Println(b("!"))
}
