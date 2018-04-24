package main

import (
		"fmt"
		"net"
		"bufio"
		"log"
		"os"
)



func main()	{
	fmt.Println("Starting UDP client \n")

	conn, err := net.Dial("udp", "localhost:17")	//Oppretter forbindelse med UDP-server på path "localhost:17"
	if err != nil 	{
		log.Fatal(err)
	}
	fmt.Println("Press enter to see the quote of the day")	//Melding til bruker om å trykke enter for at server skal printe QOTD
								

	reader := bufio.NewReader(os.Stdin)			//Oppretter leser
	stringRead, _ := reader.ReadString('\n')		

	fmt.Fprintf(conn, stringRead)				 

	message, err := bufio.NewReader(conn).ReadString(')')	//Leser melding fra serveren frem til ")"
		if err != nil	{
			log.Fatal(err)
		}
	fmt.Println("Message from server: \n" + message)	//Skriver melding fra server 

	defer conn.Close()					//Avslutter forbindelse
}
