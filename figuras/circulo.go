package figuras

type Circulo struct {
	Radio float64
}

func NuevoCirculo(radio float64) *Circulo {
	return &Circulo{
		Radio: radio,
	}
}

func (t Circulo) CalculateArea() float64 {
	return 3.14 * (t.Radio * t.Radio)
}
