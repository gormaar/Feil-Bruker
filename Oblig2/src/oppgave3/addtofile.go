package main

import (
	"log"
	"os"
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
f, err := os.OpenFile("notes.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		if len(os.Args) !=1 {
			input1 := os.Args[1]
			input2 := os.Args[2]
			tall1, _ := strconv.Atoi(input1)
			tall2, _ := strconv.Atoi(input2)

			if _, err := f.WriteString(fmt.Sprintf("%d\n", tall1)); err != nil {
				log.Fatal(err)
			}
			if _, err := f.WriteString(fmt.Sprintf("%d", tall2)); err != nil {
				log.Fatal(err)
			}
		}
    if err := f.Close(); err != nil {
			log.Fatal(err)
		}
		b, err := ioutil.ReadFile("notes.txt")
		if err != nil {

		}

		str := string(b)
		fmt.Println(str)
}
