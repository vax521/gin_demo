package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)

func hash_file(path string) []byte {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	return h.Sum(nil)
}

func main() {
	data := []byte("this is test, hello world, keep coding")
	fmt.Printf("sha1.Sum(data):%x\n", sha1.Sum(data))

	h := sha1.New()
	io.WriteString(h, "this is test, hello world, keep coding")
	fmt.Printf("%v\n", h.Sum(nil))

	fmt.Printf("%v\n", hash_file("./algorithm/hash_function/file.txt"))

	fmt.Printf("sha256.Sum224(data):%v\n", sha256.Sum224(data))
	fmt.Printf("sha256.Sum256(data):%v\n", sha256.Sum256(data))
}
