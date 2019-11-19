package main

import (
	"fmt"
	"reflect"
)

//使用反射值对象修改变量的值
//两个条件：可被寻址；被导出；

func main() {
	var a int = 1024
	//valueOfA := reflect.ValueOf(a)
	//获得变量a的地址
	valueOfA := reflect.ValueOf(&a)
	//取出A的值
	valueOfA = valueOfA.Elem()
	valueOfA.SetInt(1)
	fmt.Println(valueOfA)

	type dog struct {
		//legCount int
		LegCount int
	}
	valueOfDog := reflect.ValueOf(&dog{})
	valueOfDog = valueOfDog.Elem()
	vLegCount := valueOfDog.FieldByName("LegCount")
	vLegCount.SetInt(4)
	fmt.Println(vLegCount)

	//通过类型创建类型的实例
	typeofA := reflect.TypeOf(a)
	aIns := reflect.New(typeofA)
	fmt.Println(aIns, aIns.Type(), aIns.Kind())

	aIns = aIns.Elem()
	aIns.SetInt(4)
	fmt.Println(aIns, aIns.Type(), aIns.Kind())
}
