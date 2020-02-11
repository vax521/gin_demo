package main

import "fmt"

// https://leetcode.com/problems/print-in-order/ 按顺序打印

func First(streamSync [3]chan interface{}) {
	fmt.Println("First")
	streamSync[0] <- nil
}

func Second(streamSync [3]chan interface{}) {
	<-streamSync[0]
	fmt.Println("Second")
	streamSync[1] <- nil
}

func Third(streamSync [3]chan interface{}) {
	<-streamSync[1]
	fmt.Println("Third")
	streamSync[2] <- nil
}

func printInOrder(callOrder [3]int) {
	inputCallOrder := callOrder
	fmt.Println("[]inputCallOrder:", inputCallOrder)

	var streamSync [3]chan interface{}
	for i := range streamSync {
		streamSync[i] = make(chan interface{})
	}
	var funcNumMap = map[int]func([3]chan interface{}){
		1: First,
		2: Second,
		3: Third,
	}
	for _, fnum := range inputCallOrder {
		go funcNumMap[fnum](streamSync)
	}
	<-streamSync[2]
}

func main() {
	var testCases = [][3]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}

	for _, theCase := range testCases {
		printInOrder(theCase)
		fmt.Println()
		fmt.Println()
	}
}
