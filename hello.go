package main

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type DENEME struct{
	Test string `json:"init"`
	Arr []DataSeries `json:"dataseries"`

}

type DataSeries struct {
	Timepoint int `json:"timepoint"`
	Prectype string `json:"prec_type"`
	Wind Wind `json:"wind10m"`
}

type Wind struct {
	Direction string `json:"direction"`
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	log.Fatal(http.ListenAndServe(":3482", myRouter))
}

func main() {
	url := "http://www.7timer.info/bin/api.pl?lon=113.17&lat=23.09&product=astro&output=json"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}
	defer resp.Body.Close()
	var record = DENEME{}
	json.NewDecoder(resp.Body).Decode(&record)
	fmt.Println(record.Test);
	fmt.Println(record.Arr);
	fmt.Println("HELLO THERE")
	handleRequest()
}
