package main

import "fmt"

/*
计算数组列表的排列组合
*/
func reduceArr(arr [][]string) []string {
	if len(arr) >= 2 {
		arrLen1 := len(arr[0])
		arrLen2 := len(arr[1])
		var tempArr []string
		for i := 0; i < arrLen1; i++ {
			for j := 0; j < arrLen2; j++ {
				tempArr = append(tempArr, arr[0][i]+" "+arr[1][j])
			}
		}
		var newArr [][]string
		newArr = append(newArr, tempArr)
		for i := 2; i < len(arr); i++ {
			newArr = append(newArr, arr[i])
		}
		return reduceArr(newArr)
	} else {
		return arr[0]
	}
}

func main() {

	var stringArray = [][]string{{"xing", "meng"}, {"cai", "hai"}, {"niu", "hai"}}
	resultList := reduceArr(stringArray)
	fmt.Println(resultList)
	for _, value := range resultList {
		fmt.Println(value)
	}

}
