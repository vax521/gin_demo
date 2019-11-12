package main

import "fmt"

func square(number int, sqchan chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	fmt.Println("square finished!")
	sqchan <- sum
}
func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	fmt.Println("cube finished!")
	cubeop <- sum
}
func main() {
	number := 589
	squchan := make(chan int)
	cubechan := make(chan int)
	go square(number, squchan)
	go calcCubes(number, cubechan)
	sq, cube := <-squchan, <-cubechan
	fmt.Printf("final result:%d\n", sq+cube)

}
