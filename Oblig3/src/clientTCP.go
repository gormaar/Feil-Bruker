package main

import (
	"fmt"
	"net"
	"bufio"

	"log"
)

func main()	{
	fmt.Println("Starting TCP client")

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil	{
		log.Fatal(err)
	}

	defer conn.Close()

	message, _ := bufio.NewReader(conn).ReadString('\n')

	fmt.Print("Message from server: ", message)

}