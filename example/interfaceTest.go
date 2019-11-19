package main

import "fmt"

type SalaryCalculator interface {
	CalculateSalary() int
}

type primary struct {
	empId    int
	basicPay int
}
type senior struct {
	empId    int
	basicpay int
	pf       int
}

//结构体实现接口
func (p primary) CalculateSalary() int {
	return p.basicPay
}

func (s senior) CalculateSalary() int {
	return s.basicpay + s.pf
}

//参数是SalaryCalculator 接口的切片
func getTotalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("total expense of each month:%d", expense)
}

func main() {
	pri1 := primary{
		empId:    1,
		basicPay: 5000,
	}
	senior1 := senior{
		empId:    2,
		basicpay: 6000,
		pf:       500,
	}
	senior2 := senior{
		empId:    3,
		basicpay: 6000,
		pf:       600,
	}
	getTotalExpense([]SalaryCalculator{pri1, senior1, senior2})
}
