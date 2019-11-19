package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func server(address string, exitChan chan int) {
	l, err := net.Listen("tcp", address)
	//如果发生侦听错误，打印错误并exit
	if err != nil {
		fmt.Println(err.Error())
		exitChan <- 1
	}
	//打印监听地址，表示监听成功
	fmt.Println("listen:" + address)
	//延迟关闭侦听器
	defer l.Close()
	//侦听循环
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		//根据链接开始会话，并行执行
		go handleSession(conn, exitChan)
	}

}
func processTelnetCommand(str string, exitChan chan int) bool {
	//@close终止会话
	if strings.HasPrefix(str, "@close") {
		return false
	} else if strings.HasPrefix(str, "@shutdown") {
		fmt.Println("Server shutdown")
		exitChan <- 0
		return false
	}
	//打印要输出的字符串
	fmt.Println(str)
	return true
}
func handleSession(conn net.Conn, exitChan chan int) {
	//创建一个网络链接数据的读取器
	reader := bufio.NewReader(conn)
	//接受数据的循环
	for {
		//读取字符串，知道碰到回车返回
		str, err := reader.ReadString('\n')
		if err == nil {
			//去掉字符串尾部的回车
			str := strings.TrimSpace(str)
			//处理telnet指令
			if !processTelnetCommand(str, exitChan) {
				conn.Close()
				break
			}
			//echo逻辑
			conn.Write([]byte(str + "\r\n"))
		} else {
			fmt.Println("Session Closed.")
			conn.Close()
			break
		}
	}
}

func main() {
	exitChan := make(chan int)
	go server("127.0.0.1:8000", exitChan)
	code := <-exitChan
	os.Exit(code)
}
