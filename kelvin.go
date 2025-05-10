package main

import "fmt"

const Kelvin float64 = 373.2

func main() {

	K := Kelvin
	C := K - 273.15

	fmt.Printf("O ponto de Ebulição da água em Kelvin é °K %g e o Ponto de Ebulição da água em Celcius é °C %g", K, C)

}
