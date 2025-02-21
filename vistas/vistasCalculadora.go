package vistas

import (
	"calcularFiguras/figuras"
	"calcularFiguras/funciones"
	"fmt"
)

type VistasCalculadora struct {
	interfazFigura figuras.Figura
}

func NuevaVistaCalculadora() *VistasCalculadora {
	return &VistasCalculadora{}
}

func (v *VistasCalculadora) MostrarMenu() {
	var opcion int16
	fmt.Println("Menu de Operaciones")
	fmt.Println("1. Calcular area y perimetro de figuras")
	fmt.Println("2. Calcular funcion cuadratica")
	fmt.Println("3. Salir")
	fmt.Scanln(&opcion)

	if opcion <= 0 || opcion > 3 {
		fmt.Println("Opcion no valida")
		v.MostrarMenu()
		return
	}

	switch opcion {
	case 1:
		MenuOperacionesFiguras(v)
	case 2:
		MenuCuadratica(v)
	case 3:
		fmt.Println("Hasta luego")
		return
	}
}

func MenuOperacionesFiguras(v *VistasCalculadora) {
	var opcion int16
	fmt.Println("Seleccione una figura:")
	fmt.Println("1. Circulo")
	fmt.Println("2. Cuadrado")
	fmt.Println("3. Hexagono")
	fmt.Println("4. Paralelogramo")
	fmt.Println("5. Pentagono")
	fmt.Println("6. Rectangulo")
	fmt.Println("7. Rombo")
	fmt.Println("8. Trapecio")
	fmt.Println("9. Triangulo")
	fmt.Println("10. Volver")
	fmt.Scanln(&opcion)

	if opcion <= 0 || opcion > 10 {
		fmt.Println("Opcion no valida")
		MenuOperacionesFiguras(v)
		return // Asegura que no se siga ejecutando el menú
	}

	switch opcion {
	case 1:
		MenuCirculo(v)
	case 2:
		MenuCuadrado(v)
	case 3:
		MenuHexagono(v)
	case 4:
		MenuPararelogramo(v)
	case 5:
		MenuPentagono(v)
	case 6:
		MenuRectangulo(v)
	case 7:
		MenuRombo(v)
	case 8:
		MenuTrapecio(v)
	case 9:
		MenuTriangulo(v)
	case 10:
		v.MostrarMenu()
	}
}

func MenuCuadratica(v *VistasCalculadora) {
	var a, b, c float64

	fmt.Println("Ingrese el valor de a: ")
	fmt.Scanln(&a)

	fmt.Println("Ingrese el valor de b: ")
	fmt.Scanln(&b)

	fmt.Println("Ingrese el valor de c: ")
	fmt.Scanln(&c)

	fc := funciones.NuevaFuncionCuadratica(1, -3, 2)
	x1, x2 := fc.CalcularFuncionCuadratica()

	fmt.Printf("x1: %.2f\n", x1)
	fmt.Printf("x2: %.2f\n", x2)

	v.MostrarMenu()
}

func MenuCirculo(v *VistasCalculadora) {
	var radio float64

	fmt.Println("Ingrese el radio del circulo: ")
	fmt.Scanln(&radio)

	interfazFigura := figuras.NuevoCirculo(radio)

	fmt.Printf("Área del círculo: %.2f\n", interfazFigura.CalculateArea())
	fmt.Printf("Perímetro del círculo: %.2f\n", interfazFigura.CalcularPerimetro())

	v.MostrarMenu()
}

func MenuCuadrado(v *VistasCalculadora) {
	var lado float64

	fmt.Println("Ingrese el lado del cuadrado: ")
	fmt.Scanln(&lado)

	interfazFigura := figuras.NuevoCuadrado(lado)

	// Mostrar área y perímetro del cuadrado
	fmt.Printf("Área del cuadrado: %.2f\n", interfazFigura.CalculateArea())
	fmt.Printf("Perímetro del cuadrado: %.2f\n", interfazFigura.CalcularPerimetro())

	v.MostrarMenu()
}

