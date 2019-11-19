package main

import "fmt"

// 通过指针交换值
func swap(a, b *int) {
	t := *a
	*a = *b
	*b = t
}
func main() {
	var cat int = 1
	var str string = "hello"
	// 获取变量的地址
	fmt.Printf("%p %p\n", &cat, &str)
	// 从指针获取指针指向的地址
	var house = "maliue Point 1080,96025"
	ptr := &house
	fmt.Printf("ptr type:%T\n", ptr)
	fmt.Printf("address:%p\n", ptr)
	value := *ptr
	fmt.Printf("value type:%T\n", value)
	fmt.Printf("value:%s\n", value)

	x, y := 1, 2
	swap(&x, &y)
	fmt.Println(x, y)

	// 通过new()创建指针
	str1 := new(string)
	*str1 = "ninja"
	fmt.Println(*str1)

}
