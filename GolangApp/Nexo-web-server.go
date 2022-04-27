package main

import (
	"database/sql"
	"fmt"
	"html"
	"log"
	"net/http" // web server wibrary
	"os"

	_ "github.com/lib/pq"
)

const (
	//host     = "db"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "nexo"
)

var (
	greetings string
	username  string
)

func main() {
	host := os.Args[1]
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT greetings, username FROM nexoapp;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&greetings, &username)
		if err != nil {
			panic(err)
		}

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello World, %q", html.EscapeString(r.URL.Path))
		})

		http.HandleFunc("/db", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, greetings)
			fmt.Fprintf(w, " ")
			fmt.Fprintf(w, username)
		})

		log.Fatal(http.ListenAndServe(":8080", nil))
	}

}
