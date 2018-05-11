package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"html/template"

)

func main()	{
	fmt.Println("Starting application")		//Printer melding om at applikasjonen starter
	http.HandleFunc("/", Handler)			//Setter path og bruker Handler
	http.ListenAndServe(":8080", nil)		//Setter port, og handler

}

func errorCheck(err error)	{			//Funksjon for feilhåndtering
	if err != nil	{
		log.Fatal(err)
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {	//Handler som parser og eksekverer html fil
	path := path.Join("template", "index.html")
	tmpl, _ := template.ParseFiles(path)

	tmpl.Execute(w, "")
}
