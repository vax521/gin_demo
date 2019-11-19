package main

import (
	"fmt"
	"reflect"
)

func main() {
	type cat struct {
		Name string
		Type int `json:"age,omitempty"`
	}

	ins := cat{Name: "mimi", Type: 1}
	typeofCat := reflect.TypeOf(ins)
	for i := 0; i < typeofCat.NumField(); i++ {
		fildType := typeofCat.Field(i)
		fmt.Printf("name:%v tag:%v\n", fildType.Name, fildType.Tag)
	}

}
