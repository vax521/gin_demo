package main

import (
	"fmt"
	"sync"
)

func main() {
	var scene sync.Map
	go func() {
		scene.Store("greece", 97)
		scene.Store("london", 10)
		scene.Store("sgypt", 200)
	}()

	go func() {
		fmt.Println(scene.Load("london"))
		scene.Delete("london")

		scene.Range(func(key, value interface{}) bool {
			fmt.Println("iterate:", key, value)
			return true
		})
	}()

}
