package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"

	bolt "go.etcd.io/bbolt"
)

func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func setupDB() (*bolt.DB, error) {
	db, err := bolt.Open("fortiswak.db", 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("could not open db, %v", err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		root, err := tx.CreateBucketIfNotExists([]byte("DB"))
		if err != nil {
			return fmt.Errorf("could not create root bucket: %v", err)
		}

		_, err = root.CreateBucketIfNotExists([]byte("TARGETS"))
		if err != nil {
			return fmt.Errorf("could not create weight bucket: %v", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("could not set up buckets, %v", err)
	}
	fmt.Println("DB Setup Done")
	return db, nil
}

// DB functions
func dbAddTarget(t Targetconfig) error {

	log.Println("Add target: ", t, "...")
	entrybytes, err := json.Marshal(t)
	if err != nil {
		log.Printf("ERROR: could not marshal entry json (%v)", err)
		return fmt.Errorf("E-MARSHAL")
	}
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("DB")).Bucket([]byte("TARGETS"))
		id, _ := b.NextSequence()
		err := b.Put([]byte(strconv.FormatUint(id, 10)), entrybytes)
		if err != nil {
			log.Printf("ERROR: could not add entry: %v", err)
			return fmt.Errorf("E-ADD")
		}

		return nil
	})
	log.Print("OK")
	return err
}

func dbUpdateTarget(t Targetconfig) error {

	log.Println("Update target: ", t, "...")
	entrybytes, err := json.Marshal(t)
	if err != nil {
		log.Printf("ERROR: could not marshal entry json (%v)", err)
		return fmt.Errorf("E-MARSHAL")
	}
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("DB")).Bucket([]byte("TARGETS"))
		err = b.Put([]byte(t.Id), entrybytes)
		if err != nil {
			log.Printf("ERROR: could not update entry: %v", err)
			return fmt.Errorf("E-UPDATE")
		}

		return nil
	})
	log.Print("OK")
	return err
}

func dbGetTargets() ([]Targetconfig, error) {

	var targets []Targetconfig
	err := db.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte("DB")).Bucket([]byte("TARGETS"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var t Targetconfig
			json.Unmarshal(v, &t)
			t.Id = string(k)
			targets = append(targets, t)
		}
		return nil
	})
	if err != nil {
		log.Println(err)
	}

	return targets, nil
}

func dbFindTarget(name string) (Targetconfig, error) {
	targets, _ := dbGetTargets()
	for _, t := range targets {
		if t.Name == name {
			return t, nil
		}
	}
	return Targetconfig{}, errors.New("could no locate target")
}

func dbDeleteTarget(id string) error {

	log.Println("Deleting target: ", id, "...")

	err := db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte("DB")).Bucket([]byte("TARGETS")).Delete([]byte(id))
		if err != nil {
			log.Printf("ERROR: could not delete entry (%v)", err)
			return fmt.Errorf("E-DELETE: %v", err)
		}

		return nil
	})
	log.Print("OK")
	return err
}

func dbAddObj(bucket string, obj struct{}) error {

	log.Println("Add DB entry: ", bucket, "...")
	entrybytes, err := json.Marshal(obj)
	if err != nil {
		log.Printf("ERROR: could not marshal entry json (%v)", err)
		return fmt.Errorf("E-MARSHAL")
	}
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("DB")).Bucket([]byte(bucket))
		id, _ := b.NextSequence()
		err := b.Put([]byte(strconv.FormatUint(id, 10)), entrybytes)
		if err != nil {
			log.Printf("ERROR: could not add entry: %v", err)
			return fmt.Errorf("E-ADD")
		}

		return nil
	})
	log.Print("OK")
	return err
}

func dbUpdateObj(bucket string, obj struct{}, id string) error {

	log.Println("Update target: ", bucket, "...")
	entrybytes, err := json.Marshal(obj)
	if err != nil {
		log.Printf("ERROR: could not marshal entry json (%v)", err)
		return fmt.Errorf("E-MARSHAL")
	}
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("DB")).Bucket([]byte(bucket))
		err = b.Put([]byte(id), entrybytes)
		if err != nil {
			log.Printf("ERROR: could not update entry: %v", err)
			return fmt.Errorf("E-UPDATE")
		}

		return nil
	})
	log.Print("OK")
	return err
}

func dbGetObjAll(bucket string) ([]DbObj, error) {

	var objects []DbObj
	err := db.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte("DB")).Bucket([]byte(bucket))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			o := DbObj{
				Id:   k,
				Data: v,
			}
			objects = append(objects, o)
		}
		return nil
	})
	if err != nil {
		log.Println(err)
	}

	return objects, nil
}

func dbGetObjOne(bucket string, id []byte) ([]byte, error) {

	var obj []byte
	err := db.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte("DB")).Bucket([]byte(bucket))
		obj = b.Get(id)
		return nil
	})
	if err != nil {
		log.Println(err)
	}

	return obj, nil
}

func dbDelObj(bucket string, id []byte) error {

	log.Println("Deleting target: ", id, "...")

	err := db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte("DB")).Bucket([]byte(bucket)).Delete(id)
		if err != nil {
			log.Printf("ERROR: could not delete entry (%v)", err)
			return fmt.Errorf("E-DELETE: %v", err)
		}

		return nil
	})
	log.Print("OK")
	return err
}
