package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type details struct {
	Certname string `json:"certname"`
	Name     string `json:"name,omitempty"`
	Value    string `json:"value,omitempty"`
}

//var results map[string][]string

func main() {

	/* get root command for other dumb shit
	cmdbase := strings.Split(os.Args[0], "/")
	fmt.Println(cmdbase[len(cmdbase)-1])
	*/
	//base := "http://aspdo046.aus1.homeaway.live:44880/v2"
	base := "http://localhost:8080/"

	endpoint := []string{"osfamily", "is_virtual"}
	for _, v := range endpoint {
		//url := base + "/facts/" + v
		url := base + v + ".json"
		fmt.Println(url)
		fetchPuppetInfo(url)
	}
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

	counts := make(map[string]int)
	for _, v := range r {
		counts[v.Value]++
	}
	fmt.Println(counts)

}
