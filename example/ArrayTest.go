package main

import "fmt"

func main() {
	var team [3]string
	team[0] = "hammer"
	team[1] = "soldier"
	team[2] = "mum"
	fmt.Println(team)
	//初始化数组
	var team1 = [3]string{"hammer", "soldier", "mum"}
	fmt.Println(team1)
	// 遍历数组
	for k, v := range team1 {
		fmt.Println(k, v)
	}

	var arr1 = [3]int{1, 2, 3}
	fmt.Println(arr1, arr1[1:2])

	arr2 := []int{1, 2, 3}
	fmt.Println(arr2[:])
	fmt.Println(arr2[0:0])
	fmt.Println(arr2)
	//声明切片
	var strList []string
	//strList[0] = "a"
	var numList []int
	//声明一个空的整形切片
	var numListEmpty = []int{}
	fmt.Println(strList, numList, numListEmpty)
	fmt.Println(len(strList), len(numList), len(numListEmpty))
	fmt.Println(strList == nil)      //true
	fmt.Println(numList == nil)      //true
	fmt.Println(numListEmpty == nil) //false

	//通过make()构造切片
	ab := make([]int, 2)
	ac := make([]int, 2, 10)
	fmt.Println(ab, ac)
	fmt.Println(len(ab), len(ac))
	//使用append为切片添加元素
	var nums []int
	for i := 0; i < 10; i++ {
		nums = append(nums, i)
		fmt.Printf("len:%d cap:%d pointer:%p\n", len(nums), cap(nums), nums)
	}
	var car []string
	car = append(car, "Old Driver")
	fmt.Println(car)
	car = append(car, "Ice", "Snipper")
	fmt.Println(car)
	//team := []string{"pig","dog"}
	//car = append(car,team...)

}
