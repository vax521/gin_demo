package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "监控退出...")
			return
		default:
			fmt.Println(name, "goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}

func testCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx, "[监控1]")
	go watch(ctx, "[监控2]")
	go watch(ctx, "[监控3]")
	time.Sleep(10 * time.Second)
	fmt.Println("all 监控 完毕...")
	cancel()
	time.Sleep(time.Second * 5)
}

func commonDmeo() {
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Err:", ctx.Err())
				return
			default:
			}
		}
	}(ctx)

	cancel()
	wg.Wait()
}

func testDeadline() {
	deadline := time.Now().Add(1 * time.Second)
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	wg.Add(1)
	defer cancel()
	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("err:", ctx.Err())
				return
			default:
				fmt.Println("监控中....")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}(ctx)
	wg.Wait()
}

func testTimeout() {
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("err:", ctx.Err())
				return
			default:
				fmt.Println("监控中....")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}(ctx)
	wg.Wait()
}
func testValue() {
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	valueCtx := context.WithValue(ctx, "key", "add value")

	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Err:", ctx.Err())
				return
			default:
				fmt.Println(ctx.Value("key"))
				time.Sleep(1 * time.Second)
			}
		}
	}(valueCtx)

	time.Sleep(5 * time.Second)
	cancel()
	wg.Wait()
}
func main() {
	//testDeadline()
	//testTimeout()
	testValue()
}
