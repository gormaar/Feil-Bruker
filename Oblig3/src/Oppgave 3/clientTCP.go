package main

import (
	"fmt"
	"net"
	"bufio"
	"log"
)

func main()	{
	fmt.Println("Starting TCP client \n")

	conn, err := net.Dial("tcp", "localhost:17")
	if err != nil	{
		log.Fatal(err)
	}

	message, _ := bufio.NewReader(conn).ReadString(')')

	fmt.Println("Message from server: \n" + message)

	defer conn.Close()

}
