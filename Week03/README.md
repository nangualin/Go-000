作业基本是写完了。但是并不满意。由于家里人住院，在那边陪护。希望后面的时间努力学习再跟上进度。

## 学习笔记
参考自go圣经：
http://shouce.jb51.net/gopl-zh/ch8/ch8-02.html
8.2. 示例: 并发的Clock服务

```

服务端
// Clock1 is a TCP server that periodically writes the time.
package main

import (
    "io"
    "log"
    "net"
    "time"
)

func main() {
    listener, err := net.Listen("tcp", "localhost:8000")
    if err != nil {
        log.Fatal(err)
    }

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Print(err) // e.g., connection aborted
            continue
        }
        // handleConn(conn) // handle one connection at a time 由于里面是个死循环，没法结束以便让listener继续执行accept。，因此第二个连接请求他没有取出来出来。
        go handleConn(conn) // 可以接收多个请求了。   有了go之后。外层循环继续。因此能继续拿到新的请求以处理。
    }
}

func handleConn(c net.Conn) {
    defer c.Close()
    for {
        _, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
        if err != nil {
            return // e.g., client disconnected
        }
        time.Sleep(1 * time.Second)
    }
}
===============================
客户端：
// Netcat1 is a read-only TCP client.
package main

import (
    "io"
    "log"
    "net"
    "os"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:8000")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()
    mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
    if _, err := io.Copy(dst, src); err != nil {
        log.Fatal(err)
    }
}

```

### channel
读与写
ch = make(chan int)    // unbuffered channel 无缓存
ch = make(chan int, 0) // unbuffered channel 无缓存
ch = make(chan int, 3) // buffered channel with capacity 3 有缓存

带缓存的Channels
ch = make(chan string, 3)

由于channel的特性。留有容量。则可以避免一些goroutines泄露
如下代码：
func mirroredQuery() string {
    responses := make(chan string, 3)
    go func() { responses <- request("asia.gopl.io") }()
    go func() { responses <- request("europe.gopl.io") }()
    go func() { responses <- request("americas.gopl.io") }()
    return <-responses // return the quickest response
}

func request(hostname string) (response string) { /* ... */ }

关于close(ch) panic出现的情况
试图重复关闭一个channel将导致panic异常，试图关闭一个nil值的channel也将导致panic异常
不管一个channel是否被关闭，当它没有被引用时将会被Go语言的垃圾自动回收器回收
当一个channel被关闭后，再向该channel发送数据将导致panic异常。当一个被关闭的channel中已经发送的数据都被成功接收后，后续的接收操作将不再阻塞，它们会立即返回一个零值。

x, ok := <-naturals
它多接收一个结果，多接收的第二个结果是一个布尔值ok，ture表示成功从channels接收到值，false表示channels已经被关闭并且里面没有值可接收。



## channel的应用
串联的Channels（Pipeline）

#### 单方向的Channel
Go语言的类型系统提供了单方向的channel类型，分别用于只发送或只接收的channel。类型chan<- int表示一个只发送int的channel，只能发送不能接收。相反，类型<-chan int表示一个只接收int的channel，只能接收不能发送。（箭头<-和关键字chan的相对位置表明了channel的方向。）这种限制将在编译期检测。

#### channel测量容量和库存
在某些特殊情况下，程序可能需要知道channel内部缓存的容量，可以用内置的cap函数获取：

fmt.Println(cap(ch)) // "3"
同样，对于内置的len函数，如果传入的是channel，那么将返回channel内部缓存队列中有效元素的个数。因为在并发程序中该信息会随着接收操作而失效，但是它对某些故障诊断和性能优化会有帮助。

fmt.Println(len(ch)) // "2"

### 消息机制
不关心值，关心信号时。基于channels发送消息有两个重要方面。首先每个消息都有一个值，但是有时候通讯的事实和发生的时刻也同样重要。当我们更希望强调通讯发生的时刻时，我们将它称为消息事件。有些消息事件并不携带额外的信息，它仅仅是用作两个goroutine之间的同步，这时候我们可以用struct{}空结构体作为channels元素的类型，虽然也可以使用bool或int类型实现同样的功能，done <- 1语句也比done <- struct{}{}更短。

每个channel都有一个特殊的类型，也就是channels可发送数据的类型。一个可以发送int类型数据的channel一般写为chan int。

参考： http://shouce.jb51.net/gopl-zh/ch8/ch8-04.html

##happen before 概念
当我们说x事件在y事件之前发生（happens before），我们并不是说x事件在时间上比y时间更早；我们要表达的意思是要保证在此之前的事件都已经完成了，例如在此之前的更新某些变量的操作已经完成，你可以放心依赖这些已完成的事件了。
