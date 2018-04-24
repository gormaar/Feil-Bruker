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
	
	go TCP()	//Kjører serverene synkront
	go UDP()

	for    {	//Evig loop
		time.Sleep(10)
	}
}

func errorCheck(err error)	{	//Errorhandler
	if err != nil{
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func TCP()	{
	listen, err := net.Listen("tcp", ":17")	//Deklarerer TCP-serveren med port 17 for Quote of the Day
	errorCheck(err)

	tcpConn, err := listen.Accept()	//Setter opp forbindelse ved Dial fra client
	errorCheck(err)

	for {
		tcpConn.Write([]byte("Quote of the day: \n" + qotd))	//Skriver Quote of the Day og avslutter forbindelsen
		}
	defer tcpConn.Close()
	}

func UDP()	{
	addr, err := net.ResolveUDPAddr("udp", ":17")	//Deklarerer UDP-serveren med port 17 for Quote of the Day
	errorCheck(err)

	udpConn, err := net.ListenUDP("udp", addr)	//Lytter etter Dials fra UDP clients
	errorCheck(err)

	buffer := make([]byte, 1024)			//Oppretter en buffer

	for {
		_ , addr, err := udpConn.ReadFromUDP(buffer)				//Leser fra client
		udpConn.WriteToUDP([]byte("Quote of the day: \n" + qotd), addr)		// Skriver Quote of the Day
		errorCheck(err)
	}
	defer udpConn.Close()								//Avslutter forbindelsen
}
