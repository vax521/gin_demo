package main

import "fmt"

/**
方法其实就是一个函数，在 func 这个关键字和方法名中间加入了一个特殊的接收器类型。
接收器可以是结构体类型或者是非结构体类型。接收器是可以在方法的内部访问的。
*/
type Employee struct {
	name     string
	salary   int
	currency string
}

func (e Employee) displayAll() {
	fmt.Printf("name:%s,salary:%s%d\n", e.name, e.currency, e.salary)
}
func (e Employee) displaySalary() {
	fmt.Printf("the salary of %s is %s%d\n", e.name, e.currency, e.salary)
}

/**
    有了函数还需要方法的原因:
 Go 不是纯粹的面向对象编程语言，而且Go不支持类。因此，基于类型的方法是一种实现和类相似行为的途径。
相同的名字的方法可以定义在不同的类型上，而相同名字的函数是不被允许的。
*/

//使用值接收器
func (e Employee) changeName(name string) {
	e.name = name
}

// s使用值针接收器
func (e *Employee) changeSalary(newSalary int) {
	e.salary = newSalary
}

//匿名字段的方法
type address struct {
	city  string
	state string
}

func (add address) fullAddress() {
	fmt.Printf("Full Address:%s-%s\n", add.state, add.city)
}

type Person struct {
	firstName string
	lastName  string
	address
}

/**
  在方法中使用值接收器，在函数中使用值参数
  当一个函数有一个值参数，它只能接受一个值参数。
当一个方法有一个值接收器，它可以接受值接收器和指针接收器。
*/
type rectangle struct {
	width  int
	height int
}

func area(r rectangle) {
	fmt.Printf("the area is %d\n", r.height*r.width)
}
func (r rectangle) area() {
	fmt.Printf("the area is %d\n", r.height*r.width)
}

func main() {
	emp1 := Employee{
		name:     "bob",
		salary:   10000,
		currency: "$",
	}
	emp1.displaySalary()
	emp1.changeName("sam")
	emp1.displayAll()
	(&emp1).changeSalary(20000)
	emp1.displayAll()
	personElon := Person{
		firstName: "Elon",
		lastName:  "Mask",
		address: address{
			city:  "Los Angeles",
			state: "Californai",
		},
	}
	personElon.fullAddress()

	r := rectangle{
		width:  50,
		height: 100,
	}
	area(r)
	r.area()
	p := &r
	p.area()
}
