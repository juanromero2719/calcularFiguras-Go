package figuras

type Paralelogramo struct {
	Base   float64
	Altura float64
}

func NuevoParalelogramo(base, altura float64) *Paralelogramo {
	return &Paralelogramo{
		Base:   base,
		Altura: altura,
	}
}

func (t Paralelogramo) CalculateArea() float64 {
	return t.Base * t.Altura
}
