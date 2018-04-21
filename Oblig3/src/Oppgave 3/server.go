package main

import (
		"fmt"
		"net"
		"os"
		"time"
)


const quote = "Quote of the day"

func main()	{

	fmt.Println("Starting server..")

	go TCP()
	go UDP()

	for    {
		time.Sleep(100)
	}
}

func errorCheck(err error)	{
	if err != nil{
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func TCP()	{
	listen, err := net.Listen("tcp", ":8080")
	errorCheck(err)

	tcpConn, err := listen.Accept()
	errorCheck(err)

	defer tcpConn.Close()

	for {
		tcpConn.Write([]byte(quote))
		}
	}

func UDP()	{
	addr, err := net.ResolveUDPAddr("udp", ":8080")
	errorCheck(err)

	udpConn, err := net.ListenUDP("udp", addr)
	errorCheck(err)

	defer udpConn.Close()

	buf := make([]byte, 1024)

	for {
		_ , addr, err := udpConn.ReadFromUDP(buf)
		udpConn.WriteToUDP([]byte(quote), addr)
		errorCheck(err)
	}
}
