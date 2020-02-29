package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/////////////////////////////////
// API-facing functions follow //
/////////////////////////////////

func setcors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,POST,PUT,OPTIONS,DELETE")
}

func apiTargetApiCall(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	target, err := dbFindTarget(vars["target"])
	if err != nil {
		log.Println(err)
		return
	}

	reqbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(string(reqbody))

	var apiCall TargetApiCall
	json.Unmarshal(reqbody, &apiCall)
	var rawout []byte
	if target.Devtype == "FortiGate" {
		rawout, err = fosClientRaw(target, apiCall.Method, apiCall.Path, apiCall.Data, nil)
	}

	if target.Devtype == "FortiManager" {
		rawout, err = fmgClientRaw(target, apiCall.Method, apiCall.Path, apiCall.Data, nil)
	}
	fmt.Println(string(rawout))
	setcors(w)
	fmt.Fprint(w, bytes.NewBuffer(rawout))

}

func apiAddTarget(w http.ResponseWriter, r *http.Request) {

	reqbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	var newtarget Targetconfig
	json.Unmarshal(reqbody, &newtarget)
	// VALIDATION HERE
	dbAddTarget(newtarget)
	setcors(w)

}

func apiUpdateTarget(w http.ResponseWriter, r *http.Request) {

	reqbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	var newtarget Targetconfig
	json.Unmarshal(reqbody, &newtarget)
	// VALIDATION HERE
	dbUpdateTarget(newtarget)
	setcors(w)
}

func apiDeleteTarget(w http.ResponseWriter, r *http.Request) {

	reqbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	var newtarget Targetconfig
	json.Unmarshal(reqbody, &newtarget)
	// VALIDATION HERE
	log.Println("Deleting", newtarget.Id)
	dbDeleteTarget(newtarget.Id)
	setcors(w)
}

func apiGetTargets(w http.ResponseWriter, r *http.Request) {

	targets, err := dbGetTargets()
	output, err := json.Marshal(targets)
	if err != nil {
		log.Println(err)
		return
	}

	setcors(w)
	w.Write(output)

}

// Deliver parsed config to frontend as config.json
// func apiGetConfigJS(w http.ResponseWriter, r *http.Request) {

// 	marshalConf, err := json.Marshal(conf.FE)
// 	if err != nil {
// 		panic(err)
// 	}
// 	jsconf := fmt.Sprintf(`
// 	window.CONFIG = %s

// 	`, marshalConf)
// 	w.Header().Add("Content-Type", "application/json")
// 	fmt.Fprint(w, bytes.NewBuffer([]byte(jsconf)))

// }

// func apiGetTargetList(w http.ResponseWriter, r *http.Request) {

// 	type target struct {
// 		Name string `json:"name"`
// 		Ip   string `json:"ip"`
// 	}
// 	var t []target
// 	w.Header().Add("Content-Type", "application/json")
// 	for _, v := range conf.BE.Targets {
// 		t = append(t, target{Name: v.Name, Ip: v.Firewallip})
// 	}
// 	output, _ := json.Marshal(t)
// 	setcors(w)
// 	fmt.Fprint(w, bytes.NewBuffer([]byte(output)))

// }

// func apiGetFos(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	target, _ := findTarget(ps.ByName("target"))
// 	path := ps.ByName("path")
// 	resp, err := fosClientRaw(target, "GET", path, nil, nil)
// 	if err != nil {
// 		return
// 	}
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
// 	fmt.Fprint(w, bytes.NewBuffer(resp))
// }

// func apiGetApiPathlist(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
// 	fmt.Fprint(w, string(output))
// }

// func apiGetApiEndpointlist(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	path := ps.ByName("path")
// 	var endpointlist []string

// 	for _, v := range schema.Results {
// 		if v.Path == path {
// 			endpointlist = append(endpointlist, v.Name)
// 		}
// 	}
// 	sort.Strings(endpointlist)
// 	output, _ := json.Marshal(endpointlist)
// 	fmt.Fprint(w, string(output))
// }

// func apiGetDetectedDevices(w http.ResponseWriter, r *http.Request) {

// 	devs, err := fosGetDetectedDevices()
// 	if err != nil {
// 		return
// 	}

// 	output, _ := json.Marshal(devs)
// 	setcors(w)
// 	fmt.Fprint(w, string(output))

// }
