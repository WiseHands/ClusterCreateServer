package main

import (
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"log"
	"net/http"
)

//CORS enabled
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func responseHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	switch r.Method {
	case "POST":
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("responseHandler request body %s\n", reqBody)

		//Display all request params
		for k, v := range r.URL.Query() {
			log.Printf("responseHandler request param %s: %s\n", k, v)
		}

		y, err := yaml.JSONToYAML(reqBody)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		log.Println(string(y))

		user := "TarasH1"
		token := "myToken"
		url := "https://api.github.com/repos/WiseHands/ClusterDev/contents/config.yaml"

		createOrUpdateClusterConfigFile(user, token, url, string(y))

		fmt.Fprintf(w, string(y))

	}
}
