package main

import (
   "context"
   "os"
   "log"
   "net/http"
   "os"
   "os/signal"
   "golang.org/x/sync/errgroup"
   "dao"
)

func main() {
	group ,ctx := errorgroup.WithContext(context.Background())
	r := InitHttpHandler(dao.userService{},ctx)
    httpServer := func() error {
       fmt.Println("http server start")
       return http.ListenAndServe(":54119",r)
    }
    // 函数变量传入
	group.GO(httpServer)
	// 直接传入函数。
    group.Go (
       func() error {
            signal.Notify(signals)
            select{
               case sig := <-signals :
                       if err := serveOne.Shutdown(context.Background()) ; err !=nil {
                          fmt.Println("ops,has some errors",err)
                       }
                      if err := serveTwo.Shutdown(context.Background()) ; err !=nil {
                         fmt.Println("ops,has some errors",err)
                      }
                      fmt.Printf("Recevied terminal sig %s\n", sig.String())
                      return fmt.Errorf("Received signal %s",sig.String())
               case <-ctx.Done() :
                     return context.Canceled
            }
       },
    )

   if err := group.Wait() ;err != nil {
       fmt.Println("something has error",err);
   }
}