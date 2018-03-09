package main

import(
	"fmt"
	"os"
	"strconv"
	"syscall"
	"os/signal"
	"time"
)


// Func b legger sammen tall1 og tall2 som den mottar i func main. Adderer de sammen og sender til c. 
func B(c chan int, someValue int, someValue2 int){

	c <- someValue + someValue2

}


func main(){

	go sigint()

	//venter 5 sekunder i påvente av eventuell SIGINT

	time.Sleep(5*time.Second)

	// Lager en channel
	fooVal := make(chan int)
	input1:= os.Args[1]
	input2:= os.Args[2]

	// konverterer argumentene til int og skriver ut feilmelding om feks argumentene er bokstaver eller tegn
	tall1,err:= strconv.Atoi(input1)
	if err != nil {
		fmt.Print("First number was entered incorrect, try again \n")
		os.Exit(1)
	}
	tall2,err:= strconv.Atoi(input2)
	if err != nil {
		fmt.Print("Second number was entered incorrect, try again \n")
		os.Exit(1)
	}
	
	// starter func B med parametrene tall1 og tall2 som er inputverdiene. Mottar summen fra func B (fooVal) og printer resultat
	go B(fooVal, tall1, tall2)
	result:= <-fooVal
	fmt.Println(result)
	os.Exit(1)

}

func sigint(){
	Sigint := make(chan os.Signal, 1)
	signal.Notify(Sigint, syscall.SIGINT)
	<- Sigint
	fmt.Printf("\nA SIGINT made the program stop running \n")
	os.Exit(1)
}
