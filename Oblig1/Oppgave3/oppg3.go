package main

import ("fmt"
		"os"
		"bufio"
	"syscall"
	"os/signal"
)



func main() {

	for i := 0; i > 0; i++ {

	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
		text := scanner.Text()
		if text == "sigterm" {
			fmt.Println("SIGTERM")
			ch <- syscall.SIGTERM
		} else if text == "sigkill" {
			fmt.Println("SIGKILL")
			ch <- syscall.SIGKILL
		} else if text == "sigquit" {
			fmt.Println("SIGQUIT")
			ch <- syscall.SIGQUIT
		} else if text == "sigint" {
			fmt.Println("SIGINT")
			ch <- syscall.SIGINT
		}

}

