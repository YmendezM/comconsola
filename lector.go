package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main()  {
 fmt.Println("Lector")

 nuevoTexto := os.Args[1]+"\n"
 //escribir := ioutil.WriteFile("fichero.txt", nuevoTexto, 0777)
 fichero, err := os.OpenFile("fichero.txt", os.O_RDWR, 0777)
 showError(err)
 escribir, err := fichero.WriteString(nuevoTexto)
 fmt.Println(escribir)
 showError(err)
 fichero.Close()


 texto, errorFichero := ioutil.ReadFile("fichero.txt")
 showError(errorFichero)
 fmt.Println(string(texto))


}

func showError(e error)  {
	if (e != nil) {
		panic(e)
	}
}
