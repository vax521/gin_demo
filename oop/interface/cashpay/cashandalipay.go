package main

import "fmt"

//电子支付
type Alipay struct {
}

func (a *Alipay) CanUseFaceId() {
	fmt.Println("You are using Alipay with faceId")
}

//现金支付方式
type Cash struct {
}

func (c *Cash) Stolen() {

}

//具备刷脸特性的方法
type ContainCanUseFaceId interface {
	CanUseFaceId()
}
type ContainStolen interface {
	Stolen()
}

//打印支付方式具备的特点
func print(payMethod interface{}) {
	switch payMethod.(type) {
	case ContainStolen:
		fmt.Printf("%T may be stolen\n", payMethod)
	case ContainCanUseFaceId:
		fmt.Printf("%T can use faceid\n", payMethod)
	}
}
func main() {
	a := new(Alipay)
	a.CanUseFaceId()
	print(new(Alipay))
	print(new(Cash))
}
