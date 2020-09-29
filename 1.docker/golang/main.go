package main

import (
	"fmt"
	"log"
	"net/http"
        "strconv"

	"github.com/gorilla/mux"
)

func sumHandler(w http.ResponseWriter, r *http.Request) {
        queryValues := r.URL.Query()
        a, _ := strconv.Atoi(queryValues.Get("a"))
        b, _ := strconv.Atoi(queryValues.Get("b"))
	fmt.Fprintf(w, "%v", a+b)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/v1/sum", sumHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}
