package main

type pelicula struct {
	Name         string `json:"name"`
	Year         int    `json:"year"`
	Director     string `json:"director"`
	Protagonista string `json:"protagonista"`
}

type peliculas []pelicula
