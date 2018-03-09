package main

import (
	"log"
	"os"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
	"os/signal"
	"syscall"
)

func main() {

	go sigint2()

	f, err := os.OpenFile("notes.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Print("There was an error opening the file")
		time.Sleep(5*time.Second)
		os.Exit(1)
	}
	if len(os.Args) !=1 {
		input1 := os.Args[1]
		input2 := os.Args[2]
		tall1, err := strconv.Atoi(input1)
		if err != nil {
			fmt.Print("First number was entered incorrect, try again \n")
			time.Sleep(5*time.Second)
			os.Exit(1)
		}
		tall2, err := strconv.Atoi(input2)
		if err != nil {
			fmt.Print("First number was entered incorrect, try again \n")
			time.Sleep(5*time.Second)
			os.Exit(1)
		}

		if _, err := f.WriteString(fmt.Sprintf("%d\n", tall1));
		err != nil {
			fmt.Print("There was an error writing the first number to the file")
		}
		if _, err := f.WriteString(fmt.Sprintf("%d", tall2));
		err != nil {
			fmt.Print("There was an error writing the second number to the file")
		}
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadFile("notes.txt")
	if err != nil {
		fmt.Print("There was an error reading the file")
		time.Sleep(5*time.Second)
		os.Exit(1)
	}

	str := string(b)
	fmt.Println(str)
}
func sigint2(){
	Sigint := make(chan os.Signal, 1)
	signal.Notify(Sigint, syscall.SIGINT)
	<- Sigint
	fmt.Printf("\nA SIGINT made the program stop running \n")
	os.Exit(1)
}
