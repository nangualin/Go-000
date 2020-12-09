package main
import(
   "fmt"
   "os"
   "os/signal"
   "context"
   "net/http"
   "golang.org/x/sync/errgroup"
)

type handler struct {
	name string
}

var signals chan os.Signal

func (hand handler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Server %s get Request", hand.name)
}
func main() {
   signals = make(chan os.Signal,1)
   group ,ctx := errgroup.WithContext(context.Background());
   serveOne := &http.Server{Addr: ":54110",Handler:handler{"one"}}
   serveTwo := &http.Server{Addr: ":54120",Handler:handler{"two"}}

   group.Go(
       func() error {
          if err := serveOne.ListenAndServe() ; err != http.ErrServerClosed {
                return err
          }
          return nil
       },
   )
   group.Go(
        func() error {
           if err := serveTwo.ListenAndServe() ; err != http.ErrServerClosed {
               return err
           }
           return nil
        },
   )
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