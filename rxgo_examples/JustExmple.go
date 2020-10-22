package main

import (
	"errors"
	"fmt"

	"github.com/reactivex/rxgo/v2"
)

func main() {
	/*
		Just使用柯里化（currying）让它可以在第一个参数中接受多个数据，在第二个参数中接受多个选项定制行为。
		柯里化是函数化编程的思想，简单来说就是通过在函数中返回函数，以此来减少每个函数的参数个数。例如：
		func add(value int) func (int) int {
		  return func (a int) int {
			return value + a
		  }
		}
	    fmt.Prinlnt(add(5)(10)) // 15
	    由于 Go 不支持多个可变参数，Just通过柯里化迂回地实现了这个功能：

		// rxgo/factory.go
		func Just(items ...interface{}) func(opts ...Option) Observable {
		  return func(opts ...Option) Observable {
			return &ObservableImpl{
			  iterable: newJustIterable(items...)(opts...),
			}
		  }
		}
	*/
	observable := rxgo.Just(1, 2, 3, 4, 5)()
	ch := observable.Observe()
	for item := range ch {
		fmt.Println(item.V)
	}
	fmt.Println()

	observable1 := rxgo.Just(1, 2, errors.New("unknown"), 3, 4, 5)()
	for item := range observable1.Observe() {
		if item.Error() {
			fmt.Println("error:", item.E)
		} else {
			fmt.Println(item.V)
		}
	}
	fmt.Println()

	<-observable.ForEach(func(v interface{}) {
		fmt.Println("received:", v)
	}, func(err error) {
		fmt.Println("error:", err)
	}, func() {
		fmt.Println("completed")
	})
	fmt.Println()

	<-observable1.ForEach(func(v interface{}) {
		fmt.Println("received:", v)
	}, func(err error) {
		fmt.Println("error:", err)
	}, func() {
		fmt.Println("completed")
	})
}
