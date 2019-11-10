package main

import "fmt"

//环形队列实现
type CycleQueue struct {
	data  []interface{}
	front int
	rear  int
	cap   int
}

func NewCycleQueue(cap int) *CycleQueue {
	return &CycleQueue{
		data:  make([]interface{}, cap),
		front: 0,
		rear:  0,
		cap:   cap,
	}
}

func (cq *CycleQueue) Push(data interface{}) bool {
	//队列已满
	if (cq.rear+1)%cq.cap == cq.front {
		return false
	}
	cq.data[cq.rear] = data          //将元素放入队列尾部
	cq.rear = (cq.rear + 1) % cq.cap //尾部元素指向下一个空间位置,取模运算保证了索引不越界（余数一定小于除数）
	return true
}

func (cq *CycleQueue) Pop() interface{} {
	if cq.front == cq.rear {
		return nil
	}
	data := cq.data[cq.front]
	cq.data[cq.front] = nil
	cq.front = (cq.front + 1) % cq.cap
	return data
}

func main() {
	cq := NewCycleQueue(5)
	cq.Push(4)
	cq.Push(5)
	cq.Push(5)
	cq.Push(6)
	cq.Push(7)
	fmt.Println(cq.Pop())
	fmt.Println(cq.Pop())
	fmt.Println(cq.Pop())
	fmt.Println(cq.Pop())
	fmt.Println(cq.Pop())
	fmt.Println(cq.Pop())

}
