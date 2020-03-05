package main

import (
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"log"
	"net/http"
)

func responseHandler(w http.ResponseWriter, r *http.Request) {
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
		fmt.Fprintf(w, string(y))

	}
}
