package main

import (
	"fmt"
	"github.com/liyue201/gostl/ds/array"
)

func main() {
	arr := array.New(5)
	for i := 0; i < arr.Size(); i++ {
		arr.Set(i, i+1)
	}
	for i := 0; i < arr.Size(); i++ {
		fmt.Println(arr.At(i))
	}
	for iter := arr.Begin(); iter.IsValid(); iter.Next() {
		fmt.Println(iter.Value())
	}
}
