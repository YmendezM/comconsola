package main
import "fmt"

func main()  {
	fmt.Println("<!----------- Mi programa en GO -------->")
	edad := 25

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
