package main

import (
	"fmt"
	"github.com/liyue201/gostl/ds/list/bid_list"
)

func main() {
	biList := bid_list.New()
	biList.PushFront(1)
	biList.PushFront(4)
	biList.PushFront(6)
	biList.PushFront(9)
	biList.PushBack(11)
	biList.PushBack(15)
	biList.PushBack(17)
	biList.PushBack(19)
	//正序遍历
	for node := biList.FrontNode(); node != nil; node = node.Next() {
		fmt.Println(node.Value)
	}
	fmt.Println("___________________")
	for n := biList.BackNode(); n != nil; n = n.Prev() {
		fmt.Println(n.Value)
	}
}
