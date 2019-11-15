package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%2.2f", value), 64)
	return value
}

/**
returnSlice := make([]float64, n)
	for k, v := range redPocketResult {
		returnSlice[k] = Decimal(minMoney + v)
	}
	lastPacket := allMoney - sliceSum(returnSlice)
	returnSlice[n-1] = Decimal(lastPacket + minMoney)
	return returnSlice[:]
*/
func RedPacket2(n int, allMoney float64, minMoney float64) []float64 {
	redPocketResult := make([]float64, n)
	//可分配的钱数为 总钱数-n*每个钱包的最小数
	canDisMoney := Decimal(allMoney - float64(n)*minMoney)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n-1; i++ {
		coefficient := Decimal(rand.Float64()) //生成0.0-1.0随机整数
		fmt.Printf("随机系数：%.2f\n", coefficient)
		if coefficient == 0.0 || coefficient == 1.0 {
			coefficient = 0.5
		}
		redPocketResult[i] = Decimal(coefficient * (canDisMoney / float64(n-i) * 2))
		canDisMoney = Decimal(canDisMoney - redPocketResult[i])
	}
	returnSlice := make([]float64, n)
	for k, v := range redPocketResult {
		returnSlice[k] = Decimal(minMoney + v)
	}
	lastPacket := allMoney - sliceSum(returnSlice)
	returnSlice[n-1] = Decimal(lastPacket + minMoney)
	return returnSlice[:]
}

/**
  n : 红包个数
  allMoney ： 总钱数
   minMoney ：每个红包的最小钱数
*/
func RedPacket(n int, allMoney float64, minMoney float64) []float64 {
	redPocketResult := make([]float64, n)
	//可分配的钱数为 总钱数-n*每个钱包的最小数
	canDisMoney := Decimal(allMoney - float64(n)*minMoney)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n-1; i++ {
		coefficient := Decimal(rand.Float64()) //生成0.0-1.0随机整数
		fmt.Printf("随机系数：%f\n", coefficient)
		if coefficient == 0.0 || coefficient == 1.0 {
			coefficient = 0.2
		}
		fmt.Printf("实际系数：%f\n", coefficient)
		redPocketResult[i] = Decimal(coefficient * canDisMoney)
		canDisMoney = Decimal(canDisMoney - redPocketResult[i])
	}
	returnSlice := make([]float64, n)
	for k, v := range redPocketResult {
		returnSlice[k] = Decimal(minMoney + v)
	}
	lastPacket := allMoney - sliceSum(returnSlice)
	returnSlice[n-1] = Decimal(lastPacket + minMoney)
	return returnSlice[:]
}

func sliceSum(resultArray []float64) float64 {
	sum := 0.0
	for _, v := range resultArray {
		sum = sum + v
	}
	return Decimal(sum)
}

func main() {
	resultArray := RedPacket2(10, float64(100), float64(0.01))
	fmt.Println(sliceSum(resultArray))
	fmt.Println(resultArray)
}
