package main

import (
	"bytes"
	"fmt"
)

//定义一个函数，参数量为0-n,类型为字符串
func joinStrings(slist ...string) string {
	var b bytes.Buffer
	for _, s := range slist {
		b.WriteString(s)
	}
	return b.String()
}

//定义一个函数，传入可变参数类型，打印其类型和值
func printTypeValues(slist ...interface{}) string {
	var b bytes.Buffer
	for _, s := range slist {
		//将interface{}类型格式化为字符串
		str := fmt.Sprintf("%v", s)
		//类型的字符串描述
		var typeString string
		//对s进行类型断言
		switch s.(type) {
		case bool:
			typeString = "bool"
		case string:
			typeString = "string"
		case int:
			typeString = "int"
		}
		b.WriteString("value:")
		b.WriteString(str + " ")
		b.WriteString("type:")
		b.WriteString(typeString)
		b.WriteString("\n")
	}
	return b.String()
}

func main() {
	fmt.Println(joinStrings("hello", "world", "name"))
	fmt.Println(joinStrings("hello", "world", "name", "nature"))
	fmt.Println(printTypeValues(100, "str", true))
}
