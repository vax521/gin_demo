package main

import (
	"fmt"
	"github.com/liyue201/gostl/ds/list/simple_list"
)

func main() {
	listA := simple_list.New()
	listA.PushBack(3)
	listA.PushBack(4)
	listA.PushBack(5)
	listA.PushFront(6)
	listA.PushFront(78)
	listA.PushFront(9)
	listA.PushFront(88)

	//简单列表是一种单向列表，支持从头部和尾部插入数据，只支持从头部遍历数据。
	for n := listA.FrontNode(); n != nil; n = n.Next() {
		fmt.Println(n.Value)
	}
}
