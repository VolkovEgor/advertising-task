package test

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jmoiron/sqlx"
)

const (
	UsernameTestDB = "postgres"
	PasswordTestDB = "123matan123"
	HostTestDB     = "localhost"
	PortTestDB     = "5432"
	DBnameTestDB   = "postgres_test"
	SslmodeTestDB  = "disable"
	UpTestDBFile   = "migrations/000001_init.up.sql"
	DownTestDBFile = "migrations/000001_init.down.sql"
	TestDataDBFile = "scripts/insert_test_data.sql"
)

func OpenTestDatabase() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		HostTestDB, PortTestDB, UsernameTestDB, DBnameTestDB, PasswordTestDB, SslmodeTestDB))
	return db, err
}

func PrepareTestDatabase(prefix string, insert_data bool) (*sqlx.DB, error) {
	db, err := OpenTestDatabase()
	if err != nil {
		log.Fatal(err)
	}

	down, err := ioutil.ReadFile(prefix + DownTestDBFile)
	if err != nil {
		log.Fatal(err)
	}

	schema, err := ioutil.ReadFile(prefix + UpTestDBFile)
	if err != nil {
		log.Fatal(err)
	}

	db.MustExec(string(down))
	db.MustExec(string(schema))

	if insert_data {
		data, err := ioutil.ReadFile(prefix + TestDataDBFile)
		if err != nil {
			log.Fatal(err)
		}
		db.MustExec(string(data))
	}

	return db, err
}