func MenuHexagono(v *VistasCalculadora) {
	var lado, apotema float64

	fmt.Println("Ingrese el lado del hexagono: ")
	fmt.Scanln(&lado)

	fmt.Println("Ingrese el apotema del hexagono:")
	fmt.Scanln(&apotema)

	interfazFigura := figuras.NuevoHexagono(lado, apotema)

	fmt.Printf("Área del hexagono: %.2f\n", interfazFigura.CalculateArea())
	fmt.Printf("Perímetro del hexagono: %.2f\n", interfazFigura.CalcularPerimetro())

	v.MostrarMenu()
}

func MenuPararelogramo(v *VistasCalculadora) {
	var base, altura, lado float64

	fmt.Println("Ingrese la base del paralelogramo: ")
	fmt.Scanln(&base)

	fmt.Println("Ingrese la altura del paralelogramo:")
	fmt.Scanln(&altura)

	fmt.Println("Ingrese el lado del paralelogramo:")
	fmt.Scanln(&lado)

	interfazFigura := figuras.NuevoParalelogramo(base, altura, lado)

	fmt.Printf("Área del paralelogramo: %.2f\n", interfazFigura.CalculateArea())
	fmt.Printf("Perímetro del paralelogramo: %.2f\n", interfazFigura.CalcularPerimetro())

	v.MostrarMenu()
}

func MenuPentagono(v *VistasCalculadora) {
	var lado, apotema float64

	fmt.Println("Ingrese el lado del pentagono: ")
	fmt.Scanln(&lado)

	fmt.Println("Ingrese el apotema del pentagono:")
	fmt.Scanln(&apotema)

	interfazFigura := figuras.NuevoPentagono(lado, apotema)

	fmt.Printf("Área del pentagono: %.2f\n", interfazFigura.CalculateArea())
	fmt.Printf("Perímetro del pentagono: %.2f\n", interfazFigura.CalcularPerimetro())

	v.MostrarMenu()
}

func MenuRectangulo(v *VistasCalculadora) {
	var base, altura float64

	fmt.Println("Ingrese la base del rectangulo: ")
	fmt.Scanln(&base)

	fmt.Println("Ingrese la altura del rectangulo:")
	fmt.Scanln(&altura)

	interfazFigura := figuras.NuevoRectangulo(base, altura)

	fmt.Printf("Área del rectangulo: %.2f\n", interfazFigura.CalculateArea())
	fmt.Printf("Perímetro del rectangulo: %.2f\n", interfazFigura.CalcularPerimetro())

	v.MostrarMenu()
}

func MenuRombo(v *VistasCalculadora) {
	var lado, altura float64

	fmt.Println("Ingrese el diagonal mayor del rombo: ")
	fmt.Scanln(&lado)

	fmt.Println("Ingrese la diagonal menor del rombo:")
	fmt.Scanln(&altura)

	interfazFigura := figuras.NuevoRombo(lado, altura)

	fmt.Printf("Área del rombo: %.2f\n", interfazFigura.CalculateArea())
	fmt.Printf("Perímetro del rombo: %.2f\n", interfazFigura.CalcularPerimetro())

	v.MostrarMenu()
}

func MenuTrapecio(v *VistasCalculadora) {
	var baseMayor, baseMenor, altura float64

	fmt.Println("Ingrese la base mayor del trapecio: ")
	fmt.Scanln(&baseMayor)

	fmt.Println("Ingrese la base menor del trapecio:")
	fmt.Scanln(&baseMenor)

	fmt.Println("Ingrese la altura del trapecio:")
	fmt.Scanln(&altura)

	interfazFigura := figuras.NuevoTrapecio(baseMayor, baseMenor, altura)

	fmt.Printf("Área del trapecio: %.2f\n", interfazFigura.CalculateArea())
	fmt.Printf("Perímetro del trapecio: %.2f\n", interfazFigura.CalcularPerimetro())

	v.MostrarMenu()
}

func MenuTriangulo(v *VistasCalculadora) {
	var base, altura float64

	fmt.Println("Ingrese la base del triangulo: ")
	fmt.Scanln(&base)

	fmt.Println("Ingrese la altura del triangulo:")
	fmt.Scanln(&altura)

	interfazFigura := figuras.NuevoTriangulo(base, altura)

	fmt.Printf("Área del triangulo: %.2f\n", interfazFigura.CalculateArea())
	fmt.Printf("Perímetro del triangulo: %.2f\n", interfazFigura.CalcularPerimetro())

	v.MostrarMenu()
}
