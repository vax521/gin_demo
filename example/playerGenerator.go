package main

import "fmt"

//利用闭包的记一下效应实现生成器
func playerGen(name string) func() (string, int) {
	//血量一直为150
	hp := 150
	//返回创建的闭包
	return func() (string, int) {
		return name, hp
	}
}
func main() {
	generator := playerGen("high noon")
	name, hp := generator()
	fmt.Println(name, hp)
}
