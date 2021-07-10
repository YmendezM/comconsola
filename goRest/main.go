package main

import (
	"log"
	"net/http"
)

func main() {
	/*addr := ":8080"
	  http.HandleFunc("/", MyHandler)
	  log.Println("listen on", addr)
	  log.Fatal( http.ListenAndServe(addr, nil) )

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", Index)
	r.HandleFunc("/contacto", Contacto)
	r.HandleFunc("/peliculas", peliculasList)
	r.HandleFunc("/peliculas/{id}", peliculasShow)*/

	r := NewRoute()
	addr := ":8080"
	log.Println("listen on", addr)
	server := http.ListenAndServe(addr, r)
	log.Fatal(server)
}
