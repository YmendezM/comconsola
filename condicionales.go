package main
import (
	"fmt"
	"os" //ejecutar comandos terminal
	"strconv" //convertidor de variables
)

func main()  {
	fmt.Println("<!----------- Mi programa en GO -------->")
	fmt.Println("Primer parametro " + os.Args[1])

	edad,_ := strconv.Atoi(os.Args[2])

	if edad >= 18 && edad <= 99  {
		fmt.Println("Eres Mayor de edad")
	} else {
		fmt.Println("Eres Menor de edad o un viejito")
	}

	if (edad >= 18 || edad == 17) && edad != 25  {
		fmt.Println("Eres Mayor de edad o tienes 17")
	} else if edad == 16 {
		fmt.Println("Tienes 16")
	}else if edad == 25 {
		fmt.Println("Tienes 25")
	}else  {
		fmt.Println("Eres menor de edad")
	}
}
