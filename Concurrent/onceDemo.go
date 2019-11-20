package main

import (
	"fmt"
	"sync"
)

func One() {
	fmt.Println("One")
}

func Two() {
	fmt.Println("Two")
}

func main() {
	var once sync.Once
	for i, v := range make([]string, 10) {
		once.Do(One)
		fmt.Println("count:", v, "---", i)
	}
}
