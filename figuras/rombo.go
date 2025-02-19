package figuras

type Rombo struct {
	DiagonalMayor float64
	DiagonalMenor float64
}

func NuevoRombo(diagonalMayor, diagonalMenor float64) *Rombo {
	return &Rombo{
		DiagonalMayor: diagonalMayor,
		DiagonalMenor: diagonalMenor,
	}
}

func (t Rombo) CalculateArea() float64 {
	return (t.DiagonalMayor * t.DiagonalMenor) / 2

}

func (t Rombo) CalcularPerimetro() float64 {
	return 4 * t.DiagonalMayor
}
