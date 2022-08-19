package main

import (
	"GoAWSDeploy/database"
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	Id   int
	Nome string
}

func test(w http.ResponseWriter, r *http.Request) {
	db := database.Conectar()
	defer db.Close()

	users, err := db.Query("select * from test order by id desc")

	if err != nil {
		panic(err.Error())
	}

	user := User{}
	var lUsers []User

	for users.Next() {
		var id int
		var nome string

		err = users.Scan(&id, &nome)

		if err != nil {
			panic(err.Error())
		}

		user.Id = id
		user.Nome = nome

		lUsers = append(lUsers, user)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lUsers)
}

func main() {
	http.HandleFunc("/", test)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
