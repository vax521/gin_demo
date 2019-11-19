package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("当前版本：" + runtime.Version())
}
