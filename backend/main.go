package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gorilla/mux"
	bolt "go.etcd.io/bbolt"
)

const (
	RUNINTERVAL = 30 // Default run interval in seconds for pollers
)

var conf Config
var dbconf *mongo.Collection
var dbcache *mongo.Collection
var ctx context.Context

// var mongo *mgo.Session

var db *bolt.DB

func main() {

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	dbclient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	dbconf = dbclient.Database("fortiswak").Collection("conf")
	dbcache = dbclient.Database("fortiswak").Collection("cache")
	defer dbclient.Disconnect(ctx)

	dbpt, err := setupDB()
	if err != nil {
		panic(err)
	}
	db = dbpt
	defer db.Close()

	log.Println("PBRPortal: Loading config.json")
	jsonFile, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	jsonBytes, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(jsonBytes, &conf)
	if err != nil {
		panic(err)
	}
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
	log.Printf("PBRPortal: Starting server, hostname: %s port:%s\r\n", conf.FE.Hostname, conf.FE.Port)
	log.Fatal(http.ListenAndServe(":"+conf.FE.Port, router))

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
