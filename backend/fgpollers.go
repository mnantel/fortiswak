package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func pollFGState(t Targetconfig) {
	res, err := fosClientRaw(t, "GET", "/api/v2/cmdb/system/vdom", nil, nil)
	if err != nil {

	}
	fmt.Println(res)

}

func pollFG(t Targetconfig, path string, datatype interface{}) {

	// Query FOS API for target PATH
	res, err := fosClientRaw(t, "GET", path, nil, nil)
	if err != nil {
		log.Println("runPollersFG() FOS request error")
		return
	}

	// Unmarshal results
	// var vdomres datatype{}
	err = json.Unmarshal(res, &datatype)
	if err != nil {
		log.Println("runPollersFG() unmarshal error")
		return
	}

	// Update cached data in Mongo
	qFilter := bson.M{"target": bson.M{"$eq": t}, "path": bson.M{"$eq": path}}
	qUpdate := bson.M{"$set": bson.M{"target": t, "path": path, "data": datatype}}

	insertResult, err := dbcache.UpdateOne(context.TODO(), qFilter, qUpdate, options.Update().SetUpsert(true))
	if err != nil {
		log.Fatal(err)
	}

	switch insertResult.UpsertedID {
	case nil:
		log.Println("FG Cache Updated -> ", t.Name, " --> ", path)
	default:
		log.Println("FG Cache Added -> ", t.Name, " --> ", path)
	}

	// // Query cached data and display
	// var entry DBCacheEntry
	// err = dbcache.FindOne(context.TODO(), qFilter).Decode(&entry)

	// fmt.Println("Query output: ", entry.Data.Results)
}
