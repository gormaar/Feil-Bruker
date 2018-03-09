package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"strconv"
	"time"
	"os/signal"
	"syscall"
)

func main() {

	go sigint3()

	//venter 5 sekunder i påvente av eventuell SIGINT

	time.Sleep(5*time.Second)

	// Åpner filen notes.txt med rettigheter for å legge til tekst samt rettigheten writeonly (kan ikke fjerne tekst)
	// Permission code 0644; Bruker kan lese og skrive, group og other kan bare lese.
	file, err := os.OpenFile("notes.txt", os.O_APPEND| os.O_RDWR, 0644)
	// Skriver ut feilmelding om "notes.txt" ikke kan åpnes
	if err != nil {
		fmt.Print("There was an error opening the file")
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(file)
	if err != nil{
		// leser igjennom notes, hvis det er error printer den melding
		fmt.Print("Could not read file")
		os.Exit(1)
	}
	// Konverterer byte[] som tallene er skrevet som i notes.txt til string for å kunne bruke split.
	// Split separerer i dette tilfellet opp i nye strings vel linjeskift (\n) og legger verdiene til input1 og input2
	f:= string(b)
	s := strings.Split(f,"\n")
	input1, input2 := s[0],s[1]
//	fmt.Println(input1,input2)

// Konverterer til int for å kunne addere verdiene
	tall1, _ := strconv.Atoi(input1)
	if err != nil {
		fmt.Print("First number was entered incorrect, try again \n")
		os.Exit(1)
	}

	tall2, err := strconv.Atoi(input2)
	if err != nil {
		fmt.Print("Second number was entered incorrect, try again \n")
		os.Exit(1)
	}

	sum:= tall1+tall2

	fmt.Println(sum)

// Truncate sletter alt innhold i fila når size er 0. Er size 10 vil den slette alt som står etter 10 bytes
	err1 := os.Truncate("notes.txt", 0)
	if err1 != nil {
		fmt.Print("Could not truncate file")
	}
	if _,err :=  file.WriteString(fmt.Sprintf("%d",sum)); err!= nil {
		fmt.Print("Could not write sum")
	}

	defer file.Close()
}

func sigint3(){
	Sigint := make(chan os.Signal, 1)
	signal.Notify(Sigint, syscall.SIGINT)
	<- Sigint
	fmt.Printf("\nA SIGINT made the program stop running \n")
	os.Exit(1)
}
