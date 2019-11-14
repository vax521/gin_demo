package main

import (
	"fmt"
	"testing"
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
