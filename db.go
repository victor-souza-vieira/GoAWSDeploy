package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func Conectar() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:api-test-pass@id-instancia-rds-teste.c0islweyishx.us-east-1.rds.amazonaws.com:5432/api_test_db")

	if err != nil {
		panic(err.Error())
	}

	return db
}
