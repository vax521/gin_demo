package main

import "fmt"

type Invoker interface {
	// 参数为任意类型的值
	call(interface{})
}

// 结果体类型
type Struct struct {
}

// 结构体实现接口
func (s *Struct) call(p interface{}) {
	fmt.Println("from struct", p)
}

//函数定义为类型
type FuncCaller func(interface{})

// 函数体实现接口
func (f FuncCaller) call(p interface{}) {
	f(p)
}

func main() {
	//声明接口变量
	var invoker Invoker
	s := new(Struct)
	invoker = s
	invoker.call("hello")

	// 将匿名函数转为FuncCaller类型，再赋值给接口
	invoker = FuncCaller(func(v interface{}) {
		fmt.Println("from fucntion:", v)
	})
	//使用接口调用FuncCaller.Call,内部会调用函数本体
	invoker.call("hello")
}
