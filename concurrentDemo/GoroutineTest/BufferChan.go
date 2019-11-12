package main

import (
	"fmt"
	"time"
)

func write(ch chan int) {
	for i := 1; i < 10; i++ {
		ch <- i
		fmt.Println("Successfully wrote ", i, " to ch")
	}
	close(ch)
}

func main() {
	buffChan := make(chan string, 2)
	//缓冲信道的容量是指信道可以存储的值的数量。我们在使用 make 函数创建缓冲信道的时候会指定容量大小。
	//缓冲信道的长度是指信道中当前排队的元素个数。
	buffChan <- "string1"
	buffChan <- "string2"
	//buffChan <- "string3"
	fmt.Println("Cap:", cap(buffChan))
	fmt.Println("length:", len(buffChan))
	fmt.Println(<-buffChan)
	fmt.Println("Cap:", cap(buffChan))
	fmt.Println("length:", len(buffChan))
	fmt.Println(<-buffChan)
	//fmt.Println(<-buffChan)

	ch := make(chan int, 2)
	go write(ch)
	time.Sleep(2 * time.Second)
	for num := range ch {
		fmt.Println("receivied", num, " from ch")
		time.Sleep(2 * time.Second)
	}

}
