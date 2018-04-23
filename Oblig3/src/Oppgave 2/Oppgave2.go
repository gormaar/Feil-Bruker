package main

import	(
		"fmt"
		"net/http"
		"io/ioutil"
		"encoding/json"
		"log"
		"html/template"
		"path"
)


//Structs med felt for JSON data
type Parkering struct {
	Dato                string `json:"Dato"`
	Klokkeslett         string `json:"Klokkeslett"`
	Sted                string `json:"Sted"`
	Latitude            string `json:"Latitude"`
	Longitude           string `json:"Longitude"`
	AntallLedigePlasser string `json:"Antall_ledige_plasser"`
}

type SSB struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Text string `json:"text"`
}

type Lekeplass struct {
	Entries []struct {
		Latitude  string `json:"latitude"`
		Navn      string `json:"navn"`
		ID        string `json:"id"`
		Longitude string `json:"longitude"`
	} `json:"entries"`
	Page  int `json:"page"`
	Pages int `json:"pages"`
	Posts int `json:"posts"`
}

type Buss struct {
	BusStops []struct {
		CompanyCode  int     `json:"CompanyCode"`
		Dataset      int     `json:"Dataset"`
		StopCode     string  `json:"StopCode"`
		FullName     string  `json:"FullName"`
		Name         string  `json:"Name"`
		ShortName    string  `json:"ShortName"`
		Latitude     float64 `json:"Latitude"`
		Longitude    float64 `json:"Longitude"`
		Zone1        int     `json:"Zone1"`
		Zone2        int     `json:"Zone2"`
		TransferTime int     `json:"TransferTime"`
		Transfer     int     `json:"Transfer"`
		BusStopType  int     `json:"BusStopType"`
		BusStopID    int     `json:"BusStopId"`
	} `json:"BusStops"`
	HashCode    string `json:"HashCode"`
	NumBusStops int    `json:"NumBusStops"`
	APIVersion  int    `json:"ApiVersion"`
}

type Miljostasjon struct {
	Entries []struct {
		Latitude    string `json:"latitude"`
		Navn        string `json:"navn"`
		Plast       string `json:"plast"`
		GlassMetall string `json:"glass_metall"`
		TekstilSko  string `json:"tekstil_sko,omitempty"`
		Longitude   string `json:"longitude"`
	} `json:"entries"`
	Page  int `json:"page"`
	Pages int `json:"pages"`
	Posts int `json:"posts"`
}

//Variabler som holder på JSON data
var parkering []Parkering
var ssb []SSB
var lekeplass Lekeplass
var buss Buss
var miljostasjon Miljostasjon


func main()	{

	fmt.Println("Starting server..")

	//Handlefunc som definerer stier og handlers for hvert datasett
	http.HandleFunc("/1",  Handler1)

	http.HandleFunc("/2",  Handler2)

	http.HandleFunc("/3",  Handler3)

	http.HandleFunc("/4", Handler4)

	http.HandleFunc("/5", Handler5)

	http.ListenAndServe(":8080", nil)	//Definerer adresse og handler

}

func Handler1(w http.ResponseWriter, r *http.Request) {
	//Henter JSON datasett fra URL
	data1, err := http.Get("https://opencom.no/dataset/36ceda99-bbc3-4909-bc52-b05a6d634b3f/resource/d1bdc6eb-9b49-4f24-89c2-ab9f5ce2acce/download/parking.json")
	if err != nil {
		log.Fatal(err)
	}
	bytes1, err := ioutil.ReadAll(data1.Body)	//Gjør dataen om til bytes
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(bytes1, &parkering)		//Dekoder dataen som ligger i bytes, og legger den i variabelen; parkering
	
	path := path.Join("template", "parkering.html")	//Samler pathen til html-fil i en variabel
	tmpl, _ := template.ParseFiles(path)		//Parser pathen
	tmpl.Execute(w, parkering)			//Skriver datasett med html tags til serveren
	
}

func Handler2(w http.ResponseWriter, r *http.Request) {

	data2, err := http.Get("http://data.ssb.no/api/v0/no/table/")
	if err != nil {
	log.Fatal(err)
	}
	bytes2, err := ioutil.ReadAll(data2.Body)
	if err != nil {
	log.Fatal(err)
	}
	json.Unmarshal(bytes2, &ssb)

	path := path.Join("template", "lønn.html")
	tmpl, _ := template.ParseFiles(path)
	tmpl.Execute(w, ssb)

	}

	func Handler3(w http.ResponseWriter, r *http.Request)	{
		data3, err := http.Get("https://hotell.difi.no/api/json/bergen/lekeplasser?")
		if err != nil {
			log.Fatal(err)
		}
		bytes3, err := ioutil.ReadAll(data3.Body)
		if err != nil {
			log.Fatal(err)
		}
		json.Unmarshal(bytes3, &lekeplass)

		path := path.Join("template", "lekeplass.html")
		tmpl, _ := template.ParseFiles(path)
		tmpl.Execute(w, lekeplass)
	}

	func Handler4(w http.ResponseWriter, r *http.Request)	{
		data4, err := http.Get("http://sanntidsappservice-web-prod.azurewebsites.net/busstops?format=json")
		if err != nil	{
			log.Fatal(err)
		}
		bytes4, err := ioutil.ReadAll(data4.Body)
		if err != nil 	{
			log.Fatal(err)
		}
		json.Unmarshal(bytes4, &buss)

		path := path.Join("template", "buss.html")
		tmpl, _ := template.ParseFiles(path)
		tmpl.Execute(w, buss)
	}

func Handler5(w http.ResponseWriter, r *http.Request)	{
		data5, err := http.Get("https://hotell.difi.no/api/json/stavanger/miljostasjoner")
		if err != nil	{
			log.Fatal(err)
		}
		bytes5, err := ioutil.ReadAll(data5.Body)
		if err != nil 	{
			log.Fatal(err)
		}
		json.Unmarshal(bytes5, &miljostasjon)

		path := path.Join("template", "miljøstasjon.html")
		tmpl, _ := template.ParseFiles(path)
		tmpl.Execute(w, miljostasjon)
	}
