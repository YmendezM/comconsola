package main

import (
	"fmt"
	"log"
	"net/http"
	//"log"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(writer, "Hola Mundo server Go")
	})

	server := http.ListenAndServe(":8080", nil)
	log.Fatal(server)
}
