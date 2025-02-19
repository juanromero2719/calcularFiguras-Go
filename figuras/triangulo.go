package figuras

type Triangulo struct {
	Base   float64
	Altura float64
}

func NuevoTriangulo(base, altura float64) *Triangulo {
	return &Triangulo{
		Base:   base,
		Altura: altura,
	}
}

func (t Triangulo) CalculateArea() float64 {
	return 0.5 * t.Base * t.Altura
}

func (t Triangulo) CalcularPerimetro() float64 {
	return 3 * t.Base
}
