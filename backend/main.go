package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	bolt "go.etcd.io/bbolt"
)

const (
	RUNINTERVAL = 30 // Default run interval in seconds for pollers
)

var db *bolt.DB

func main() {

	dbpt, err := setupDB()
	if err != nil {
		panic(err)
	}
	db = dbpt
	defer db.Close()

	router := mux.NewRouter()

	router.Use(loggingMiddleware)

	router.Methods("OPTIONS").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			setcors(w)
		})
	router.HandleFunc("/targets/{target}/api", apiTargetApiCall).Methods("POST")
	router.HandleFunc("/targets", apiGetTargets).Methods("GET")
	router.HandleFunc("/targets", apiAddTarget).Methods("POST")
	router.HandleFunc("/targets", apiDeleteTarget).Methods("DELETE")
	router.HandleFunc("/targets", apiUpdateTarget).Methods("PUT")

	// API: Monitor
	// router.HandleFunc("/monitor/user/detected-devices", apiGetDetectedDevices).Methods("GET")

	// Start Pollers
	go runPollers()

	// Start web server
	log.Printf("PBRPortal: Starting server, hostname: %s port:%s\r\n", "localhost", "4000")
	log.Fatal(http.ListenAndServe(":"+"4000", router))

}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func runPollers() {
	for {
		log.Println("POLLING START")
		targets, err := dbGetTargets()
		if err != nil {
			panic("Big bad problem with the DB.")
		}

		for _, t := range targets {
			if t.Devtype == "FortiGate" {
				log.Println("Polling FG -> ", t.Name, "(", t.Firewallip, ")")
				go runPollersFG(t)
			}
			if t.Devtype == "FortiManager" {
				log.Println("Polling FMG -> ", t.Name, "(", t.Firewallip, ")")
				go runPollersFMG(t)
			}
		}
		<-time.After(time.Second * RUNINTERVAL)
	}

}

type DBCacheEntry struct {
	Target Targetconfig             `json:"target"`
	Path   string                   `json:"path"`
	Data   FOSCMDBSystemVdomResults `json:"data"`
}

func runPollersFG(t Targetconfig) {
	// pollFG(t, "/api/v2/cmdb/system/vdom", FOSCMDBSystemVdomResults{})

}

func runPollersFMG(t Targetconfig) {

}
