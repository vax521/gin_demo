package main

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func MarshalJson(v interface{}) (string, error) {
	//准备一个缓冲池
	var b bytes.Buffer
	if err := writeAny(&b, reflect.ValueOf(v)); err == nil {
		return b.String(), nil
	} else {
		return "", err
	}
}

func writeAny(buff *bytes.Buffer, value reflect.Value) error {
	switch value.Kind() {
	case reflect.String:
		buff.WriteString(strconv.Quote(value.String()))
	case reflect.Int:
		buff.WriteString(strconv.FormatInt(value.Int(), 10))
	case reflect.Slice:
		return writeSlice(buff, value)
	case reflect.Struct:
		return writeStruct(buff, value)
	default:
		return errors.New("unsupported kind:" + value.Kind().String())
	}
	return nil
}
func writeSlice(buff *bytes.Buffer, value reflect.Value) error {
	buff.WriteString("[")
	for s := 0; s < value.Len(); s++ {
		sliceValue := value.Index(s)
		writeAny(buff, sliceValue)
		//每个元素尾部加入逗号
		if s < value.Len()-1 {
			buff.WriteString(",")
		}
	}
	buff.WriteString("]")
	return nil
}

func writeStruct(buff *bytes.Buffer, value reflect.Value) error {
	valueType := value.Type()
	buff.WriteString("{")
	for i := 0; i < value.NumField(); i++ {
		fieldValue := value.Field(i)
		fieldType := valueType.Field(i)
		buff.WriteString("\"")
		buff.WriteString(fieldType.Name)
		buff.WriteString("\":")
		writeAny(buff, fieldValue)
		if i < value.NumField()-1 {
			buff.WriteString(",")
		}
	}
	buff.WriteString("}")
	return nil
}
func main() {
	//将结构体转为json文本
	type Skill struct {
		Name  string
		Level int
	}

	type Actor struct {
		Name   string
		Age    int
		Skills []Skill
	}

	a := Actor{
		Name: "cowboy",
		Age:  37,
		Skills: []Skill{
			{Name: "Roll", Level: 1},
			{Name: "Flash", Level: 2},
			{Name: "Timw", Level: 3},
		},
	}
	if result, err := MarshalJson(a); err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}
