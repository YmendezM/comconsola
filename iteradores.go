package main

import "fmt"

func main()  {

	array := []string{"n", "n", "n", ""}
	/*for i := 0; i<15; i++{
		if i%2 == 0 {
			fmt.Println(i, " <<!---- iteradores Par---->>")
		}else{
			fmt.Println(i, " <<!---- iteradores Impar ---->>")
		}
	}*/
	for i := 0; i<len(array); i++{
		if i%2 == 0 {
			fmt.Println(i, " <<!---- iteradores Par---->>")
		}else{
			fmt.Println(i, " <<!---- iteradores Impar ---->>")
		}
	}




}
