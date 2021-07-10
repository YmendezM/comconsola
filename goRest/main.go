package main

import (
	"fmt"
	//"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	addr := ":8080"
    http.HandleFunc("/index", MyHandler)
    log.Println("listen on", addr)
    log.Fatal( http.ListenAndServe(addr, nil) )

	/*r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", Index)

	server := http.ListenAndServe("localhost:8080", r)
	log.Fatal(server)*/
}

func MyHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hola Mundo server Go")
}
