package funciones

type CuatroNumeros struct {
	A int
	B int
	C int
	D int
}

func (cn *CuatroNumeros) Ordenar() [4]int {

	x, y, z, w := cn.A, cn.B, cn.C, cn.D

	if x > y {
		x, y = y, x
	}
	if z > w {
		z, w = w, z
	}

	if x > z {
		x, z = z, x
	}

	if y > w {
		y, w = w, y
	}

	if y > z {
		y, z = z, y
	}

	return [4]int{x, y, z, w}
}
