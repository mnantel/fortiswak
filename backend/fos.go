package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func fosClientRaw(target Targetconfig, method string, path string, data []byte, param []byte) ([]byte, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := &http.Client{
		Timeout: time.Second * 2,
	}
	fmt.Println(target.Apikey, target.Firewallip+path, string(data))
	req, _ := http.NewRequest(method, "https://"+target.Firewallip+path, bytes.NewBuffer(data))
	req.Header.Set("Authorization", "Bearer "+target.Apikey)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	resbody, _ := ioutil.ReadAll(res.Body)

	return resbody, nil
}

// func fosGetSchema() (ApiSchemaResponse, error) {
// 	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
// 	client := &http.Client{}

// 	req, _ := http.NewRequest("GET", "https://172.16.217.3:443/api/v2/cmdb?action=schema", nil)
// 	req.Header.Set("Authorization", "Bearer x")
// 	res, err := client.Do(req)
// 	if err != nil {
// 		panic(err)
// 	}
// 	resbody, _ := ioutil.ReadAll(res.Body)
// 	var schema ApiSchemaResponse
// 	err = json.Unmarshal(resbody, &schema)
// 	return schema, nil
// }

// func getapipathlist() {
// 	seenlist := make(map[string]bool)
// 	var apipathlist []string
// 	for _, v := range schema.Results {
// 		if strings.Contains(v.Path, "__tree__") {
// 			continue
// 		}
// 		if _, exists := seenlist[v.Path]; !exists {
// 			apipathlist = append(apipathlist, v.Path)
// 			seenlist[v.Path] = true
// 		}
// 	}
// 	sort.Strings(apipathlist)
// 	output, _ := json.Marshal(apipathlist)
// 	fmt.Println(string(output))
// }
