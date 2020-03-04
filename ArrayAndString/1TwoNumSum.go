package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var resultArr []int
	dict := map[int]int{}
	for index, value := range nums {
		reNum := target - value
		_, ok := dict[reNum]
		if ok {
			resultArr = append(resultArr, dict[reNum])
			resultArr = append(resultArr, index)
			return resultArr
		}
		dict[value] = index
	}
	return nil
}

func main() {
	testArr := []int{2, 3, 4, -4, 5, 6, 7, 9}
	fmt.Println(twoSum(testArr, 5))
}
