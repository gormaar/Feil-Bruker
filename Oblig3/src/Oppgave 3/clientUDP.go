package main

import (
		"fmt"
		"os"
		"net"
		"bufio"
		"log"
)



func main()	{
	fmt.Println("Starting UDP client")

	conn, err := net.Dial("udp", "localhost:8080")
	if err != nil 	{
		log.Fatal(err)
	}

	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	read, _ := reader.ReadString('\n')

	fmt.Fprintf(conn, read)

	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println("Message from server: ", message)
}
