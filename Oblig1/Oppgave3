package main
import ("fmt"
		"os"
		"os/signal"
		"syscall"
		"time"
)

func main() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall. SIGINT, syscall.SIGKILL, syscall.SIGABRT)

	go func() {
		for sig := range ch {
			switch sig {
			case syscall.SIGTERM:
				fmt.Println("sigterm recieved")
				os.Exit(0)
			case syscall.SIGINT:
				fmt.Println("sigint recieved")
				os.Exit(0)
			case syscall.SIGKILL:
				fmt.Println("sigkill recieved")
				os.Exit(0)
			case syscall.SIGQUIT:
				fmt.Println("sigquit recieved")
				os.Exit(0)
			}
		}
	}()
	time.Sleep(time.Minute)
}
