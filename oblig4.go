package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"html/template"
	"io/ioutil"
	"encoding/json"
)

type Butikker []struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}


var markers Butikker

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

func Handler(w http.ResponseWriter, r *http.Request)	{
	path := path.Join("template", "index.html")
	tmpl, _ := template.ParseFiles(path)

	markers, _ := ioutil.ReadFile("markers.json")
	json.Unmarshal(markers, &markers)
	tmpl.Execute(w, markers)
	fmt.Fprintf(w, "<h1>Butikker nærmest deg:</h1>")



//AIzaSyAvxEV8EQtWbqyCA-NeOfJ94-4QXscKi88
//https://maps.googleapis.com/maps/api/js?key=AIzaSyAvxEV8EQtWbqyCA-NeOfJ94-4QXscKi88&callback=initMap
}
//'https://storage.googleapis.com/maps-devrel/google.json'
//https://maps.googleapis.com/maps/api/distancematrix/json?parameters
//<script src="https://maps.googleapis.com/maps/api/js?key=AIzaSyAvxEV8EQtWbqyCA-NeOfJ94-4QXscKi88&libraries=places&callback=initMap" async defer></script>