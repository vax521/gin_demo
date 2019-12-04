package main

import "fmt"

type BitMap struct {
	bits []byte
	max  int
}

////初始化一个BitMap
////一个byte有8位,可代表8个数字,取余后加1为存放最大数所需的容量
func NewBitMap(max int) *BitMap {
	bits := make([]byte, (max>>3)+1)
	return &BitMap{bits: bits, max: max}
}

func (bitmap *BitMap) String() string {
	return fmt.Sprintf("%b", bitmap.bits)
}

//添加一个数字到位图
//计算添加数字在数组中的索引index,一个索引可以存放8个数字
//计算存放到索引下的第几个位置,一共0-7个位置
//原索引下的内容与1左移(*2)到指定位置后做或运算
func (bitmap *BitMap) Add(num uint) {
	index := num >> 3
	pos := num & 0x07
	fmt.Println(pos)
	fmt.Println(1 << pos)
	bitmap.bits[index] |= 1 << pos
}

func main() {
	bitmap := NewBitMap(9)
	bitmap.Add(0)
	bitmap.Add(1)
	bitmap.Add(2)
	bitmap.Add(3)
	bitmap.Add(4)
	bitmap.Add(5)
	bitmap.Add(6)
	bitmap.Add(7)
	fmt.Println(bitmap)
	bitmap.Add(8)
	fmt.Println(bitmap)
	bitmap.Add(9)
	fmt.Println(bitmap)
	bitmap.Add(10)
	fmt.Println(bitmap)
	bitmap.Add(15)
	fmt.Println(bitmap)

}
