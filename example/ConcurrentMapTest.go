package main

func main() {
	//map在并发情况下读写线程不安全
	m := make(map[int]int)
	go func() {
		for {
			m[1] = 1
		}
	}()
	go func() {
		for {
			_ = m[1]
		}
	}()
	select {}
}
