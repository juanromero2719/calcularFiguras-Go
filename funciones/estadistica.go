package funciones

import "math"

type Estadisticas struct {
	Minimo             float64
	Maximo             float64
	Promedio           float64
	Varianza           float64
	DesviacionEstandar float64
}

func CalcularEstadisticas(n1, n2, n3, n4 float64) Estadisticas {

	minimo := n1
	if n2 < minimo {
		minimo = n2
	}
	if n3 < minimo {
		minimo = n3
	}
	if n4 < minimo {
		minimo = n4
	}

	maximo := n1
	if n2 > maximo {
		maximo = n2
	}
	if n3 > maximo {
		maximo = n3
	}
	if n4 > maximo {
		maximo = n4
	}

	promedio := (n1 + n2 + n3 + n4) / 4.0

	diff1 := n1 - promedio
	diff2 := n2 - promedio
	diff3 := n3 - promedio
	diff4 := n4 - promedio
	varianza := (diff1*diff1 + diff2*diff2 + diff3*diff3 + diff4*diff4) / 4.0

	desviacion := math.Sqrt(varianza)

	return Estadisticas{
		Minimo:             minimo,
		Maximo:             maximo,
		Promedio:           promedio,
		Varianza:           varianza,
		DesviacionEstandar: desviacion,
	}
}
