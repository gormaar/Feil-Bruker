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

	conn, err := net.Dial("udp", "localhost:17")
	if err != nil 	{
		log.Fatal(err)
	}
	fmt.Println("Press enter to see the quote of the day")

	reader := bufio.NewReader(os.Stdin)
	stringRead, _ := reader.ReadString('\n')

	fmt.Fprintf(conn, stringRead)

	message, err := bufio.NewReader(conn).ReadString(')')
		if err != nil	{
			log.Fatal(err)
		}
	fmt.Println("Message from server: \n" + message)

	defer conn.Close()
}
