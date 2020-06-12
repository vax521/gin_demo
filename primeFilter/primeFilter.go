package main

import "fmt"

func GenerateNatural() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

//通道过滤器，删除能被素数整除的数
func PrimeFilter(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func main() {
	ch := GenerateNatural()
	for i := 0; i < 100; i++ {
		prime := <-ch
		fmt.Printf("%v:%v\n", i+1, prime)
		ch = PrimeFilter(ch, prime)
	}
}
