package storage

import (
	"fmt"

	"encoding/json"

	"github.com/Sirupsen/logrus"
	"github.com/boltdb/bolt"
)

//Storage is main storage struct
type Storage struct {
	db *bolt.DB
}

//Record define a data stucture
type Record struct {
	Name        string
	Mountpoint  string
	Status      map[string]interface{}
	Options     map[string]string
	Connections int
}

//Init method init boltdb
func Init() *Storage {

	//1. инициализация базы данных
	logrus.Info("Init database")
	db, err := bolt.Open("/tmp/my.db", 0666, nil)
	if err != nil {
		panic(err)
		//log.Fatal(err)
	}
	logrus.Info("Initialize collection")
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("Volumes"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	return &Storage{db: db}

}

//ListVolumes return the slice of volumes
func (storage *Storage) ListVolumes() {

}

//Update method
func (storage *Storage) Update(k string, rec Record) {
	storage.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Volumes"))
		data, _ := json.Marshal(rec)
		err := b.Put([]byte(rec.Name), data)
		return err
	})
}

//Create put into a database data about volume
func (storage *Storage) Create(rec Record) {

	storage.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Volumes"))
		data, _ := json.Marshal(rec)
		err := b.Put([]byte(rec.Name), data)
		return err
	})

}
