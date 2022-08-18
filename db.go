package main

import (
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
)

func Conectar() *sql.DB {
	godotenv.Load(".env")

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	connStr := "postgres://" + user + ":" + pass + "@" + host + ":" + port + "/" + dbName

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err.Error())
	}

	return db
}
