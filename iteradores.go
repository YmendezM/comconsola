package main

import (
	"fmt"
	"time"
)

func main()  {

	array := []string{"n", "n", "n", "n"}
	/*for i := 0; i<15; i++{
		if i%2 == 0 {
			fmt.Println(i, " <<!---- iteradores Par---->>")
		}else{
			fmt.Println(i, " <<!---- iteradores Impar ---->>")
		}
	}*/
	for i := 0; i<len(array); i++{
		if i%2 == 0 {
			fmt.Println(i, " <<!---- iteradores "+ array[i] +" Par---->>")
		}else{
			fmt.Println(i, " <<!---- iteradores "+ array[i] +" Impar ---->>")
		}
	}

	for _, arr :=  range array{
		fmt.Println(arr)
	}

	momento := time.Now()
	hoy := momento.Weekday()
	switch hoy {
	case 0 :
		fmt.Println("Hoy es DOMINGO")
	case 1 :
		fmt.Println("Hoy es LUNES")
	case 2 :
		fmt.Println("Hoy es MARTES")
	case 3 :
		fmt.Println("Hoy es MIERCOLES")
	case 4 :
		fmt.Println("Hoy es JUEVES")
	case 5 :
		fmt.Println("Hoy es VIERNES")
	default:
		fmt.Println("Hoy es SABADO")

	}




}
