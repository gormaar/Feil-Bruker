package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"html/template"

)

func main()	{
	fmt.Println("Starting application")
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)

}

func errorCheck(err error)	{
	if err != nil	{
		log.Fatal(err)
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	path := path.Join("template", "index.html")
	tmpl, _ := template.ParseFiles(path)

	tmpl.Execute(w, "")
}
