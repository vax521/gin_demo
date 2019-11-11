package main

import (
	"fmt"
	"github.com/liyue201/gostl/ds/queue"
)

//队列是一种先进先出的数据结构，底层使用双端队列或者链表作为容器，默认使用双端队列，若想使用链表，可以在创建对象时使用queue.WithListContainer()参数。支持线程安全。

func simpleExam1() {
	queA := queue.New()
	queA.Push(5)
	queA.Push(6)
	queA.Push(8)
	queA.Push(9)
	for !queA.Empty() {
		fmt.Printf("%v\n", queA.Pop())
	}
}

func main() {
	simpleExam1()
}
