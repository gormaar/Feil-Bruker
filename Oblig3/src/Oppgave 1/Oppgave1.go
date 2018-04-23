package main

import (
	"fmt"
	"net/http"
)

func main()	{
	fmt.Println("Starting server..")
	http.HandleFunc("/", handler)		//Oppretter en path og bruker handler til å skrive melding
	http.ListenAndServe(":8080",nil)	//Port og handler
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello, client")		//Skriver melding
}
