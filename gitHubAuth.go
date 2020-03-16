package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	user := "TarasH1"
	token := "5dc1f34ac0306dbd0224f8ce12ad07d885222935"
	url := "https://api.github.com/repos/WiseHands/ClusterDev/contents/config.yaml"

	checkIsFileCreated(user, token, url)
}

func checkIsFileCreated(user string, token string, url string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "token "+token)

	client := &http.Client{}
	res, err := client.Do(req)

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		createFile(user, token, url)
	}

	fmt.Printf("Body: %s\n", body)

	type FileType struct {
		Sha  string `json:"sha"`
		Type string `json:"type"`
	}

	//Parse JSON
	jsonData := body
	var dat map[string]interface{}

	if err := json.Unmarshal([]byte(jsonData), &dat); err != nil {
		panic(err)
	}

	shaParam, ok := dat["sha"].(string)
	log.Println("sha", shaParam, ok)

	fileTypeParam, ok := dat["type"].(string)
	log.Println("type", fileTypeParam, ok)

	response := FileType{shaParam, fileTypeParam}
	js, err := json.Marshal(response)
	if err != nil {
		log.Println("JSON response from server error: " + err.Error())

	}
	log.Println("JSON raw response from server: " + string(js))

	updateFile(user, token, url, shaParam)
}

func createFile(user string, token string, url string) {
	data := "some data"
	encoded := base64.StdEncoding.EncodeToString([]byte(data))

	type Committer struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	type RequestBody struct {
		Message   string    `json:"message"`
		Committer Committer `json:"committer"`
		Content   string    `json:"content"`
	}

	committer := Committer{
		Name:  "Bohdan Tsap",
		Email: "bohdaq@gmail.com",
	}

	body := RequestBody{
		Message:   "commit",
		Committer: committer,
		Content:   encoded,
	}

	requestByte, _ := json.Marshal(body)
	requestReader := bytes.NewBuffer(requestByte)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, url, requestReader)
	req.Header.Set("Authorization", "token "+token)
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}

func updateFile(user string, token string, url string, sha string) {
	data := "somekljkl data"
	encoded := base64.StdEncoding.EncodeToString([]byte(data))

	type Committer struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	type RequestBody struct {
		Message   string    `json:"message"`
		Committer Committer `json:"committer"`
		Content   string    `json:"content"`
		Sha       string    `json:"sha"`
	}

	committer := Committer{
		Name:  "TarasH1",
		Email: "research010@gmail.com",
	}

	body := RequestBody{
		Message:   "commit",
		Committer: committer,
		Content:   encoded,
		Sha:       sha,
	}

	requestByte, _ := json.Marshal(body)
	requestReader := bytes.NewBuffer(requestByte)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, url, requestReader)
	req.Header.Set("Authorization", "token "+token)
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
