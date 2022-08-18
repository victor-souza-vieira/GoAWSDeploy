package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
)

func Conectar() *sql.DB {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	db, err := sql.Open("postgres", "postgres://"+user+":"+pass+"@"+host+":"+port+"/"+dbName)

	if err != nil {
		panic(err.Error())
	}

	return db
}
