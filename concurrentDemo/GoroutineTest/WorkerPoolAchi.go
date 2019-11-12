package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
缓冲信道的重要应用之一就是实现工作池。
一般而言，工作池就是一组等待任务分配的线程。一旦完成了所分配的任务，这些线程可继续等待任务的分配。
我们会使用缓冲信道来实现工作池。我们工作池的任务是计算所输入数字的每一位的和。例如，如果输入 234，结果会是 9（即 2 + 3 + 4）。向工作池输入的是一列伪随机数。
我们工作池的核心功能如下：
创建一个 Go 协程池，监听一个等待作业分配的输入型缓冲信道。
将作业添加到该输入型缓冲信道中。
作业完成后，再将结果写入一个输出型缓冲信道。
从输出型缓冲信道读取并打印结果。
*/

type Job struct {
	id        int
	randomNum int
}
type Result struct {
	job         Job
	sumOfDigits int
}

func sumCount(num int) int {
	sum := 0
	no := num
	for no != 0 {
		digit := no % 10
		sum += digit
		no = no / 10
	}
	time.Sleep(1 * time.Second)
	return sum
}

var jobsChan = make(chan Job, 10)
var resultsChan = make(chan Result, 10)

func worker(wg *sync.WaitGroup) {
	for job := range jobsChan {
		output := Result{job: job, sumOfDigits: sumCount(job.randomNum)}
		resultsChan <- output
	}
	wg.Done()
}
func createPoolOfWorker(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(resultsChan)
}

//把作业分配给工作者。
func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randint := rand.Int()
		job := Job{id: i, randomNum: randint}
		jobsChan <- job
	}
	close(jobsChan)
}

//从信道中读取输入打印输出
func getResults(done chan bool) {
	for result := range resultsChan {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomNum, result.sumOfDigits)
	}
	done <- true
}
func main() {
	startTime := time.Now()
	noOfJobs := 100
	go allocate(noOfJobs)
	done := make(chan bool)
	go getResults(done)
	noOfWorkers := 100
	createPoolOfWorker(noOfWorkers)
	<-done
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Println("total time taken ", duration.Seconds(), "seconds")
}
