package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// TODO see if I can run this in the same pod as the cloudeos
// and still make the curl work

// TODO retry or at least fail out when it doesn't work causing container to reschedule

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
		if strings.HasSuffix(key, ":9910") {
			cvURL := "https://" + strings.TrimSuffix(key, ":9910")
			log.Printf("Cloudvision located at : %s", cvURL)
		}
	}
}
