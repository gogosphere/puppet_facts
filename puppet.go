package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type details struct {
	Certmap Certmap `json:"certname"`
	Name    string  `json:"name,omitempty"`
	Value   string  `json:"value,omitempty"`
}

type Certmap struct {
	Hostname string `json:"cert"`
	HostID   uint64 `json:"id"`
}

type certmap Certmap

var results map[string][]string

func main() {

	finder := make(map[string][]string)
	var results map[string][]string
	cmdbase := strings.Split(os.Args[0], "/")
	fmt.Println(cmdbase[len(cmdbase)-1])

	base := "http://aspdo046.aus1.homeaway.live:44880/v2"

	endpoint := []string{"osfamily", "is_virtual"}
	for _, v := range endpoint {
		url := base + "/facts/" + v
		fmt.Println("sending to next function: ", url)
		results = fetchPuppetInfo(url, finder)

	}
	for k, v := range results {
		fmt.Println(k, v)
	}

}

func fetchPuppetInfo(url string, finder map[string][]string) map[string][]string {

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

	results := hostList(body, finder)
	return results

}

func hostList(pContents []byte, finder map[string][]string) map[string][]string {

	var r []details
	json.Unmarshal(pContents, &r)

	for _, v := range r {

		if len(finder[v.Certmap.Hostname]) == 0 {
			finder[v.Certmap.Hostname] = []string{v.Name, v.Value}

		} else {
			a := append(finder[v.Certmap.Hostname], v.Name, v.Value)
			finder[v.Certmap.Hostname] = a
		}
	}
	return finder

}
