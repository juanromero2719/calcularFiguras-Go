package funciones

import "math"

type FuncionCuadratica struct {
	A float64
	B float64
	C float64
}

func NuevaFuncionCuadratica(a, b, c float64) *FuncionCuadratica {
	return &FuncionCuadratica{
		A: a,
		B: b,
		C: c,
	}
}

func (f *FuncionCuadratica) CalcularFuncionCuadratica() (float64, float64) {
	discriminante := f.B*f.B - 4*f.A*f.C
	raizDiscriminante := math.Sqrt(discriminante)
	x1 := (-f.B + raizDiscriminante) / (2 * f.A)
	x2 := (-f.B - raizDiscriminante) / (2 * f.A)
	return x1, x2
}
