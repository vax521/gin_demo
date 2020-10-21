# 生成器
在Python中我们可以使用yield关键字来让一个函数成为生成器，在Go中我们可以使用信道来制造生成器(一种lazy load类似的东西)。
当然我们的信道并不是简单的做阻塞主线的功能来使用的哦。
下面是一个制作自增整数生成器的例子，直到主线向信道索要数据，我们才添加数据到信道
```
func xrange() chan int{ // xrange用来生成自增的整数
    var ch chan int = make(chan int)
    go func() { // 开出一个goroutine
        for i := 0; ; i++ {
            ch <- i  // 直到信道索要数据，才把i添加进信道
        }
    }()
    return ch
}

func main() {
    generator := xrange()
    for i:=0; i < 1000; i++ {  // 我们生成1000个自增的整数！
        fmt.Println(<-generator)
    }
}
```
# 服务化
比如我们加载一个网站的时候，例如我们登入新浪微博，我们的消息数据应该来自一个独立的服务，这个服务只负责 返回某个用户的新的消息提醒。
```
func get_notification(user string) chan string{
   /*
    * 此处可以查询数据库获取新消息等等..
    */
    notifications := make(chan string)

    go func() { // 悬挂一个信道出去
        notifications <- fmt.Sprintf("Hi %s, welcome to weibo.com!", user)
    }()

    return notifications
}

func main() {
    jack := get_notification("jack") //  获取jack的消息
    joe := get_notification("joe") // 获取joe的消息

    // 获取消息的返回
    fmt.Println(<-jack)
    fmt.Println(<-joe)
}
```

# 多路复合
上面的例子都使用一个信道作为返回值，可以把信道的数据合并到一个信道的。 不过这样的话，我们需要按顺序输出我们的返回值（先进先出）。
```
func do_stuff(x int) int { // 一个比较耗时的事情，比如计算
    time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond) //模拟计算
    return 100 - x // 假如100-x是一个很费时的计算
}

func branch(x int) chan int{ // 每个分支开出一个goroutine做计算并把计算结果流入各自信道
    ch := make(chan int)
    go func() {
        ch <- do_stuff(x)
    }()
    return ch
}

func fanIn(chs... chan int) chan int {
    ch := make(chan int)

    for _, c := range chs {
        // 注意此处明确传值
        go func(c chan int) {ch <- <- c}(c) // 复合
    }

    return ch
}


func main() {
    result := fanIn(branch(1), branch(2), branch(3))

    for i := 0; i < 3; i++ {
        fmt.Println(<-result)
    }
}
```

# select监听信道
Go有一个语句叫做select，用于监测各个信道的数据流动。
如下的程序是select的一个使用例子，我们监视三个信道的数据流出并收集数据到一个信道中。
```
func foo(i int) chan int {
    c := make(chan int)
    go func () { c <- i }()
    return c
}


func main() {
    c1, c2, c3 := foo(1), foo(2), foo(3)

    c := make(chan int)

    go func() { // 开一个goroutine监视各个信道数据输出并收集数据到信道c
        for {
            select { // 监视c1, c2, c3的流出，并全部流入信道c
            case v1 := <- c1: c <- v1
            case v2 := <- c2: c <- v2
            case v3 := <- c3: c <- v3
            }
        }
    }()

    // 阻塞主线，取出信道c的数据
    for i := 0; i < 3; i++ {
        fmt.Println(<-c) // 从打印来看我们的数据输出并不是严格的1,2,3顺序
    }
}
```

