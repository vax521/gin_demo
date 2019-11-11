package main

import (
	"fmt"
	"github.com/liyue201/gostl/algorithm/sort"
	"github.com/liyue201/gostl/ds/deque"
)

func main() {
	//双端队列支持从头部和尾部高效的插入数据，支持随机访问和迭代器访问。
	dequeA := deque.New()
	dequeA.PushBack(2)
	dequeA.PushFront(4)
	dequeA.PushBack(6)
	dequeA.PushFront(8)
	fmt.Printf("%v\n", dequeA)
	sort.Sort(dequeA.Begin(), dequeA.End())
	fmt.Printf("%v\n", dequeA)
	fmt.Println(dequeA.PopBack())
	fmt.Println(dequeA.PopBack())
	fmt.Println(dequeA.PopFront())
	fmt.Println(dequeA.PopFront())
	fmt.Println(dequeA)
}
