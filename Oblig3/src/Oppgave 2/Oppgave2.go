package main

import	(
		"fmt"
		"net/http"
		"io/ioutil"
		"encoding/json"
		"log"
)

type Parkering struct {
	Dato                string `json:"Dato"`
	Klokkeslett         string `json:"Klokkeslett"`
	Sted                string `json:"Sted"`
	Latitude            string `json:"Latitude"`
	Longitude           string `json:"Longitude"`
	AntallLedigePlasser string `json:"Antall_ledige_plasser"`
}
var parkering []Parkering

func main()	{

	fmt.Println("Starting server..")
	http.HandleFunc("/1",  Handler1)
	http.HandleFunc("/2",  Handler2)
	http.HandleFunc("/3",  Handler3)
	http.HandleFunc("/4",  Handler4)
	http.HandleFunc("/5",  Handler5)
	http.ListenAndServe(":8080", nil)

	}

func Handler1(w http.ResponseWriter, r *http.Request) {

	data1, err := http.Get("https://opencom.no/dataset/36ceda99-bbc3-4909-bc52-b05a6d634b3f/resource/d1bdc6eb-9b49-4f24-89c2-ab9f5ce2acce/download/parking.json")
		if err != nil {
			log.Fatal(err)
	}
	bytes1, err := ioutil.ReadAll(data1.Body)
		if err != nil {
			log.Fatal(err)
	}
	json.Unmarshal(bytes1, &parkering)
	
	path := path.Join("templates", "index.html")
	tmpl, _ := template.ParseFiles(path)
	tmpl.Execute(w, parkering)


	for i := 0; i < len(parkering); i++ {
		fmt.Fprintf(w, "%v\n", parkering[i])
	}

	}

func Handler2(w http.ResponseWriter, r *http.Request) {

	data2, err := http.Get()
	if err != nil {
		log.Fatal(err)
	}
	bytes2, err := ioutil.ReadAll(data2.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(bytes2, &parkering)

	for i := 0; i < len(parkering); i++ {
		fmt.Fprintf(w, "%v\n", parkering[i])
	}

}
func Handler3(w http.ResponseWriter, r *http.Request) {

	data3, err := http.Get("https://opencom.no/dataset/36ceda99-bbc3-4909-bc52-b05a6d634b3f/resource/d1bdc6eb-9b49-4f24-89c2-ab9f5ce2acce/download/parking.json")
	if err != nil {
		log.Fatal(err)
	}
	bytes3, err := ioutil.ReadAll(data3.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(bytes3, &parkering)

	for i := 0; i < len(parkering); i++ {
		fmt.Fprintf(w, "%v\n", parkering[i])
	}

}
func Handler4(w http.ResponseWriter, r *http.Request) {

	data4, err := http.Get("https://opencom.no/dataset/36ceda99-bbc3-4909-bc52-b05a6d634b3f/resource/d1bdc6eb-9b49-4f24-89c2-ab9f5ce2acce/download/parking.json")
	if err != nil {
		log.Fatal(err)
	}
	bytes4, err := ioutil.ReadAll(data4.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(bytes4, &parkering)

	for i := 0; i < len(parkering); i++ {
		fmt.Fprintf(w, "%v\n", parkering[i])
	}

}

func Handler5(w http.ResponseWriter, r *http.Request) {

	data5, err := http.Get("https://opencom.no/dataset/36ceda99-bbc3-4909-bc52-b05a6d634b3f/resource/d1bdc6eb-9b49-4f24-89c2-ab9f5ce2acce/download/parking.json")
	if err != nil {
		log.Fatal(err)
	}
	bytes5, err := ioutil.ReadAll(data5.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(bytes5, &parkering)

	for i := 0; i < len(parkering); i++ {
		fmt.Fprintf(w, "%v\n", parkering[i])
	}

}

 
