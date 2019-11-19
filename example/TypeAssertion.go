package main

import "fmt"

func assrt(i interface{}) {
	s, ok := i.(int)
	fmt.Println(s, ok)
}

func main() {
	var i interface{} = "heelop"
	assrt(i)
}
