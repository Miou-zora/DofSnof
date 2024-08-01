package database

import (
	"errors"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"log"
	"os"

	"database/sql"
	"fmt"
	"reflect"
	"strings"
)

// type IDB interface {
// 	Setup() error
// 	Close()
// 	Update(any_value interface{}) (sql.Result, error)
// 	Exist(any_value interface{}) bool
// 	Save(any_value interface{}) (sql.Result, error)
// }

type DB struct {
	sqlxDB *sqlx.DB
}

func (db *DB) Setup() error {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")
	if !areEnvVarsValids(user, password, dbname, host) {
		return errors.New("environment variables not set")
	}

	return connectToDb(db, user, dbname, password, host)
}

func (db *DB) Close() {
	db.sqlxDB.Close()
}

func (db *DB) Update(any_value interface{}) (sql.Result, error) {
	var query string = "UPDATE "
	query += reflect.TypeOf(any_value).Name()
	query += " SET "
	for i := 0; i < reflect.TypeOf(any_value).NumField(); i++ {
		query += reflect.TypeOf(any_value).Field(i).Name
		query += " = :"
		query += strings.ToLower(reflect.TypeOf(any_value).Field(i).Name)
		if i != reflect.TypeOf(any_value).NumField()-1 {
			query += ", "
		}
	}
	query += " WHERE id = :id"
	return db.sqlxDB.NamedExec(query, any_value)
}

func (db *DB) Exist(any_value interface{}) bool {
	var query string = "SELECT * FROM "
	query += reflect.TypeOf(any_value).Name()
	query += " WHERE id = :id"
	rows, err := db.sqlxDB.NamedQuery(query, any_value)
	if err != nil {
		fmt.Println("Error checking if item exists: ", err)
		return false
	}
	defer rows.Close()
	return rows.Next()
}

func (db *DB) Save(any_value interface{}) (sql.Result, error) {
	var query string = "INSERT INTO "
	query += reflect.TypeOf(any_value).Name()
	query += " ("
	for i := 0; i < reflect.TypeOf(any_value).NumField(); i++ {
		query += reflect.TypeOf(any_value).Field(i).Name
		if i != reflect.TypeOf(any_value).NumField()-1 {
			query += ", "
		}
	}
	query += ") VALUES ("
	for i := 0; i < reflect.TypeOf(any_value).NumField(); i++ {
		query += ":" + strings.ToLower(reflect.TypeOf(any_value).Field(i).Name)
		if i != reflect.TypeOf(any_value).NumField()-1 {
			query += ", "
		}
	}
	query += ")"
	return db.sqlxDB.NamedExec(query, any_value)
}

func connectToDb(db *DB, user string, dbname string, password string, host string) error {
	var err error
	db.sqlxDB, err = sqlx.Connect("postgres", "user="+user+" dbname="+dbname+" sslmode=disable password="+password+" host="+host)
	if err != nil {
		return err
	}
	if err := db.sqlxDB.Ping(); err != nil {
		log.Println(err)
		return err
	} else {
		log.Println("Successfully Connected to: ", dbname)
	}
	return nil
}

func areEnvVarsValids(user, password, dbname, host string) bool {
	var allSet bool = true
	if user == "" {
		log.Println("POSTGRES_USER not set")
		allSet = false
	}
	if password == "" {
		log.Println("POSTGRES_PASSWORD not set")
		allSet = false
	}
	if dbname == "" {
		log.Println("POSTGRES_DB not set")
		allSet = false
	}
	if host == "" {
		log.Println("POSTGRES_HOST not set")
		allSet = false
	}
	return allSet
}
