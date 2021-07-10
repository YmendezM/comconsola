package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	/*addr := ":8080"
    http.HandleFunc("/", MyHandler)
    log.Println("listen on", addr)
    log.Fatal( http.ListenAndServe(addr, nil) )*/

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", Index)
	r.HandleFunc("/contacto", Contacto)
	r.HandleFunc("/addServes", Index)

	addr := ":8080"
	log.Println("listen on", addr)
	server := http.ListenAndServe(addr, r)
	log.Fatal(server)
}

func Index(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hola Mundo server Go")
}

func Contacto(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ysrael Mendez Marciales.")
}