package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//GetDataEndpoint : Get data from database for each serialNO
func GetDataEndpoint(w http.ResponseWriter, req *http.Request) {
	//params := mux.Vars(req)
	queryValues := req.URL.Query()
	//weatherDeatils := getData(params["id"], params["size"], params["num"])
	weatherDeatils := getData(queryValues.Get("id"), queryValues.Get("offset"), queryValues.Get("limit"))
	json.NewEncoder(w).Encode(&weatherDeatils)
}

//SaveDataEndpoint : Save client data
func SaveDataEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var weather Weather
	_ = json.NewDecoder(req.Body).Decode(&weather)
	weather.SerialNo = params["id"]
	//weatherDeatils = append(weatherDeatils, weather)
	//json.NewEncoder(w).Encode(&weatherDeatils)
	saveData(weather)
}

func initHttp() {
	router := mux.NewRouter()
	//router.HandleFunc("/weather/{id}/{size}/{num}", GetDataEndpoint).Methods("GET")
	router.HandleFunc("/weather", GetDataEndpoint).Methods("GET")
	router.HandleFunc("/weather/{id}", SaveDataEndpoint).Methods("POST")
	log.Fatal(http.ListenAndServe(":12345", router))
}
