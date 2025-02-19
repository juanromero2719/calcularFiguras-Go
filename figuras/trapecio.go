package figuras

type Trapecio struct {
	BaseInferior float64
	BaseSuperior float64
	Altura       float64
}

func NuevoTrapecio(baseInferior, baseSuperior, altura float64) *Trapecio {
	return &Trapecio{
		BaseInferior: baseInferior,
		BaseSuperior: baseSuperior,
		Altura:       altura,
	}
}

func (t Trapecio) CalculateArea() float64 {
	return ((t.BaseInferior + t.BaseSuperior) / 2) * t.Altura
}
