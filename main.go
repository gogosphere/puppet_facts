package main

// TODO: the hostList and hostFactValue need get an io.Reader from httpclient.
// TODO: add optionparser for --action [--filter]
// TODO: maybe kickout reuseable funtions or make them reusable.

import (
<<<<<<< HEAD
<<<<<<< HEAD
	"bytes"
=======
>>>>>>> parent of a84f20f... dicking around with the puppet puller
=======
>>>>>>> parent of a84f20f... dicking around with the puppet puller
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
<<<<<<< HEAD
<<<<<<< HEAD
	"log"
	"net/http"
=======
>>>>>>> parent of a84f20f... dicking around with the puppet puller
=======
>>>>>>> parent of a84f20f... dicking around with the puppet puller
	"os"
	"strings"
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

func main() {
	cmdbase := strings.Split(os.Args[0], "/")
	fmt.Println(cmdbase[len(cmdbase)-1])
	//hostFactValue()
	//hostList()
<<<<<<< HEAD
<<<<<<< HEAD

	var jsonByte []byte

	url := "http://puppetdb.homeawaycorp.com/nodes"

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
	/*
=======
}

func hostList() {
>>>>>>> parent of a84f20f... dicking around with the puppet puller
=======
}

func hostList() {
>>>>>>> parent of a84f20f... dicking around with the puppet puller
	file, e := ioutil.ReadFile("/Users/whancock/Dropbox/nodesoutput.txt")
	if e != nil {
		fmt.Println("ioutil read bork")
	}
<<<<<<< HEAD
<<<<<<< HEAD
	*/
	var r []Details
	json.Unmarshal(pContents, &r)
=======
	var r []Details
	json.Unmarshal(file, &r)
>>>>>>> parent of a84f20f... dicking around with the puppet puller
=======
	var r []Details
	json.Unmarshal(file, &r)
>>>>>>> parent of a84f20f... dicking around with the puppet puller
	for k, v := range r {
		fmt.Println(k, v.Name)
	}
	fmt.Println(len(r))
}

func hostFactValue() {
	file, e := ioutil.ReadFile("/Users/whancock/Dropbox/factsoutput.txt")
	if e != nil {
		fmt.Println("ioutil read bork")
	}
	var r []Details
	json.Unmarshal(file, &r)
	for k, v := range r {
		fmt.Println(k, v.Certmap.Hostname, v.Name, v.Value)
	}
	fmt.Println(len(r))
}

func (a *Certmap) UnmarshalJSON(b []byte) (err error) {
	j, s, n := certmap{}, "", uint64(0)
	if err = json.Unmarshal(b, &j); err == nil {
		*a = Certmap(j)
		return
	}
	if err = json.Unmarshal(b, &s); err == nil {
		a.Hostname = s
		return
	}
	if err = json.Unmarshal(b, &n); err == nil {
		a.HostID = n
	}
	return
}

func Decode(r io.Reader) (x *Details, err error) {
	x = new(Details)
	err = json.NewDecoder(r).Decode(x)
	return
}
