package main

import "fmt"

type Describer interface {
	describe()
}
type familyAddress struct {
	state string
	city  string
}
type person struct {
	name string
	age  int
}

//使用值接受者实现
func (p person) describe() {
	fmt.Printf("name:%s,age:%d\n", p.name, p.age)
}

//使用地址接受者实现
func (fd *familyAddress) describe() {
	fmt.Printf("state:%s-city:%s\n", fd.state, fd.city)
}
func main() {
	var d1 Describer
	p1 := person{
		name: "hii",
		age:  45,
	}
	d1 = p1
	d1.describe()
	p2 := person{
		name: "jiao",
		age:  24,
	}
	d1 = &p2
	d1.describe()
	//总结：使用值接受者声明的方法，既可以用值来调用，也能用指针调用
	var d2 Describer
	a1 := familyAddress{
		state: "shandong",
		city:  "linyi",
	}
	d2 = &a1
	d2.describe()
}
