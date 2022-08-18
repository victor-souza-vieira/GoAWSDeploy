package main

import (
	"log"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message":"hello intyx!", "new_field":"pera la amigáum!"}`))
}

func main() {
	http.HandleFunc("/", test)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
