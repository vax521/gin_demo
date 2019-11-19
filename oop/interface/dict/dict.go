package main

import "fmt"

type Dictionary struct {
	data map[interface{}]interface{}
}

func (d *Dictionary) Get(key interface{}) interface{} {
	return d.data[key]
}
func (d *Dictionary) Set(key interface{}, value interface{}) {
	d.data[key] = value
}

//遍历字典
func (d *Dictionary) Visit(callback func(k, v interface{}) bool) {
	if callback == nil {
		return
	}
	for k, v := range d.data {
		if !callback(k, v) {
			return
		}
	}
}

// clear all data
func (d *Dictionary) Clear() {
	d.data = make(map[interface{}]interface{})
}

//create a dict
func NewDict() *Dictionary {
	d := &Dictionary{}
	d.Clear()
	return d
}

func main() {
	dict := NewDict()
	dict.Set("MyFactory", 60)
	dict.Set("TerraCraft", 36)
	dict.Set("Dont Hungry", 25)

	favorite := dict.Get("TerraCraft")
	fmt.Println("favorite:", favorite)

	//遍历
	dict.Visit(func(k, v interface{}) bool {
		if v.(int) > 40 {
			fmt.Println(k, "is expensive")
			return true
		}
		fmt.Println(k, "is cheap")
		return true
	})
}
