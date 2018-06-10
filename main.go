package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"bytes"
	"encoding/json"
)

func main() {
	response, err := http.Get("https://api.coinbase.com/v2/prices/spot?currency=USD")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Printf(string(data))
	}
	jsonData := map[string]string{"firstname":"Maple", "lastname": "Liu"}
	jsonValue, _ := json.Marshal(jsonData)
	//Requesting with a Header, Can add more customization
	request, _ := http.NewRequest("POST", "https://httpbin.org/post", bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err = client.Do(request)
	//Requesting without a Header
	//response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Printf(string(data))
	}
}