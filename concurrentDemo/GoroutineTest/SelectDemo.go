package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/**
select语句选择一组可能的send操作和receive操作去处理。它类似switch,但是只是用来处理通讯(communication)操作。
它的case可以是send语句，也可以是receive语句，亦或者default。
receive语句可以将值赋值给一个或者两个变量。它必须是一个receive操作。
最多允许有一个default case,它可以放在case列表的任何位置，尽管我们大部分会将它放在最后。
*/
func fabonacci(c, q chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-q:
			fmt.Println("quit...")
			return
		}
	}
}

func RandomGene(ch chan int) {
	for {
		select {
		case ch <- 0:
		case ch <- 1:
		case ch <- 2:
		case ch <- 3:
		case ch <- 4:
		case ch <- 5:
		}
	}
}

//select和default分支可以很容易实现一个Goroutine的退出控制:
func worker(ch chan bool) {
	for {
		select {
		default:
			fmt.Println("hello")
		case <-ch:
			return
		}
	}
}

/*我们通过close()来关闭cancel通道，向多个Goroutine广播退出的指令。
不过这个程序依然不够稳健：当每个Goroutine收到退出指令退出时一般会进行一定的清理工作，但是退出的清理工作并不能保证被完成，
因为main线程并没有等待各个工作Goroutine退出工作完成的机制。我们可以结合sync.WaitGroup来改进：*/
func workerWithWg(wg *sync.WaitGroup, ch chan bool) {
	defer wg.Done()
	for {
		select {
		default:
			fmt.Println("hello")
		case <-ch:
			return
		}
	}
}

/*标准库增加了一个context包，用来简化对于处理单个请求的多个Goroutine之间与请求域的数据、超时和退出等操作*/
func workerWithContext(context context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()
	for {
		select {
		default:
			fmt.Println("hello")
		case <-context.Done():
			return context.Err()
		}
	}
}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go workerWithContext(ctx, &wg)
	}
	time.Sleep(5 * time.Second)
	cancel()
	wg.Wait()
	/*cancel := make(chan bool)
	var wg sync.WaitGroup
	for i:=0;i < 10; i++{
		wg.Add(1)
		go workerWithWg(&wg, cancel)
	}
	time.Sleep(5*time.Second)
	close(cancel)
	wg.Wait()*/
	/*通道的发送操作和接收操作是一一对应的，如果要停止多个Goroutine，那么可能需要创建同样数量的通道，这个代价太大了。
	其实我们可以通过close()关闭一个通道来实现广播的效果，所有从关闭通道接收的操作均会收到一个零值和一个可选的失败标志。*/
	/*cancel := make(chan bool)
	for i:=0;i < 10; i++{
		go worker(cancel)
	}
	time.Sleep(5*time.Second)
	close(cancel)*/
	/*cancel := make(chan bool)
		go worker(cancel)
	    time.Sleep(5*time.Second)
		cancel <- true*/
	/* ch := make(chan int);
	go RandomGene(ch);
	for v := range ch{
		fmt.Println(v)
	}*/
	/*c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fabonacci(c, quit)*/

	//超时处理timeout

	/*chnl := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		chnl <- "result 1"
	}()
	for {
		select {
		case res := <-chnl:
			fmt.Println(res)
			return
		case <-time.After(time.Second * 1):
			fmt.Println("time out 1s")
		}
	}*/

}
