package figuras

import "fmt"

type figura interface{
	area() float64
	perimetro() float64
}

func CalcularArea(figura figura) {
	fmt.Println(figura.area())
}

func CalcularPerimetro(figura figura) {
	fmt.Println(figura.perimetro())
}