package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type details struct {
	Certname string `json:"certname"`
	Name     string `json:"name,omitempty"`
	Value    string `json:"value,omitempty"`
}

//var results map[string][]string``

func main() {

	//base := "http://localhost:8180/v2/nodes/"
	
	var hosts []string
	hosts = fetchPuppetInfo(base)
	i := 1
	for _, v := range hosts {
		//v = "qb004.wvrgroup.internal"
		url := base + v + "/facts"
		i++

		time.Sleep(10 * time.Millisecond)
		go fetchPuppetInfo(url)
	}
}

func fetchPuppetInfo(url string) []string {
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

	hosts := hostList(body)
	return hosts
}

func hostList(pContents []byte) []string {

	var H []string
	var r []details
	json.Unmarshal(pContents, &r)

	for _, v := range r {
		H = append(H, v.Certname, v.Name, v.Value)
		fmt.Println(v.Certname, v.Name, v.Value)
	}
	return H
}
