sync 包
sync 包包含了对低级别内存访问同步最有用的并发原语，是 “内存访问同步” 的最有利工具，也是传统并发模型解决临界区问题的常用工具。

WaitGroup
WaitGroup 是等待一组并发操作完成的方法，包含了三个函数：

func (wg *WaitGroup) Add(delta int)
func (wg *WaitGroup) Done()
func (wg *WaitGroup) Wait()
复制代码
其中，Add() 用来添加 goroutine 的个数，Done() 是 goroutine 用来表明执行完成并退出，将计数减一，而 Wait() 用来等待所有 goroutine 退出。

用法如下：

func main() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Printf("goroutine 结束\n")
	}()

	wg.Wait()
}
复制代码
需要注意的是，Add() 方法需要在 goroutine 之前执行。

互斥锁和读写锁
互斥是保护程序中临界区的一种方式。一个互斥锁只能同时被一个 goroutine 锁定，其它 goroutine 将阻塞直到互斥锁被解锁（重新争抢对互斥锁的锁定）。

用法如下：

func main() {
	var lock sync.Mutex
	var count int
	var wg sync.WaitGroup

	wg.Add(1)
	// count 加 1
	go func() {
		defer wg.Done()
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Println("count=", count)
	}()

	// count 减 1
	wg.Add(1)
	go func() {
		defer wg.Done()
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Println("count=", count)
	}()

	wg.Wait()
	fmt.Println("count=", count)
}
复制代码
需要注意的是，在 goroutine 里用 defer 来调用 Unlock 是个常见的习惯用法，确保了即使出现了 panic，调用也总是执行，防止出现死锁。

读写锁在概念上跟互斥锁是一样的：保护对内存的访问，读写锁让你对内存有更多的控制。读写锁与互斥锁最大的不同就是可以分别对读、写进行锁定。一般用在大量读操作、少量写操作的情况。

读写锁的 Lock() 和 Unlock() 是对写操作的锁定和解锁；Rlock() 和 RUnlock() 是对读操作的锁定和解锁，需要配对使用。而读锁和写锁的关系：

同时只能有一个 goroutine 能够获得写锁定。
同时可以有任意多个 gorouinte 获得读锁定。
同时只能存在写锁定或读锁定（读和写互斥）。
Channel
Channel 是 CSP 派生的同步原语之一，是 Golang 推崇的 “使用通信来共享内存，而不是通过共享内存来通信” 理念的最有利的工具。

Channel 的基本使用这里不展开讲，但对不同状态下的 Channel 不同操作的结果做一个总结：

操作	Channel 状态	结果
Read	nil	阻塞
打开非空	输出值
打开但空	阻塞
关闭	<默认值>, false
只写	编译错误
Write	nil	阻塞
打开但填满	阻塞
打开不满	写入值
关闭	panic
只读	编译错误
Close	nil	panic
打开非空	关闭 Channel; 读取成功，直到 Channel 耗尽，读取产生值的默认值
打开但空	关闭 Channel；读到生产者的默认值
关闭	panic
只读	编译错误
for-select
select 语句是将 Channel 绑定在一起的粘合剂，能够让一个 goroutine 同时等待多个 Channel 达到准备状态。

select 语句是针对 Channel 的操作，语法上看上去与 switch 很像，但不同的是，select 块中的 case 语句没有测试顺序，如果没有满足任何条件，执行也不会失败。用法如下：

var c1, c2 <-chan interface{}
select {
  case <- c2:
    // 某段逻辑
  case <- c2:
    // 某段逻辑
}
复制代码
上面这个 select 控制结构会等待所有 case 条件语句任意一个的返回，无论哪一个返回都会立刻执行 case 中的代码，不过如果了 select 中的两个 case 同时被触发，就会随机选择一个 case 执行。

for-select 是一个很常见的用法，通常在 “向 Channel 发送迭代变量” 和 “循环等待停止” 两种情况下会用到，用法如下：

向 Channel 发送迭代变量：

func main() {
	c := make(chan int, 3)
	for _, s := range []int{1, 2, 3} {
		select {
		case c <- s:
		}
	}
}
复制代码
循环等待停止：

