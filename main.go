package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	//TODO Add http server to provide API to data gathered
	for {
		url := "http://localhost:6060/rest/Eos/TerminAttr/connection"
		resp, err := http.Get(url)
		if err != nil {
			log.Printf("Unable to reach TerminAttr REST endpoint\n")
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
		time.Sleep(300 * time.Second)
	}

}
