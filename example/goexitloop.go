package main

import "fmt"

func main() {
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if y == 2 {
				goto breakHere
			}
		}
	}
	//手动返回，避免进入标签
	return
breakHere:
	fmt.Println("Done")
}
