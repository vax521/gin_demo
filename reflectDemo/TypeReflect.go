package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int
	typeofa := reflect.TypeOf(a)
	fmt.Println(typeofa, typeofa.Name(), typeofa.Kind())

	type Enum int
	const Zero Enum = 0

	type cat struct {
	}
	maomi := cat{}

	typeofmaomi := reflect.TypeOf(maomi)
	fmt.Println(typeofmaomi.Name(), typeofmaomi.Kind())

	typeofZero := reflect.TypeOf(Zero)
	fmt.Println(typeofZero.Name(), typeofZero.Kind())

	ins := &cat{}
	typeofIns := reflect.TypeOf(ins)
	fmt.Printf("name:%v, kind:%v\n", typeofIns.Name(), typeofIns.Kind())
	//获取类型的元素
	typeofIns = typeofIns.Elem()
	fmt.Printf("element name:%v,element kind:%v\n", typeofIns.Name(), typeofIns.Kind())

}
