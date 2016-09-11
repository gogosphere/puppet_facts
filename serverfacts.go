package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type details struct {
	Certname string `json:"certname"`
	Name     string `json:"name,omitempty"`
	Value    string `json:"value,omitempty"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Syntax: pfacts infopath \n nodes\n nodes/($servername)/facts\n nodes/($servername)/facts/($fact)\n facts/($fact)\n nodes/($fact)/($value)\n")
		os.Exit(1)
	}

	path := os.Args[1]
	base := "http://aspdo046.aus1.homeaway.live:44880/v2/"
	url := base + path
	fetchPuppetInfo(url)
}

func fetchPuppetInfo(url string) {
	var jsonByte []byte
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonByte))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatal("Die: ", err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Die: ", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	hostList(body)
}

func hostList(pContents []byte) {
	var r []details
	json.Unmarshal(pContents, &r)

	for _, v := range r {
		fmt.Printf("%s: %s : %s\n", v.Certname, v.Name, v.Value)
	}
}
