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

// Programmet kan kjøres på 2 forskjellige måter: Kjøres fra terminal med 2 parametre. eks. go run addtofile.go 4 5
// Da vil programmet legge tallene inn på linje 1 og 2 i notes.txt. og man kan deretter addere i sumfromfile.go
// Eller så kan programmet kjøres uten parametre og det vil da kun printe ut det som allerede ligger i notes.txt

func main() {

	go sigint2()

	//venter 5 sekunder i påvente av eventuell SIGINT

	time.Sleep(5*time.Second)

	// Åpner filen notes.txt med rettigheter for å legge til tekst, lage dokumentet om det ikke allerede finnes samt rettigheten writeonly (kan ikke fjerne tekst)
	// Permission code 0644; Bruker kan lese og skrive, group og other kan bare lese.
	f, err := os.OpenFile("notes.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Print("There was an error opening the file")
		os.Exit(1)
	}
	// Sjekker om lengden på argumentet ikke er 1; om det er 1 er det ikke lagt inn noen parametre og da vil programmet kun skrive ut det som ligger i fila
	if len(os.Args) !=1 {
		input1 := os.Args[1]
		input2 := os.Args[2]

		tall1, err := strconv.Atoi(input1)
		if err != nil {
			fmt.Print("First number was entered incorrect, try again \n")
			os.Exit(1)
		}
		tall2, err := strconv.Atoi(input2)
		if err != nil {
			fmt.Print("First number was entered incorrect, try again \n")
			os.Exit(1)
		}

		if _, err := f.WriteString(fmt.Sprintf("%d\n", tall1));
		err != nil {
			fmt.Print("There was an error writing the first number to the file")
			os.Exit(1)
		}
		if _, err := f.WriteString(fmt.Sprintf("%d", tall2));
		err != nil {
			fmt.Print("There was an error writing the second number to the file")
			os.Exit(1)
		}
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)

	}
	// Leser filen på nytt ( med tall1 og tall2 lagt inn, og skriver ut resultatet.
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
