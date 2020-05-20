package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
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
		cvURL := ""
		for key, value := range connections {
			log.Printf("key %s value %+v \n", key, value)
			if strings.HasSuffix(key, ":9910") {
				cvURL = "https://" + strings.TrimSuffix(key, ":9910")
				log.Printf("CloudVision located at : %s\n", cvURL)
			}
		}

		podName := os.Getenv("CLOUDEOS_POD_NAME")
		podNamespace := os.Getenv("CLOUDEOS_POD_NAMESPACE")
		pod, err := clientset.CoreV1().Pods(podNamespace).Get(context.TODO(), podName, metav1.GetOptions{})
		if err != nil {
			log.Println(err)
		}
		log.Printf("Annotating pod %s in namespace %s with CloudVision URL: %s\n", podName, podNamespace, cvURL)
		pod.SetAnnotations(map[string]string{
			"CloudVisionURL": cvURL,
		})
		clientset.CoreV1().Pods("").Update(context.TODO(), pod, metav1.UpdateOptions{})
		time.Sleep(300 * time.Second)
	}

}
