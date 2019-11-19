package main

import "fmt"

//利用闭包的记忆效应实现一个累加器
//提供一个值，每次调用函数会指定对值进行累加
func Accumulate(value int) func() int {
	//返回一个闭包
	return func() int {
		value++
		return value
	}
}
func main() {
	accumulator := Accumulate(1)
	fmt.Println(accumulator())
	fmt.Println(accumulator())
	fmt.Println(accumulator())
	fmt.Println(accumulator())
	//打印累加器的函数地址
	fmt.Printf("%p\n", accumulator)
	//创建一个累加器，初始值为10
	accumulator2 := Accumulate(10)
	fmt.Println(accumulator2())
	fmt.Printf("%p\n", accumulator2)

}
