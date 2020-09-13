package main

import "fmt"

func partition(array []int, left, right int) int {
	baseNum := array[left]
	for left < right {
		for array[right] >= baseNum && right > left {
			right--
		}
		array[left] = array[right]
		for array[left] <= baseNum && right > left {
			left++
		}
		array[right] = array[left]
	}
	array[right] = baseNum
	return right
}

func quickSort(array []int, left, right int) {
	if left >= right {
		return
	}
	//fmt.Printf("%v\n",array)
	index := partition(array, left, right)
	quickSort(array, left, index-1)
	quickSort(array, index+1, right)
}

//带有返回值的快速排序写法
func qucickSortWithBackValue(array []int) []int {
	if len(array) == 0 {
		return []int{}
	}
	if len(array) == 1 {
		return array
	}
	pivot := array[0]
	var pivots []int
	var smallArray []int
	var bigArray []int
	for _, value := range array {
		if value > pivot {
			bigArray = append(bigArray, value)
		} else if value > pivot {
			smallArray = append(smallArray, value)
		} else {
			pivots = append(pivots, pivot)
		}
	}
	tempArray := append(qucickSortWithBackValue(smallArray), pivots...)
	tempArray = append(tempArray, qucickSortWithBackValue(bigArray)...)
	return tempArray
}

/*//实现快速排序
func quickSort(nums []int, start, end int) []int {
	if start < end {
		i, j := start, end
		key := nums[(start+end)/2]
		print(key)
		for nums[i] < key{
			i++
		}
		for nums[j] > key {
			j--
		}
		if i <= j{
			nums[i],nums[j] = nums[j],nums[i]
			i++
			j--
		}
		nums = append(quickSort(nums, start, j), key)
		nums = append(nums, quickSort(nums,i,end)...)

	}

	return nums
}*/

func main() {
	nums := []int{45, 63, 3, 1, 29, 77, 20, 4, 30}
	quickSort(nums, 0, len(nums)-1)
	fmt.Printf("%v\n", nums)
	fmt.Printf("%v\n", qucickSortWithBackValue(nums))
}
