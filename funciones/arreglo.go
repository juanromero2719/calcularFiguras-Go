package funciones

import (
	"math/rand"
	"time"
)

type Arreglo struct {
	Valores [10]int
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerarValor() int {
	n := rand.Intn(21) - 10
	if n%2 == 0 {
		return 0
	}
	if n%3 == 0 {
		return 999
	}
	return n
}

func (a *Arreglo) GenerarArreglo() [10]int {
	a.Valores = [10]int{
		GenerarValor(),
		GenerarValor(),
		GenerarValor(),
		GenerarValor(),
		GenerarValor(),
		GenerarValor(),
		GenerarValor(),
		GenerarValor(),
		GenerarValor(),
		GenerarValor(),
	}
	return a.Valores
}
