package main

import "fmt"

func xrange() chan int { // xrange用来生成自增的整数
	var ch chan int = make(chan int)

	go func() { // 开出一个goroutine
		for i := 2; ; i++ {
			ch <- i // 直到信道索要数据，才把i添加进信道
		}
	}()

	return ch
}

// 输入一个整数队列，筛出是number倍数的, 不是number的倍数的放入输出队列
// in:  输入队列
func filter(in chan int, number int) chan int {
	ch := make(chan int)
	go func() {
		for {
			i := <-in
			if i%number != 0 {
				ch <- i
			}
		}
	}()

	return ch
}

func main() {
	const max = 100
	nums := xrange()
	number := <-nums
	for number <= max {
		fmt.Println(number)
		nums = filter(nums, number)
		number = <-nums
	}
}
