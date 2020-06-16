package main

import "fmt"

func main() {
	var numero float32
	var numero2 int
	var divisor int = 100
	fmt.Println("Indica el valor del cual quieres sacar el porcentaje")

	fmt.Scan(&numero)

	fmt.Println("Indica que porcentaje quieres saber:")

	fmt.Scan(&numero2)

	var porcentaje float32 = float32(numero2) / float32(divisor)

	resultado := porcentaje * numero

	fmt.Println(numero2, "% De  ", numero, " es ", resultado)

	restante := 100 - numero2

	porcentaje2 := float32(restante) / float32(divisor)

	resultado2 := porcentaje2 * numero

	fmt.Println("Por si te interesa tambien te doy el porcentaje restante", restante, "%", "de", numero, "es", resultado2)

	fmt.Println("Presiona cualquier tecla para salir")
	var tecla string
	fmt.Scan(&tecla)

}
