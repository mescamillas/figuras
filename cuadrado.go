package figuras

type Cuadrado struct{
	Ancho float64
	Alto float64
}

func (c *Cuadrado) area() float64 {
	return c.Alto * c.Ancho
}

func (c *Cuadrado) perimetro() float64 {
	return 2*(c.Alto+c.Ancho)
}