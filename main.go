package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := "http://localhost:6060/rest/Eos/TerminAttr/connection"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	log.Printf("Retreived from TerminAttr %+v\n", resp)
	body, err := ioutil.ReadAll(resp.Body)
	var connections map[string]interface{}
	json.Unmarshal(body, &connections)
	log.Printf("connections %+v\n", connections)
	for key, value := range connections {
		log.Printf("key %s value %+v \n", key, value)
	}
}
