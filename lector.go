package main

import (
	"fmt"
	"io/ioutil"
)

func main()  {
 fmt.Println("Lector")
 texto, errorFichero := ioutil.ReadFile("fichero.txt")
 showError(errorFichero)
 fmt.Println(string(texto))

}

func showError(e error)  {
	if (e != nil) {
		panic(e)
	}
}
