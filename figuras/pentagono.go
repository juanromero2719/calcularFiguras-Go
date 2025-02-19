package figuras

type Pentagono struct {
	Lado    float64
	Apotema float64
}

func NuevoPentagono(lado, apotema float64) *Pentagono {
	return &Pentagono{
		Lado:    lado,
		Apotema: apotema,
	}
}

func (t Pentagono) CalculateArea() float64 {
	perimetro := 5 * t.Lado
	return ((perimetro * t.Apotema) / 2)
}

func (t Pentagono) CalcularPerimetro() float64 {
	return t.Lado * 5
}
