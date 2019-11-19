package main

import "fmt"

func main() {
	//switch语句的使用
	var a = "hello"
	switch a {
	case "hello":
		fmt.Println(1)
	case "world":
		fmt.Println(2)
	default:
		fmt.Println(0)
	}

	var relation = "mum"
	switch relation {
	case "mum", "daddy":
		fmt.Println("family")
		fallthrough
	default:
		fmt.Println("stranger")
	}
}
