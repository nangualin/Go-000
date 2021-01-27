package main

import (
	"bufio"
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

//用 Go 实现一个 tcp server ，
// 用两个 goroutine 读写 conn，两个 goroutine 通过 chan 可以传递 message，
// 能够正确退出
var signals chan os.Signal

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	signals = make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT,os.Interrupt)
	fmt.Println("TCP服务已启动，端口是：520")
	listen, err := net.Listen("tcp", ":520")
	if err != nil {
		log.Fatalf("listen error.%v\n", err)
	}
	go func(){
	    sig := <-signals
		cancel()
		if err := listen.Close(); err != nil {
		    fmt.Println("ops,has some errors", err)
		}
		fmt.Println("tcp server close by ",sig.String())
	}()

	serve(ctx,listen)
	fmt.Println("system shutdown")
}

func serve(ctx context.Context,listen net.Listener) error{
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context.error %v\n ", ctx.Err())
			log.Fatalf("ctx error .%v\n", ctx.Err())
			return fmt.Errorf("Received  %s",ctx.Err())
		default:
			conn, err := listen.Accept()
			if err != nil {
				log.Printf("accept error:%v", err)
				continue
			}
			go handleConn(ctx, conn)
		}
	}
	return nil
}

/**
 * @desc 处理连接请求
 */
func handleConn(ctx context.Context, conn net.Conn) {
	defer conn.Close()
	ch := make(chan []byte)
	group, _ := errgroup.WithContext(ctx)

	group.Go(func() error {
		return read(ctx, conn, ch)
	})
	group.Go(func() error {
		return write(ctx, conn, ch)
	})

	group.Wait()
	fmt.Println(conn.RemoteAddr(), "connect has been closed")
}

// 读方法只管循环的读取连接上获取到的消息。然后把消息传递给写操作。
func read(ctx context.Context, conn net.Conn, ch chan []byte) error {
	readCn := bufio.NewReader(conn)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			message, _, err := readCn.ReadLine()
			if err != nil {
				close(ch)
				return err
			}
			fmt.Printf("read conn message %v \n", string(message))
			ch <- message
		}
	}
}

func write(ctx context.Context, conn net.Conn, ch chan []byte) error {
	writeCon := bufio.NewWriter(conn)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			message, ok := <-ch
			if !ok {
				return nil
			}
			if len(message) <= 0 {
				continue
			}
			writeCon.WriteString("hi:")
			writeCon.Write(message)
			writeCon.WriteString("\n")
			writeCon.Flush()
		}
	}
}
