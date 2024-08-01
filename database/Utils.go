package database

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Update(db *sqlx.DB, any_value interface{}) (sql.Result, error) {
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
	return db.NamedExec(query, any_value)
}

func Exist(db *sqlx.DB, any_value interface{}) bool {
	var query string = "SELECT * FROM "
	query += reflect.TypeOf(any_value).Name()
	query += " WHERE id = :id"
	rows, err := db.NamedQuery(query, any_value)
	if err != nil {
		fmt.Println("Error checking if item exists: ", err)
		return false
	}
	defer rows.Close()
	return rows.Next()
}

func Save(db *sqlx.DB, any_value interface{}) (sql.Result, error) {
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
	return db.NamedExec(query, any_value)
}