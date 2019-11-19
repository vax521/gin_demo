package main

import (
	"fmt"
	"time"
)

func running() {
	var times int
	for {
		times++
		fmt.Println("tick", times)
		time.Sleep(time.Second)
	}
}

func main() {
	//并发执行
	go running()

	//接受命令行输入
	var input string
	fmt.Scanln((&input))
}
