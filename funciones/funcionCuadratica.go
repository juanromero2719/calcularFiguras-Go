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

func CalcularFuncionCuadratica(a, b, c float64) (float64, float64, error) {
	discriminante := b*b - 4*a*c
	raizDiscriminante := math.Sqrt(discriminante)
	x1 := (-b + raizDiscriminante) / (2 * a)
	x2 := (-b - raizDiscriminante) / (2 * a)
	return x1, x2, nil
}

func (f FuncionCuadratica) CalcularFuncionCuadratica() (float64, float64, error) {
	return CalcularFuncionCuadratica(f.A, f.B, f.C)
}
