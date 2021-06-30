package main

import (
	"fmt"
	"time"
)
func main(){
	var suma  int = 9 + 9
	fmt.Println("Hola Mundo en GO");
	time.Sleep(time.Second * 5) // Permite manejar tiempo
	fmt.Println(suma)
}
