package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

//套接字接受过程
func socketRecvWg(conn net.Conn, wg *sync.WaitGroup) {
	//创建接收缓冲
	buff := make([]byte, 1024)
	for {
		_, err := conn.Read(buff)
		if err != nil {
			break
		}
	}
	//函数结束发出通知
	wg.Done()
}

//避免在不必要的地方使用通道
func main() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Println(err)
		return
	}
	var wg sync.WaitGroup

	wg.Add(1)
	go socketRecvWg(conn, &wg)
	time.Sleep(time.Second)
	conn.Close()
	//等待通知
	wg.Wait()
	fmt.Println("receive done")

}
