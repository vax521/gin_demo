package main

import (
	"fmt"
	"runtime"
)

/*
https://mp.weixin.qq.com/s?__biz=MzAxMTA4Njc0OQ==&mid=2651438591&idx=2&sn=bfa676374883c5b6e2df955d93b717db&chksm=80bb630db7ccea1bfcd3a3b97bd39eafaaa59281e91cb80bb1c428a49f7fd32fe47f54849c55&mpshare=1&scene=1&srcid=&sharer_sharetime=1582197398329&sharer_shareid=26f90233aaca457c116a35751fe88b75&exportkey=AVZb3qvSDY6Okf02wwSK4wY%3D&pass_ticket=2d6pr8Qx2j6BiTZ3CG7EErJwBUHK%2FvNqtE28Kitj3oF1zuoXWfeSMXAN%2ByVUrCzl#rd
相同的一个 ZeroEvenOdd 类实例将会传递给三个不同的线程：

线程 A 将调用 zero()，它只输出 0 。
线程 B 将调用 even()，它只输出偶数。
线程 C 将调用 odd()，它只输出奇数。
每个线程都有一个 printNumber 方法来输出一个整数。请修改给出的代码以输出整数序列 010203040506... ，其中序列的长度必须为 2n。*/

type ZeroOddEven struct {
	n              int
	chanZeroToOdd  chan struct{}
	chanZeroToEven chan struct{}
	chanOddToZero  chan struct{}
	chanEvenToZero chan struct{}
	chanZeroToend  chan struct{}
}

func PrintNum(n int) {
	fmt.Printf("%d\n", n)
}

func (zeo *ZeroOddEven) Zero(printNum func(int)) {
	for i := 0; i < zeo.n; i++ {
		select {
		case <-zeo.chanOddToZero:
			printNum(0)
			zeo.chanZeroToEven <- struct{}{}
		case <-zeo.chanEvenToZero:
			printNum(0)
			zeo.chanZeroToOdd <- struct{}{}
		default:
			runtime.Gosched()
			i--
		}
	}
	if zeo.n%2 == 0 {
		<-zeo.chanEvenToZero //等待even结束
	} else {
		<-zeo.chanOddToZero //等待odd结束
	}
	zeo.chanZeroToend <- struct{}{} //通知zero结束
}

func (zeo *ZeroOddEven) Odd(printNum func(int)) {
	oddUpper := ((zeo.n + 1) - (zeo.n+1)%2) - 1
	for i := 1; i <= oddUpper; i = i + 2 {
		<-zeo.chanZeroToOdd
		printNum(i)
		zeo.chanOddToZero <- struct{}{}
	}
}

func (zeo *ZeroOddEven) Even(printNum func(int)) {
	evenUpper := zeo.n - zeo.n%2
	for i := 2; i <= evenUpper; {
		<-zeo.chanZeroToEven
		printNum(i)
		i += 2
		zeo.chanEvenToZero <- struct{}{}
	}
}

func main() {
	var printZeroOddeven = func(testNum int) {
		var zeo = &ZeroOddEven{
			n:              testNum,
			chanEvenToZero: make(chan struct{}),
			chanZeroToEven: make(chan struct{}),
			chanOddToZero:  make(chan struct{}),
			chanZeroToOdd:  make(chan struct{}),
			chanZeroToend:  make(chan struct{}),
		}
		go func() { zeo.chanEvenToZero <- struct{}{} }() //点燃火种
		go zeo.Zero(PrintNum)
		go zeo.Odd(PrintNum)
		go zeo.Even(PrintNum)
		<-zeo.chanZeroToend
		fmt.Println()
	}

	printZeroOddeven(15)
	/*for testNum := range [14]int{}{
	          fmt.Printf("%2d\n",testNum)
	          printZeroOddeven(testNum)
		}*/
}
