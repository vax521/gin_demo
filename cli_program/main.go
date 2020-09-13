package main

import (
	"backend/cli_program/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute error:%v\n", err)
	}
}
