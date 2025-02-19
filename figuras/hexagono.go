package figuras

type Hexagono struct {
	Lado    float64
	Apotema float64
}

func NuevoHexagono(lado, apotema float64) *Hexagono {
	return &Hexagono{
		Lado:    lado,
		Apotema: apotema,
	}
}

func (t Hexagono) CalculateArea() float64 {
	perimetro := 6 * t.Lado
	return (perimetro * t.Apotema) / 2
}
