package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	/*addr := ":8080"
	  http.HandleFunc("/", MyHandler)
	  log.Println("listen on", addr)
	  log.Fatal( http.ListenAndServe(addr, nil) )*/

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", Index)
	r.HandleFunc("/contacto", Contacto)
	r.HandleFunc("/peliculas", peliculasList)
	r.HandleFunc("/peliculas/{id}", peliculasShow)

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

func peliculasList(w http.ResponseWriter, r *http.Request) {

	listPaliculas := peliculas{
		pelicula{"forrest gump", 1994, "Robert Zemeckis", "Tom Hanks"},
		pelicula{"En busca de la felicidad", 2006, "Gabriele Muccino", "Will Smith"},
		pelicula{"Mas alla de los sueños", 1998, "Vincent Ward", "Robin Williams"},
	}
	//fmt.Fprintf(w, "Listado de peliculas")
	json.NewEncoder(w).Encode(listPaliculas)
}

func peliculasShow(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	pelicula_id := params["id"]
	fmt.Fprintf(w, "Cargando pelicula numero %s", pelicula_id)
}
