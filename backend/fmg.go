package main

import (
	"crypto/tls"
	"encoding/json"
	"fmg"
	"fmt"
	"log"
	"net/http"
)

var (
	FMG fmg.FMG
)

func fmgClientRaw(target Targetconfig, method string, path string, data []byte, param []byte) ([]byte, error) {

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	FMG.Username = target.Username
	FMG.Password = target.Apikey
	FMG.IP = target.Firewallip

	err := FMG.Login()
	if err != nil {
		log.Println(err)
		return make([]byte, 0), err
	}
	res, err := FMG.Call(method, path, nil)
	if err != nil {
		fmt.Println(err)
		return make([]byte, 0), err
	}
	output, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
		return make([]byte, 0), err
	}
	return output, nil
}
