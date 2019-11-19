package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 1024
	valueOfA := reflect.ValueOf(a)
	var getA int = valueOfA.Interface().(int)
	var getA2 int = int(valueOfA.Int())
	fmt.Println(getA, getA2)
}
