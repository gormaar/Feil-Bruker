package main

import (
	"fmt"
	"net"
	"bufio"
	"log"
)

func main()	{
	fmt.Println("Starting TCP client \n")

	conn, err := net.Dial("tcp", "localhost:17")		//Setter opp forbindelse med TCP-server på path "localhost:17"
	if err != nil	{
		log.Fatal(err)
	}

	message, _ := bufio.NewReader(conn).ReadString(')')	//Leser melding fra server frem til ")"

	fmt.Println("Message from server: \n" + message)	//Skriver melding fra server

	defer conn.Close()					//Avslutter forbindelse

}
