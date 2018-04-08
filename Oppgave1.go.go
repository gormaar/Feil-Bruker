package main

import (
	"fmt"
	"net/http"
)

func main()	{
	fmt.Println("Starting server..")
	http.HandleFunc("/", handler)		//Path og handler-funksjon som skriver melding
	http.ListenAndServe(":8080",nil)	//port og handler
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello, client")		//Skriver melding
}
