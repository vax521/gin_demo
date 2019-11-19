package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	str1 := "I am a boy"
	fmt.Println(len(str1))

	//go语言的字符串以UTF-8存储，每个中文占三个字节
	str2 := "忍者无敌"
	fmt.Println(len(str2))

	//计算UTF-8字符的个数
	fmt.Println(utf8.RuneCountInString(str2))

	//字符串遍历
	theme := "狙击start"
	// 按ascii遍历
	for i := 0; i < len(theme); i++ {
		fmt.Printf("ascii:%c %d\n", theme[i], theme[i])
	}
	// 按unicode遍历
	for _, s := range theme {
		fmt.Printf("Unicode:%c %d\n", s, s)
	}

	tracerString := "死神来了,死神bye bye"
	comma := strings.Index(tracerString, ",")
	fmt.Println(comma)
	fmt.Println(tracerString[comma:])
	pos := strings.Index(tracerString[comma:], "死神")
	fmt.Println(tracerString[comma+pos:])
	fmt.Println(comma, pos, tracerString[comma+pos:])

	//修改字符串(不可变对象）
	angel := "Heros never die"
	angelBytes := []byte(angel)
	for i := 5; i <= 10; i++ {
		angelBytes[i] = ' '
	}
	fmt.Println(string(angelBytes))

	// 拼接字符串
	hammer := "锤子"
	sickel := "died"
	//声明一个字节缓冲
	var stringBuiler bytes.Buffer
	stringBuiler.WriteString(hammer)
	stringBuiler.WriteString(sickel)
	fmt.Println(stringBuiler.String())

	//格式化
	var process = 2
	var target = 8
	title := fmt.Sprintf("已采集%d个，还需要%d个完成任务", process, target)
	fmt.Println(title)
	pi := 3.1415926
	variant := fmt.Sprintf("%v %v %v", "基地", pi, true)
	fmt.Println(variant)
	//匿名结构体申明，并赋予初值
	profile := &struct {
		Name string
		HP   int
	}{
		Name: "rat",
		HP:   150,
	}
	fmt.Printf("使用'%%+v:'%+v:\n", profile)
	fmt.Printf("使用'%%#v'%#v:\n", profile)
	fmt.Printf("使用'%%T':%T\n", profile)

}
