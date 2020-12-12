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
