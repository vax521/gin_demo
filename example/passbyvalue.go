package main

import "fmt"

type Data struct {
	complex64 []int
	instance  InnerData
	ptr       *InnerData
}
type InnerData struct {
	a int
}

func passByValue(inFunc Data) Data {
	fmt.Printf("in Func value:%+v\n", inFunc)
	fmt.Printf("inFunc ptr:%p\n", &inFunc)
	return inFunc
}

func main() {
	in := Data{
		complex64: []int{1, 2, 3},
		instance:  InnerData{5},
		ptr:       &InnerData{1},
	}
	fmt.Printf("in value:%+v\n", in)
	fmt.Printf("in ptr:%p\n", &in)
	//传入结构体，返回同类型的结构体
	out := passByValue(in)
	fmt.Printf("out value:%+v\n", out)
	fmt.Printf("out ptr: %p\n", &out)
}
