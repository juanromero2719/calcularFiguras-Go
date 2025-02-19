package figuras

type Cuadrado struct {
	Lado float64
}

func NuevoCuadrado(lado float64) *Cuadrado {
	return &Cuadrado{
		Lado: lado,
	}
}

func (t Cuadrado) CalculateArea() float64 {
	return t.Lado * t.Lado
}

func (t Cuadrado) CalcularPerimetro() float64 {
	return t.Lado * 4
}
