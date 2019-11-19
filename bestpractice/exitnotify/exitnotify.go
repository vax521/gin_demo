package main

import (
	"fmt"
	"net"
	"time"
)

//套接字接受过程
func socketRecv(conn net.Conn, exitChan chan string) {
	//创建接收缓冲
	buff := make([]byte, 1024)
	for {
		_, err := conn.Read(buff)
		if err != nil {
			break
		}
	}
	exitChan <- "recev exit "
}

//避免在不必要的地方使用通道
func main() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Println(err)
		return
	}
	exit := make(chan string)
	go socketRecv(conn, exit)
	time.Sleep(time.Second)
	conn.Close()

	fmt.Println(<-exit)
}
