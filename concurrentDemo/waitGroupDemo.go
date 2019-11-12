package main

import (
	"fmt"
	"sync"
	"time"
)

/**
WaitGroup用于等待一组线程的结束。
父线程调用Add方法来设定应等待的线程的数量。
每个被等待的线程在结束时应调用Done方法。
同时，主线程里可以调用Wait方法阻塞至所有线程结束。
*/
func demo1() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("goroutine1 结束！")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("goroutine2 结束！")
	}()

	wg.Wait()

}

func demo2() {

	//sync.WaitGroup结构定义后就不能被复制，所以这里要使用指针。
	sayHello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("goroutine %d start.....\n ", id)
		time.Sleep(2 * time.Second)
		fmt.Printf("goroutine %d exit..... \n", id)
	}

	var wg sync.WaitGroup
	const N = 5
	wg.Add(N)
	for i := 0; i < N; i++ {
		go sayHello(&wg, i)
	}
	fmt.Println("All goroutine prepare to run...")
	wg.Wait()
	fmt.Println("All goroutine finished....")
}

/**
过WaitGroup提供的三个函数：Add,Done,Wait，可以轻松实现等待某个协程或协程组完成的同步操作。但在使用时要注意:
Add的数量和Done的调用数量必须相等。
WaitGroup结构一旦定义就不能复制的原因。
*/
func main() {
	demo2()
}
