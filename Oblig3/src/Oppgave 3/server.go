package main

import (
		"fmt"
		"net"
		"os"
		"time"
)


const qotd = "“Measuring programming progress by lines \n of code is like measuring aircraft \n building progress by weight.”\n(Bill Gates) "


func main()	{

	fmt.Println("Starting server..")

	go TCP()
	go UDP()

	for    {
		time.Sleep(10)
	}
}

func errorCheck(err error)	{
	if err != nil{
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func TCP()	{
	listen, err := net.Listen("tcp", ":17")
	errorCheck(err)

	tcpConn, err := listen.Accept()
	errorCheck(err)

	for {
		tcpConn.Write([]byte("Quote of the day: \n" + qotd))
		}
	defer tcpConn.Close()
	}

func UDP()	{
	addr, err := net.ResolveUDPAddr("udp", ":17")
	errorCheck(err)

	udpConn, err := net.ListenUDP("udp", addr)
	errorCheck(err)

	buffer := make([]byte, 1024)

	for {
		_ , addr, err := udpConn.ReadFromUDP(buffer)
		udpConn.WriteToUDP([]byte("Quote of the day: \n" + qotd), addr)
		errorCheck(err)
	}
	defer udpConn.Close()
}
