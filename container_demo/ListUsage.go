package main

import (
	"container/list"
	"fmt"
)

func printList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func main() {
	listA := list.New()
	listA.PushFront(1)
	listA.PushBack(2)
	listA.PushFront(3)
	printList(listA)
	fmt.Println("===========================")
	for e := listA.Front(); e != nil; e = e.Next() {
		if e.Value == 1 {
			listA.InsertBefore(1.1, e)
		}
		if e.Value == 2 {
			listA.InsertBefore(1.1, e)
		}
	}
	printList(listA)
	fmt.Println("===========================")
	fmt.Println(listA.Front().Value)
	fmt.Println(listA.Back().Value)
	fmt.Println("===========================")
	listA.MoveToFront(listA.Back())
	printList(listA)
	fmt.Println("===========================")
	for e := listA.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}

}
