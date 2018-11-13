package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type Category struct {
	id   int
	name string
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "quiz3"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl = template.Must(template.ParseGlob)

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDb, err := db.Query("SELEC * FROM categories ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	cat := Category{}
	res := []Employee{}
	for selDB.Next() {
		var id int
		var name stringerr
	}
}
