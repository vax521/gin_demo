package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	//声明一个d等待组
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.github.com/",
		"http://www/qiniu.com/",
		"http://www.baidu.com/",
	}
	//遍历地址
	for _, url := range urls {
		wg.Add(1)
		//开启一个并发
		go func(url string) {
			//函数完成时将等待组数量减一
			defer wg.Done()
			_, err := http.Get(url)
			fmt.Println(url, err)
		}(url)
	}
	//等待所有任务完成
	wg.Wait()
	fmt.Println("over")

}
