package main

import (
	"fmt"
	"sort"
)

func main() {
	scene := make(map[string]int)
	scene["route"] = 66
	fmt.Println(scene["route"])
	v := scene["route"]
	fmt.Println(v)
	//声明时初始化
	m := map[string]string{
		"W": "forward",
		"S": "backward",
		"D": "right",
		"A": "left",
	}
	fmt.Println(m)
	var mList []string
	//遍历map
	for key := range m {
		fmt.Println(key)
		mList = append(mList, key)
	}
	sort.Strings(mList)
	fmt.Println(mList)
	//使用delete()函数删除键值对
	delete(m, "A")
	for k, v := range m {
		fmt.Println(k, v)
	}
	//清空map
	m = make(map[string]string)
}
