package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var lock sync.Mutex
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("count1++ is %d\n", count)
	}()
	go func() {
		defer wg.Done()
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("count2++ is %d\n", count)
	}()

	wg.Add(2)
	go func() {
		defer wg.Done()
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("count1-- is %d\n", count)
	}()
	go func() {
		defer wg.Done()
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("count2-- is %d\n", count)
	}()

	wg.Wait()
	fmt.Printf("count is %d\n", count)

}
