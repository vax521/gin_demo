package main

import "fmt"

func visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}
func main() {
	//声明时调用匿名函数
	func(data int) {
		fmt.Println("hello,", data)
	}(100)
	//将匿名函数赋值给变量
	f := func(data int) {
		fmt.Println("hello,", data)
	}
	f(100)
	// 匿名函数用作回调函数
	visit([]int{1, 2, 3, 4}, func(v int) {
		fmt.Println(v)
	})
}
