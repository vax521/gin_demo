package main

import "fmt"

func GetData() (int, int) {
	return 100, 200
}
func main() {
	//var c_group []float32
	//var d func() bool
	//var e struct{ x int}

	// 变量定义
	var a int = 10
	var b = 10
	c := 10
	fmt.Println(a, b, c)

	var x, y int
	x, y = 1, 2
	fmt.Println(x, y)

	var (
		z int
		k bool
	)
	z, k = 5, true
	fmt.Println(z, k)

	f, g := 99, 100
	fmt.Println(f, g)

	// 匿名变量，用_表示
	first, _ := GetData()
	_, last := GetData()
	fmt.Println(first, last)
}
