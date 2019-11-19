package main

import (
	"container/list"
	"fmt"
)

func main() {
	//初始化列表
	l := list.New()
	l.PushBack("canon")
	l.PushFront(67)
	fmt.Println(l)
	//尾部添加后保存元素句柄
	element := l.PushBack("first")
	//first之后添加high
	l.InsertAfter("high", element)
	//first之前添加noon
	l.InsertBefore("noon", element)
	l.Remove(element)
	//遍历列表
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
