package main

import (
	"fmt"
	"strings"
)

func StringProcess(list []string, chain []func(string) string) {
	// 遍历每一个字符串
	for index, str := range list {
		result := str
		//遍历每一个处理链
		for _, proc := range chain {
			result = proc(result)
		}
		//将结果放回切片
		list[index] = result
	}
}
func removePrefix(str string) string {
	return strings.TrimPrefix(str, "go")
}

func main() {
	list := []string{
		"go scanner",
		"go parser",
		"go pointer",
		"go printer",
	}
	//处理函数链
	chain := []func(string) string{
		removePrefix, strings.TrimSpace, strings.ToUpper,
	}
	StringProcess(list, chain)
	for _, str := range list {
		fmt.Println(str)
	}
}
