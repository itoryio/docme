package storage

import (
	"fmt"
	"os"

	"encoding/json"

	"errors"

	log "github.com/Sirupsen/logrus"
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
	MountID     string
}

//Init method init boltdb
func Init() *Storage {

	//1. инициализация базы данных
	log.Info("Init database")
	db, err := bolt.Open("/tmp/my.db", 0666, nil)
	if err != nil {
		panic(err)
		//log.Fatal(err)
	}
	log.Info("Initialize collection")
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("Volumes"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	return &Storage{db: db}

}

//Get method return recordset with volume data
func (storage *Storage) Get(index string) (Record, error) {
	recordset := Record{}
	err := storage.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Volumes"))
		v := b.Get([]byte(index))
		if len(v) == 0 {
			return fmt.Errorf("Storage error: Cannot get volume: %s", index)
			//return errors.New("Cannot find volume ")
		}
		_ = json.Unmarshal(v, &recordset)
		return nil
	})
	if err != nil {
		log.Errorf(err.Error())
		return recordset, errors.New(err.Error())
	}
	return recordset, nil
}

//Delete method remove specified volume from db
func (storage *Storage) Delete(index string) error {
	err := storage.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Volumes"))
		err := b.Delete([]byte(index))
		return err
	})

	if err != nil {
		log.Errorf(err.Error())
		return errors.New(err.Error())
	}

	return nil
}

//List method
func (storage *Storage) List() map[string]Record {

	recordset := make(map[string]Record)
	storage.db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("Volumes"))
		b.ForEach(func(k, v []byte) error {
			item := Record{}
			_ = json.Unmarshal(v, &item)
			recordset[string(k)] = item
			return nil
		})
		return nil
	})
	return recordset
}

//Update method call Create.
func (storage *Storage) Update(k string, rec Record) {
	storage.Create(rec)
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

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)
}
