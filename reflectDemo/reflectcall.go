package main

import (
	"fmt"
	"reflect"
)

func add(a, b int) int {
	return a + b
}

func main() {
	//将函数包装为反射值对象
	funcValue := reflect.ValueOf(add)
	//构造函数参数，传入2个参数
	paramList := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20)}
	//反射调用函数
	retList := funcValue.Call(paramList)
	fmt.Println(retList[0].Int())

}
