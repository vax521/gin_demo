package main

import (
	"fmt"
	"log"
)

func main() {
	cache := newCache()
	cache.set("key", []byte("the value"))
	value, err := cache.get("key")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(value))
}
