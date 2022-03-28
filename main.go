package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

func restConf() {
	var content []byte
	trgt := "1.1.1.1"
	url := "https://" + trgt + "/restconf/data/ietf-interfaces:interfaces"
	ignoreCert := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: ignoreCert}
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(content))
	if err != nil {
		fmt.Println(err)
	}

	req.SetBasicAuth("restconf-user", "pass123")
	req.Header.Add("Content-Type", "application/yang-data+json")
	req.Header.Add("Accept", "application/yang-data+json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	fmt.Println(string(body))
}

func main() {
	restConf()
}
