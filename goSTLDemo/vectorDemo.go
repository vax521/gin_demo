package main

import (
	"fmt"
	"github.com/liyue201/gostl/ds/vector"
)

func main() {
	v := vector.New()
	v.PushBack(3)
	v.PushBack(5)
	v.PushBack(7)
	v.PushBack(9)

	for i := 0; i < v.Size(); i++ {
		fmt.Println(v.At(i))
	}

	for iter := v.Begin(); iter.IsValid(); iter.Next() {
		fmt.Println(iter.Value())
		fmt.Println(iter.Position())
	}
}
