package main

import "fmt"

func main() {
	//常量
	const pi = 3.1415926
	const e = 2.71
	const (
		name = "xing"
		er   = "3"
	)
	const size = 4
	//var arr [size]int

	//使用常量配合iota实现枚举
	type Weapon int
	const (
		Arrow Weapon = iota
		Shuriken
		Sniper
		Rifle
		Blower
	)
	fmt.Println(Arrow, Shuriken, Sniper, Rifle, Blower)
	var weapon Weapon = Blower
	fmt.Println(weapon)
	//通过移位操作实现常量值生成器

	const (
		FlagNone = 1 << iota
		FlagRed
		FlagGreen
		FlagBlue
	)
	fmt.Printf("%d %d %d\n", FlagRed, FlagGreen, FlagBlue)
	//二进制表示
	fmt.Printf("%b %b %b\n", FlagRed, FlagGreen, FlagBlue)
}
