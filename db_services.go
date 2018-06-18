package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/dgraph-io/badger"
	"github.com/dispatchlabs/commons/utils"
)

var (
	dbServiceInstance *DBService
	dbServiceOnce     sync.Once
)

// DBService holds the connection to the badger DB.
type DBService struct {
	db      *badger.DB
	running bool
}

func newDBService(running bool) *DBService {
	return &DBService{
		running: running,
	}
}

// GetDBService gets an instance of DBService which we can use to get a connection
// to the DB and interact with it.
func GetDBService() *DBService {
	dbServiceOnce.Do(func() {
		dbServiceInstance = newDBService(true)
		err := dbServiceInstance.OpenDB()
		if err != nil {
			utils.Fatal(err)
			os.Exit(1)
			return
		}
	})
	return dbServiceInstance
}

// GetDB returns is an instance of the Database.
func (d *DBService) GetDB() *badger.DB {
	return d.db
}

// IsRunning returns the state of the DB connection.
func (d *DBService) IsRunning() bool {
	return d.running
}

// OpenDB opens a connection to the database.
func (d *DBService) OpenDB() error {
	opts := badger.DefaultOptions
	opts.Dir = "." + string(os.PathSeparator) + "blockchain.db"
	opts.ValueDir = "." + string(os.PathSeparator) + "blockchain.db"

	fmt.Printf("Creating Badger DB to hold our Blockchain \n")

	db, err := badger.Open(opts)
	if err != nil {
		return err
	}

	d.db = db
	return nil
}

func NewTxn(update bool) *badger.Txn {
	return GetDBService().db.NewTransaction(update)
}
