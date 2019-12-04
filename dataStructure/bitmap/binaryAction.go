package main

import (
	"fmt"
)

/**
与操作：&
1 & 1 = 1
1 & 0 = 0
0 & 1 = 0
0 & 0 = 0

或操作：！
1 | 1 = 1
1 | 0 = 1
0 | 1 = 1
0 & 0 = 0

异或：^
1 ^ 1 = 0
1 ^ 0 = 1
0 ^ 1 = 1
0 ^ 0 = 0

左移：<<
1 << 10 = 1024
1 << 20 = 1M
1 << 30 = 1G

右移：>>
1024 >> 10 = 1
1024 >>1 = 512
1024 >>2 = 256
*/
type users struct {
	name string
	flag uint8
}

// 这里通过位移的方式默认00000 从左边一次为vip,svip,blue,red,yellow
const (
	vip   = 1
	svip  = 1 << 1
	blue  = 1 << 2
	red   = 1 << 3
	yello = 1 << 4
)

// setFlag 用于设置用户开通了哪些特权
func setFlag(user users, isSet bool, typeFlag uint8) users {
	if isSet == true {
		fmt.Printf("%b\n", user.flag)
		fmt.Printf("%b\n", typeFlag)
		user.flag = user.flag | typeFlag
		fmt.Printf("%b\n", user.flag)
		fmt.Println("---------------")
	} else {
		//异或操作
		user.flag = user.flag ^ typeFlag
	}
	return user
}

//isFlag 用于判断用户是否开通某项特权
func isFlag(user users, typeFlag uint8) bool {
	result := user.flag & typeFlag
	return result == typeFlag
}

func binaryTest() {
	var user users
	user.name = "coder"
	user.flag = 0

	/*	//判断用户是否是vip
		result := isFlag(user, vip)
		fmt.Printf("user is Vip:%t\n", result)

		//给用户开通vip,并看用户是否开通vip
		user = setFlag(user, true, vip)

		result = isFlag(user, vip)
		fmt.Printf("user is Vip:%t\n", result)

		//取消用户的vip，并查看用户是否还是vip
		user = setFlag(user, false, vip)
		result = isFlag(user, vip)
		fmt.Printf("user is Vip:%t\n", result)*/

	user = setFlag(user, true, svip)
	user = setFlag(user, true, red)
	user = setFlag(user, true, blue)
	user = setFlag(user, true, yello)
	fmt.Printf("%b\n", user.flag)
	user = setFlag(user, true, yello)
}

func main() {
	binaryTest()
}
