package main

import (
	"fmt"
	"time"
)
func main(){
	var suma  int = 9 + 9
	var nombre string = "Ysrael Mendez"
	fmt.Println("Hola Mundo en GO " + nombre)
	time.Sleep(time.Second * 5) // Permite manejar tiempo
	fmt.Println(suma)
}
