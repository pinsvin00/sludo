package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)


func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func log_bytes(bytes []byte){
	sb := string(bytes)
	log.Printf(sb)
}

func load_request(request * http.Request)map[string]string {
	//fmt.Println("New request...")
	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		panic(err)
	}

	//log_bytes(body)
	request_dict := make(map[string]string)
	json.Unmarshal(body, &request_dict)
	return request_dict
}