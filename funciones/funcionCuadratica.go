package funciones

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

type FuncionCuadratica struct {
	Ecuacion string
}

func (f *FuncionCuadratica) Resolver() (float64, float64, error) {

	re := regexp.MustCompile(`^([+-]?\d*)x\^2([+-]\d*)x([+-]\d+)=0$`)
	matches := re.FindStringSubmatch(f.Ecuacion)
	if matches == nil {
		return 0, 0, fmt.Errorf("formato de ecuación inválido")
	}

	aStr := matches[1]
	bStr := matches[2]
	cStr := matches[3]

	var a float64
	if aStr == "" || aStr == "+" {
		a = 1
	} else if aStr == "-" {
		a = -1
	} else {
		var err error
		a, err = strconv.ParseFloat(aStr, 64)
		if err != nil {
			return 0, 0, fmt.Errorf("error en el coeficiente a")
		}
	}

	var b float64
	if bStr == "+" {
		b = 1
	} else if bStr == "-" {
		b = -1
	} else {
		var err error
		b, err = strconv.ParseFloat(bStr, 64)
		if err != nil {
			return 0, 0, fmt.Errorf("error en el coeficiente b")
		}
	}

	c, err := strconv.ParseFloat(cStr, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("error en el coeficiente c")
	}

	discriminante := b*b - 4*a*c
	if discriminante < 0 {
		return 0, 0, fmt.Errorf("la ecuación tiene raíces complejas")
	}

	sqrtDisc := math.Sqrt(discriminante)
	x1 := (-b + sqrtDisc) / (2 * a)
	x2 := (-b - sqrtDisc) / (2 * a)
	return x1, x2, nil
}
