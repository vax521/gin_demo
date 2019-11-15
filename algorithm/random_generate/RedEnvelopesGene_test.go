package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/**
  单元测试文件与被测试文件需要在一个包下面
*/
func TestRedPacket(t *testing.T) {
	//随机发一万次红包
	for i := 0; i < 10000; i++ {
		resultArray := RedPacket(10, float64(100), float64(0.01))
		fmt.Println(resultArray)
		for _, v := range resultArray {
			if v < 0.01 {
				t.Errorf("%f小于红包最小要求", v)
			}
		}
		expected := 100.00
		actual := sliceSum(resultArray)

		if actual != expected {
			t.Errorf("Expect %f, but got %f!", expected, actual)
		}
	}
}

func TestRedPacket2(t *testing.T) {
	startTime := time.Now().Second()
	//随机发一万次红包
	for i := 0; i < 10000; i++ {
		resultArray := RedPacket2(10, float64(100), float64(0.01))
		fmt.Println(resultArray)
		for _, v := range resultArray {
			if v < 0.01 {
				t.Errorf("%f小于红包最小要求", v)
			}
		}
		expected := 100.00
		actual := sliceSum(resultArray)
		if actual != expected {
			t.Errorf("Expect %f, but got %f!", expected, actual)
		}
		time.Sleep(time.Millisecond)
	}
	fmt.Printf("程序执行时间：%d s\n", time.Now().Second()-startTime)
}

func RedPacket222(wg *sync.WaitGroup) {
	defer wg.Done()
	resultArray := RedPacket2(10, float64(100), float64(0.01))
	fmt.Println(resultArray)
	for _, v := range resultArray {
		if v < 0.01 {
			fmt.Printf("%f小于红包最小要求", v)
		}
	}
	expected := 100.00
	actual := sliceSum(resultArray)

	if actual != expected {
		fmt.Printf("Expect %f, but got %f!", expected, actual)
	}
	time.Sleep(time.Millisecond)
}

//并行测试
//增加time.Sleep以后，性能差异显著
func TestRedPacket3(t *testing.T) {
	startTime := time.Now().Second()
	wg := sync.WaitGroup{}
	//随机发一万次红包
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go RedPacket222(&wg)
	}
	wg.Wait()
	fmt.Printf("程序执行时间：%d s\n", time.Now().Second()-startTime)
}

//匿名函数写法
func TestRedPacket4(t *testing.T) {
	startTime := time.Now().Second()
	wg := sync.WaitGroup{}
	//随机发一万次红包
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(wgTemp *sync.WaitGroup) {
			defer wg.Done()
			resultArray := RedPacket2(10, float64(100), float64(0.01))
			fmt.Println(resultArray)
			for _, v := range resultArray {
				if v < 0.01 {
					fmt.Printf("%f小于红包最小要求", v)
				}
			}
			expected := 100.00
			actual := sliceSum(resultArray)

			if actual != expected {
				fmt.Printf("Expect %f, but got %f!", expected, actual)
			}
			time.Sleep(time.Millisecond)
		}(&wg)
	}
	wg.Wait()
	fmt.Printf("程序执行时间：%d s\n", time.Now().Second()-startTime)
}
