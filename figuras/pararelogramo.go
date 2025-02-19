package figuras

type Paralelogramo struct {
	Base   float64
	Altura float64
	Lado   float64
}

func NuevoParalelogramo(base, altura, lado float64) *Paralelogramo {
	return &Paralelogramo{
		Base:   base,
		Altura: altura,
		Lado:   lado,
	}
}

func (t Paralelogramo) CalculateArea() float64 {
	return t.Base * t.Altura
}

func (t Paralelogramo) CalcularPerimetro() float64 {
	return 2 * (t.Base + t.Lado)
}