// 第一种
for {
  select {
  case <- done:
    return
  default:
    // 进行非抢占式任务
  }
}
// 第二种
for {
  select {
  case <- done:
    return
  default:
  }
  // 进行非抢占式任务
}
复制代码
第一种是指，当我们输入 select 语句时，如果完成的 Channel 尚未关闭，我们将执行 default 语句；第二种是指，如果已经完成的 Channel 未关闭，我们将退出 select 语句并继续执行 for 循环的其余部分。

done channel
虽然 goroutine 廉价且易于利用，运行时可以将多个 goroutine 复用到任意数量的操作系统线程，但我们需要知道的是 goroutine 是需要消耗资源的，并且是不会被运行时垃圾回收的。如果出现 goroutine 泄露的情况，严重的时候会导致内存利用率的下降。

而 done channel 就是防止 goroutine 泄露的利器。用 done channel 在父子 goroutine 之间建立一个 “信号通道”，父 goroutine 可以将该 channel 传递给子 goroutine ，然后在想要取消子 goroutine 的时候关闭该 channel。用法如下：

func main() {
	doneChan := make(chan interface{})

	go func(done <-chan interface{}) {
	   for {
		  select {
		  case <-done:
		    return
		  default:
		  }
		}
	}(doneChan)

	// 父 goroutine 关闭子 goroutine
	close(doneChan)
}
复制代码
确保 goroutine 不泄露的方法，就是规定一个约定：如果 goroutine 负责创建 goroutine，它也负责确保它可以停止 goroutine。

Context 包
Context 包是专门用来简化对于处理单个请求的多个 goroutine 之间与请求域的数据、取消信号、截止时间等相关操作，这些操作可能涉及多个 API 调用。Context 包的目的主要有两个：提供一个可以取消你的调用图中分支的 API，提供用于通过呼叫传输请求范围数据的数据包。

如果使用 Context 包，那么位于顶级并发调用下游的每个函数都会将 context 作为其第一个参数。

Context 的类型如下：

type Context interface {
  Deadline() (deadline time.Time, ok bool)
  Done() <-chan struct{}
  Err() error
  Value(key interface{}) interface{}
}
复制代码
其中，Deadline 函数用于指示在一定时间后 goroutine 是否会被取消；Done 方法返回当我们的函数被抢占时关闭的 Channel；Err 方法返回取消的错误原因，因为什么 Context 被取消；Value 函数返回与此 Context 关联的 key 或 nil。

Context 虽然是个接口，但是我们在使用它的时候并不需要实现，context 包内置的两个方法来创建上下文的实例：

func Background() Context
func TODO() Context
复制代码
Background 主要用于 main 函数、初始化以及测试代码中，作为Context 这个树结构的最顶层的 Context，不能被取消；TODO，如果我们不知道该使用什么 Context 的时候，可以使用这个，但是实际应用中，暂时还没有使用过这个 TODO。

然后以此作为最顶层的父 Context，衍生出子 Context 启动调用链。而这些 Context 对象形成了一棵树，当父 Context 对象被取消时，它的所有子 Context 都会被取消。context 包还提供了一系列函数用以产生子 Context：

func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
func WithValue(parent Context, key, val interface{}) Context
复制代码
其中，WithCancel 返回一个新的 Context，在调用返回的 cancel 函数时关闭其 done channel；WithDeadline 返回一个新的 Context，当机器的时钟超过给定的最后期限时，它关闭完成的 channel；WithTimeout 返回一个新的 Context，在给定的超时时间后关闭其完成的 channel；WithValue 生成一个绑定了一个键值对数据的 Context，这个绑定的数据可以通过 Context.Value 方法访问到。

下面来看使用方法：

WithCancel
func main() {
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
复制代码
WithDeadline
func main() {
	d := time.Now().Add(1 * time.Second)
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

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

	wg.Wait()
}
复制代码
WithTimeout
func main() {
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

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

	wg.Wait()
}
复制代码
WithValue
func main() {
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

	time.Sleep(5*time.Second)
	cancel()
	wg.Wait()
}
# 定时器
我们刚才其实已经接触了信道作为定时器, time包里的After会制作一个定时器。

https://blog.csdn.net/u013474436/article/details/86651335