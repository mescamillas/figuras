package figuras

import (
	"math"
)

type Circulo struct{
	Radio float64
}

func (c *Circulo) area() float64 {
	return math.Pi * math.Pow(float64(c.Radio),2.0)
}

func (c *Circulo) perimetro() float64 {
	return 2.0 * math.Pi * c.Radio
}