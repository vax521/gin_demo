package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//使用匿名函数创建goroutine
	go func() {
		var times int
		for {
			times++
			fmt.Println("tick", times)
			time.Sleep(time.Second)
		}
	}()
	fmt.Println("CPU核心数：", runtime.NumCPU())
	var input string
	fmt.Scanln(&input)

}
