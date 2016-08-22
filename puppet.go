package main

// TODO: the hostList and hostFactValue need get an io.Reader from httpclient.
// TODO: add optionparser for --action [--filter]
// TODO: maybe kickout reuseable funtions or make them reusable.

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Details struct {
	Certmap Certmap `json:"certname"`
	Name    string  `json:"name,omitempty"`
	Value   string  `json:"value,omitempty"`
}

type Certmap struct {
	Hostname string `json:"cert"`
	HostID   uint64 `json:"id"`
}

type certmap Certmap

var url string
var uri string
var results map[string][]string

func main() {

	qanda := make(map[string][]string)
	finder := make(map[string][]string)
	var results map[string][]string
	//cmdbase := strings.Split(os.Args[0], "/")
	//fmt.Println(cmdbase[len(cmdbase)-1])

	url := "http://aspdo046.aus1.homeaway.live:44880/v2"
	qanda["osfamily"] = []string{"true", "false"}
	qanda["is_virtual"] = []string{"Windows", "RedHat", "Ubuntu", "Debian"}

	endpoint := []string{"osfamily", "is_virtual"}
	for _, v := range endpoint {
		uri := url + "/facts/" + v
		fmt.Println("sending to next function: ", uri)
		results = fetchPuppetInfo(uri, finder)

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

	var r []Details
	json.Unmarshal(pContents, &r)

	for _, v := range r {
		//data := []string{v.Name + ":" + v.Value}

		if len(finder[v.Certmap.Hostname]) == 0 {
			finder[v.Certmap.Hostname] = []string{v.Name, v.Value}

		} else {
			a := append(finder[v.Certmap.Hostname], v.Name, v.Value)
			finder[v.Certmap.Hostname] = a
		}
	}
	return finder

}

func hostFactValue(pContents []byte) {

	var r []Details
	json.Unmarshal(pContents, &r)
	for k, v := range r {
		fmt.Println(k, v.Certmap.Hostname, v.Name, v.Value)
	}
	fmt.Println(len(r))
}

func (a *Certmap) UnmarshalJSON(b []byte) (err error) {
	j, s := certmap{}, ""
	if err = json.Unmarshal(b, &j); err == nil {
		*a = Certmap(j)
		return
	}
	if err = json.Unmarshal(b, &s); err == nil {
		a.Hostname = s
		return
	}

	return
}
