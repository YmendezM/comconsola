package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)

var listPaliculas = peliculas{
	pelicula{"forrest gump", 1994, "Robert Zemeckis", "Tom Hanks"},
	pelicula{"En busca de la felicidad", 2006, "Gabriele Muccino", "Will Smith"},
	pelicula{"Mas alla de los sueños", 1998, "Vincent Ward", "Robin Williams"},
}

func getSession() *mgo.Session {

	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	return session
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola Mundo server Go")
}

func Contacto(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ysrael Mendez Marciales.")
}

func peliculasList(w http.ResponseWriter, r *http.Request) {

	//fmt.Fprintf(w, "Listado de peliculas")
	json.NewEncoder(w).Encode(listPaliculas)
}

func peliculasShow(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	pelicula_id := params["id"]
	fmt.Fprintf(w, "Cargando pelicula numero %s", pelicula_id)
}

func peliculasAdd(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var peliculasData pelicula
	err := decoder.Decode(&peliculasData)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()
	log.Println(peliculasData)
	listPaliculas = append(listPaliculas, peliculasData)

}
