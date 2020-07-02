package main

import (
	"fmt"
	"time"
)

type LeakyBucket struct {
	cap            int
	interval       time.Duration
	dropNums       int //固定时间内的访问次数
	lastAccessTime time.Time
}

func (bucket *LeakyBucket) accessControl() bool {
	now := time.Now()
	pastTime := now.Sub(bucket.lastAccessTime)
	leaks := int(float64(pastTime) / float64(bucket.interval))
	if leaks > 0 {
		if bucket.dropNums <= leaks {
			bucket.dropNums = 0
		} else {
			bucket.dropNums -= leaks
		}
	}
	bucket.lastAccessTime = now

	//允许访问范围内
	if bucket.dropNums < bucket.cap {
		bucket.dropNums++
		return true
	} else {
		return false
	}

}

func main() {
	bucket := &LeakyBucket{
		cap:      10,
		interval: time.Second,
	}

	for i := 0; i < 12; i++ {
		allowed := bucket.accessControl()
		fmt.Println("i", i)
		fmt.Println("i", allowed)
		time.Sleep(time.Millisecond * 500)
	}

	time.Sleep(time.Second * 3) // 模拟3秒中时间内没有访问
	fmt.Println("空档期走完")

	for j := 0; j < 20; j++ {
		fmt.Println("BEFORE", bucket.dropNums)
		allowed := bucket.accessControl()
		fmt.Println("AFTER", bucket.dropNums)
		fmt.Println("j", j)
		fmt.Println("j", allowed)
		time.Sleep(time.Millisecond * 500)
	}
}
