package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	signals := make(chan os.Signal, 1)
	group, ctx := errgroup.WithContext(context.Background())
	var msg string = "hello word"
	r := InitHttpHandler(msg, ctx)
	s := http.Server{
		Addr:    ":54119",
		Handler: r,
	}
	httpServer := func() error {
		fmt.Println("http server start")
		return s.ListenAndServe()
	}
	// 函数变量传入
	group.Go(httpServer)
	// 直接传入函数。
	group.Go(
		func() error {
			signal.Notify(signals, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
			select {
			case sig := <-signals:
				if err := s.Shutdown(context.Background()); err != nil {
					fmt.Println("ops,has some errors", err)
				}
				fmt.Printf("Recevied terminal sig %s\n", sig.String())
				return fmt.Errorf("Received signal %s ", sig.String())
			case <-ctx.Done():
				return context.Canceled
			}
		},
	)

	if err := group.Wait(); err != nil {
		fmt.Println("something has error", err)
	}
}
