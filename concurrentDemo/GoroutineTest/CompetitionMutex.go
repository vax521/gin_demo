package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup, mut *sync.Mutex) {
	mut.Lock()
	x = x + 1
	mut.Unlock()
	wg.Done()
}
func main() {
	var wg sync.WaitGroup
	var mut sync.Mutex
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment(&wg, &mut)

	}
	wg.Wait()
	fmt.Println("final value of x", x)

}
