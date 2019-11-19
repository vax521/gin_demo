package main

import "fmt"

func iMap(s []int, f func(int) int) []int {
	var r []int
	for _, v := range s {
		r = append(r, f(v))
	}
	return r
}

func main() {
	a := []int{1, 2, 3, 5, 6, 7}
	result := iMap(a, func(i int) int {
		return i * 5
	})
	fmt.Println(result)
}
