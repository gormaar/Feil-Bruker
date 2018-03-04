package main

import(
	"fmt"
	"os"
	"strconv"
)

func B(c chan int, someValue int, someValue2 int){

	c <- someValue + someValue2

}

func main(){
	fooVal := make(chan int)
	input1:= os.Args[1]
	input2:= os.Args[2]

	tall1,_:= strconv.Atoi(input1)
	tall2,_:= strconv.Atoi(input2)

	go B(fooVal, tall1, tall2)
	result:= <-fooVal
	fmt.Println(result)

}
