package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	message := "Away from keyboard.http://golang.org/"
	//编码信息
	encodedMess := base64.StdEncoding.EncodeToString([]byte(message))
	fmt.Println(encodedMess)
	//解码消息
	data, err := base64.StdEncoding.DecodeString(encodedMess)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}
}
