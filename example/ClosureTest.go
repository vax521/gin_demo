package main

import "fmt"

func main() {
	var s = "Hello,world"
	fmt.Println(s)
	//在闭包内部修改引用的变量
	foo := func() {
		//在匿名函数中访问s
		s = "hello dude"
		fmt.Println(s)
	}
	foo()
}
