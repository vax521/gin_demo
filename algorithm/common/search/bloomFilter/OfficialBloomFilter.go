package bloomFilter

import (
	"fmt"
	"github.com/willf/bloom"
)

func main() {
	n := uint(1000)
	filter := bloom.New(20*n, 5) // load of 20, 5 keys
	filter.Add([]byte("Love"))
	fmt.Println(filter.Test([]byte("Love")))
	fmt.Println(filter.K())
	str1 := "hello,bloom filter!"
	filter.Add([]byte(str1))
	fmt.Println(filter.Test([]byte(str1)))
	fmt.Println(filter.Test([]byte("bloom hello, filter！")))
}
