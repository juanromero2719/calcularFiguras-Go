package figuras

type Rectangulo struct {
	Base   float64
	Altura float64
}

func NuevoRectangulo(base, altura float64) *Rectangulo {
	return &Rectangulo{
		Base:   base,
		Altura: altura,
	}
}

func (t Rectangulo) CalculateArea() float64 {
	return t.Base * t.Altura
}

func (t Rectangulo) CalcularPerimetro() float64 {
	return 2 * (t.Base + t.Altura)
}
