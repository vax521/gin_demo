package main

import "fmt"

func main() {
	// 键值循环
	for key, value := range []int{1, 2, 3, 4} {
		fmt.Printf("key:%d value:%d\n", key, value)
	}
	//遍历字符串
	var str = "hello你好"
	for key, value := range str {
		fmt.Printf("key:%d value:%c\n", key, value)
	}
	//遍历通道
	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
}
