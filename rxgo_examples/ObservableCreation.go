package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/reactivex/rxgo/v2"
)

func printOberser(obse rxgo.Observable) {
	<-obse.ForEach(func(v interface{}) {
		fmt.Println("received:", v)
	}, func(err error) {
		fmt.Println("error:", err)
	}, func() {
		fmt.Println("completed")
	})
	fmt.Println()
}
func Supplier1(ctx context.Context) rxgo.Item {
	return rxgo.Of(1)
}
func Supplier2(ctx context.Context) rxgo.Item {
	return rxgo.Of(2)
}
func Supplier3(ctx context.Context) rxgo.Item {
	return rxgo.Of(3)
}

func main() {
	//使用Create方法创建Observable
	obser1 := rxgo.Create([]rxgo.Producer{
		func(ctx context.Context, next chan<- rxgo.Item) {
			next <- rxgo.Of(1)
			next <- rxgo.Of(2)
			next <- rxgo.Of(3)
			next <- rxgo.Error(errors.New("unknow"))
			next <- rxgo.Of(4)
			next <- rxgo.Of(5)
		}})
	printOberser(obser1)

	//使用FromChannel创建
	//FromChannel可以直接从一个已存在的<-chan rxgo.Item对象中创建 Observable：
	ch := make(chan rxgo.Item)
	go func() {
		for i := 1; i < 5; i++ {
			ch <- rxgo.Of(i)
		}
		close(ch)
	}()
	obseFromChannel := rxgo.FromChannel(ch)
	printOberser(obseFromChannel)
	/*
		Interval
	   Interval以传入的时间间隔生成一个无穷的数字序列，从 0 开始：
	*/

	/*	obserInterval := rxgo.Interval(rxgo.WithDuration(1*time.Second))
		<-obserInterval.ForEach(func(v interface{}) {
			fmt.Println("received:", v)
		}, func(err error) {
			fmt.Println("error:", err)
		}, func() {
			fmt.Println("completed")
		})
		fmt.Println()*/
	/*
		使用Range
	*/
	obserRange := rxgo.Range(0, 5)
	printOberser(obserRange)

	/*
		Repeat
		在已存在的 Observable 对象上调用Repeat，可以实现每隔指定时间，重复一次该序列，一共重复指定次数
	*/
	/* obserRepeat := obserRange.Repeat(3,rxgo.WithDuration(3*time.Second))
	printOberser(obserRepeat)*/

	/*
		Start
		可以给Start方法传入[]rxgo.Supplier作为参数，它可以包含任意数量的rxgo.Supplier类型。rxgo.Supplier的底层类型为：
		// rxgo/types.go
		var Supplier func(ctx context.Context) rxgo.Item
		Observable 内部会依次调用这些rxgo.Supplier生成rxgo.Item：
	*/
	obserStart := rxgo.Start([]rxgo.Supplier{Supplier1, Supplier2, Supplier3})
	printOberser(obserStart)

	/*
		Observable 分类
	     根据数据在何处生成，Observable 被分为 Hot 和 Cold 两种类型（类比热启动和冷启动）。
		数据在其它地方生成的被成为 Hot Observable。
		相反，在 Observable 内部生成数据的就是 Cold Observable。
		上面创建的是 Hot Observable。但是有个问题，第一次Observe()消耗了所有的数据，第二个就没有数据输出了。

	而 Cold Observable 就不会有这个问题，因为它创建的流是独立于每个观察者的。即每次调用Observe()都创建一个新的 channel。
		我们使用Defer()方法创建 Cold Observable，它的参数与Create()方法一样。
	*/
	obserDefer := rxgo.Defer([]rxgo.Producer{func(_ context.Context, ch chan<- rxgo.Item) {
		for i := 0; i < 3; i++ {
			ch <- rxgo.Of(i)
		}
	}})
	printOberser(obserDefer)
	printOberser(obserDefer)

}
