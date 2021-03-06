package main

import (
	"fmt"
	"time"
)

type Gorra struct {
	marca string
	color string
	precio float32
	plana bool
}

type Orden struct {
	texto string
	cantidad int
}

func main() {

	var gorra1 = Gorra {
		marca: "nike",
		color: "negra",
		precio: 100,
		plana: true,
	}
	var gorra2 = Gorra {
		"Adidas",
		"negra",
		100,
		true,
	}

	holaMundo(gorra1, gorra2)

	println(tiempoVencer())
	println(pedidos(24))
	recorrerPropiedades("Adidas","negra")
}

func holaMundo(gorra1 Gorra, gorra2 Gorra){

	fmt.Println(gorra1)
	var suma int = 9 + 9 //decalracion de variables
	var nombre string = "Ysrael Mendez"
	fmt.Println("Hola Mundo en GO " + nombre)
	time.Sleep(time.Second * 5) // Permite manejar tiempo
	fmt.Println(suma)

	valorInferido := 14 //reconoce el tipo de variable que sera
	const numero int = 14
	fmt.Println(valorInferido + numero) //opera
	fmt.Println(valorInferido, numero) // concatena

	fmt.Println(gorra2.marca)

}

func tiempoVencer() (string, int) { //Definimos el parametro que retornara
	var dato1 string = "YSRAEL"
	var dato2 int = 14
	return dato1, dato2

}

func pedidos(cantidad int)(string, int)  {
	precio := func() int{
		return cantidad*7
	}
	return "El precio de su pedido es : ", precio()
}

func recorrerPropiedades(caracteristicas ...string){
	for _, caracteristica := range caracteristicas{
		fmt.Println(caracteristica)
	}
}





